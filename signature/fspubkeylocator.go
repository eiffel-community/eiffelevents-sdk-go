// Copyright Axis Communications AB.
//
// For a full list of individual contributors, please see the commit history.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package signature

import (
	"context"
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const tracerName = "github.com/eiffel-community/eiffelevents-sdk-go/signature"

// FSPublicKeyLocator is a PublicKeyLocator implementation that locates
// public keys in a file system directory containing PEM files. The files should
// be named after the distinguished name (DN) of its owner, plus a .pem extension,
// e.g. "CN=joe,O=Acme.pem".
//
// A PEM file may contain multiple keys for an identity, and there may be multiple
// PEM files for a given identity, if different but equivalent representations of
// the DN are used. For example, "CN=joe,O=Acme.pem" and "cn=joe, o=Acme.pem" contain
// equivalent DNs and the keys in both files are returned if a lookup is made for
// any of those DN (or some other equivalent form of that DN).
type FSPublicKeyLocator struct {
	cfg       FSPublicKeyLocatorConfig
	keyLoader func(pemData []byte) ([]crypto.PublicKey, error) // Allows mocking in tests.

	// Fields protected by the mutex.
	mu       sync.RWMutex
	keyCache []keyCacheEntry
	lastScan time.Time
}

type FSPublicKeyLocatorConfig struct {
	// KeyDirectory is path to the file system directory containing the PEM files
	// from which public keys should be loaded.
	KeyDirectory string `json:"key_directory" yaml:"key_directory"`

	// CacheTTL is how old the key cache is allowed to get before it's rescanned
	// from disk. Zero means that the cache is disabled.
	CacheTTL time.Duration `json:"cache_ttl" yaml:"cache_ttl"`
}

type keyCacheEntry struct {
	identity *AuthorIdentity
	keys     []crypto.PublicKey
}

func NewFSPublicKeyLocator(cfg FSPublicKeyLocatorConfig) *FSPublicKeyLocator {
	return &FSPublicKeyLocator{
		cfg:       cfg,
		keyLoader: publicKeysFromPEMData,
		keyCache:  make([]keyCacheEntry, 0, 50),
	}
}

// Locate looks up the given identity and returns a set of matching public keys.
// If no keys match an empty or nil slice is returned.
func (pkl *FSPublicKeyLocator) Locate(ctx context.Context, identity *AuthorIdentity) ([]crypto.PublicKey, error) {
	if err := pkl.MaybeScan(ctx); err != nil {
		return nil, fmt.Errorf("error refreshing key cache: %w", err)
	}

	// Using a slice of structs for looking up the DNs isn't terribly efficient,
	// but there's no simple way of rewriting the DNs on a canonical form that
	// would fit into a map key. We only have the AuthorIdentity.Equal. If this
	// turns out to be a performance bottleneck there are some obvious
	// optimizations, like attempting the most recently used key first.
	pkl.mu.RLock()
	defer pkl.mu.RUnlock()
	var result []crypto.PublicKey
	for _, entry := range pkl.keyCache {
		if entry.identity.Equal(identity) {
			result = append(result, entry.keys...)
		}
	}
	return result, nil
}

// MaybeScan (re)scans the directory of public keys if the cache's TTL has expired
// or the TTL is disabled.
func (pkl *FSPublicKeyLocator) MaybeScan(ctx context.Context) (err error) {
	pkl.mu.Lock()
	defer pkl.mu.Unlock()

	if pkl.cfg.CacheTTL != 0 && time.Since(pkl.lastScan) <= pkl.cfg.CacheTTL {
		return nil
	}

	_, span := otel.GetTracerProvider().Tracer(tracerName).
		Start(ctx, "Rescan public keys", trace.WithSpanKind(trace.SpanKindInternal))
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	files, err := os.ReadDir(pkl.cfg.KeyDirectory)
	if err != nil {
		return fmt.Errorf("error scanning public key directory: %w", err)
	}

	// Clear the cache before starting to repopulate it. If there's an error
	// scanning for keys we'll just continuously return errors, i.e. we won't
	// attempt to be clever and use stale data until the problem has been
	// corrected. This isn't so much for philosophical reasons but rather to
	// keep the code simple. It's up to the caller to throttle retries.
	pkl.keyCache = pkl.keyCache[:0]
	for _, f := range files {
		if filepath.Ext(f.Name()) != ".pem" {
			continue
		}
		pemData, err := os.ReadFile(filepath.Join(pkl.cfg.KeyDirectory, f.Name()))
		if errors.Is(err, fs.ErrNotExist) {
			// It's okay if the file doesn't exist. It probably indicates that someone
			// changed the contents of the key directory after we read its contents,
			// and we're supposed to support hot reloads of the keys.
			continue
		} else if err != nil {
			return fmt.Errorf("error reading PEM file: %w", err)
		}
		keys, err := pkl.keyLoader(pemData)
		if err != nil {
			return fmt.Errorf("error extracting public keys from %q: %w", f.Name(), err)
		}
		identity, err := NewAuthorIdentity(strings.TrimSuffix(f.Name(), ".pem"))
		if err != nil {
			return fmt.Errorf("error parsing %q as a DN: %w", f.Name(), err)
		}
		pkl.keyCache = append(pkl.keyCache, keyCacheEntry{identity: identity, keys: keys})
	}
	pkl.lastScan = time.Now().UTC()
	return nil
}

func publicKeysFromPEMData(pemData []byte) ([]crypto.PublicKey, error) {
	var result []crypto.PublicKey
	for block, remaining := pem.Decode(pemData); block != nil; block, remaining = pem.Decode(remaining) {
		// The block types used in PEM files are a mess. Trigger on all
		// conceivable block types to be on the safe side. Block types
		// picked up from https://stackoverflow.com/a/5356351/414355.
		switch block.Type {
		case "PUBLIC KEY", "RSA PUBLIC KEY", "ECDSA PUBLIC KEY":
			key, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("error parsing %q PEM block: %w", block.Type, err)
			}
			result = append(result, key)
		}
	}
	return result, nil
}
