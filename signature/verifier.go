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
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/go-ldap/ldap"
	"github.com/gowebpki/jcs"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/eiffel-community/eiffelevents-sdk-go"
)

// PublicKeyLocator locates one or more public keys that can be used
// to verify the signature of a given identity.
type PublicKeyLocator interface {
	// Locate returns one or more public keys corresponding to the provided identity,
	// or an empty or nil slice if no public keys were found. An error return indicates
	// that the lookup itself failed.
	Locate(ctx context.Context, identity *AuthorIdentity) ([]crypto.PublicKey, error)
}

// AuthorIdentity is a representation of the distinguished name
// in the meta.security.authorIdentity field of an Eiffel event.
// It can be compared to other values of the same type.
type AuthorIdentity struct {
	dn       *ldap.DN
	original string
}

func NewAuthorIdentity(s string) (*AuthorIdentity, error) {
	dn, err := ldap.ParseDN(s)
	if err != nil {
		return nil, fmt.Errorf("error parsing author identity %q: %w", s, err)
	}
	return &AuthorIdentity{
		dn:       dn,
		original: s,
	}, nil
}

// Equal returns true if the provided *AuthorIdentity is equal to this one,
// igoring differences in whitespace etc.
func (ai *AuthorIdentity) Equal(other *AuthorIdentity) bool {
	return ai.dn.Equal(other.dn)
}

func (ai *AuthorIdentity) String() string {
	return ai.original
}

// Verifier can verify whether the signature of a given Eiffel event matches
// any of the keys known by the associated PublicKeyLocator.
type Verifier struct {
	keyLocator    PublicKeyLocator
	identityCache map[string]*AuthorIdentity
}

func NewVerifier(keyLocator PublicKeyLocator) *Verifier {
	return &Verifier{
		keyLocator:    keyLocator,
		identityCache: make(map[string]*AuthorIdentity),
	}
}

// Verify attempts to verify the signature of the provided event payload,
// using the Verifier's PublicKeyLocator to obtain a public key suitable
// for verifying the event.
//
//   - If the event can't be verified because of its contents, e.g. because
//     it doesn't include a signature, ErrUnverifiableEvent is returned.
//   - If the event is signed with an unsupported algorithm,
//     ErrUnsupportedAlgorithm is returned.
//   - If no public key that matches the event sender's identity was found,
//     ErrPublicKeyNotFound is returned.
//   - If the public key that was found doesn't match the algorithm in
//     the event payload, ErrKeyTypeMismatch is returned.
//   - If something goes wrong while modifying the event in preparation of
//     the verification, ErrMarshaling is returned.
//   - If the verification itself fails for all of the public keys,
//     ErrVerificationFailed is returned. This error will wrap the individual
//     errors returned for each public key, including ErrSignatureMismatch.
//
// Errors will be returned in wrapped form so make sure you use errors.Is
// rather than direct comparisons.
func (v *Verifier) Verify(ctx context.Context, event []byte) error {
	// Extract the signature itself and the other fields we need for
	// the verification and return an error if either of them are missing.
	values := gjson.GetManyBytes(event, algorithmField, authorIdentityField, signatureField)
	alg := values[0].String()
	identity := values[1].String()
	sig := values[2].String()

	if alg == "" {
		return fmt.Errorf("%w: %s", ErrUnverifiableEvent, algorithmField)
	}
	if identity == "" {
		return fmt.Errorf("%w: %s", ErrUnverifiableEvent, authorIdentityField)
	}
	if sig == "" {
		return fmt.Errorf("%w: %s", ErrUnverifiableEvent, signatureField)
	}

	var (
		hash       crypto.Hash
		hashFunc   func([]byte) []byte
		verifyFunc func(pk crypto.PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error
	)

	switch alg {
	case string(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_RS256):
		hash = crypto.SHA256
		hashFunc = hashSHA256
		verifyFunc = verifyPKCS1v15
	case string(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_RS384):
		hash = crypto.SHA384
		hashFunc = hashSHA384
		verifyFunc = verifyPKCS1v15
	case string(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_RS512):
		hash = crypto.SHA512
		hashFunc = hashSHA512
		verifyFunc = verifyPKCS1v15
	case string(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_ES256):
		hash = crypto.SHA256
		hashFunc = hashSHA256
		verifyFunc = verifyECDSA
	case string(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_ES384):
		hash = crypto.SHA384
		hashFunc = hashSHA384
		verifyFunc = verifyECDSA
	case string(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_ES512):
		hash = crypto.SHA512
		hashFunc = hashSHA512
		verifyFunc = verifyECDSA
	case string(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_PS256):
		hash = crypto.SHA256
		hashFunc = hashSHA256
		verifyFunc = verifyPSS
	case string(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_PS384):
		hash = crypto.SHA384
		hashFunc = hashSHA384
		verifyFunc = verifyPSS
	case string(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_PS512):
		hash = crypto.SHA512
		hashFunc = hashSHA512
		verifyFunc = verifyPSS
	default:
		return fmt.Errorf("%w: %s", ErrUnsupportedAlgorithm, alg)
	}

	// Clear the signature field and transform the event to canonical JSON.
	var err error
	if event, err = sjson.SetBytes(event, signatureField, ""); err != nil {
		return errors.Join(ErrMarshaling, err)
	}
	if event, err = jcs.Transform(event); err != nil {
		return errors.Join(ErrMarshaling, err)
	}

	// DecodedLen returns the worst case decoded length (when no padding was
	// required), thus the amount we need to allocate. Using the actual number
	// of decoded bytes to truncate sigBytes afterwards is crucial to avoid
	// trailing null bytes.
	sigBytes := make([]byte, base64.RawStdEncoding.DecodedLen(len(sig)))
	n, err := base64.StdEncoding.Decode(sigBytes, []byte(sig))
	if err != nil {
		return errors.Join(ErrMarshaling, err)
	}
	sigBytes = sigBytes[:n]

	// Use the identity cache to avoid having to reparse the same DN string over and over.
	dn, found := v.identityCache[identity]
	if !found {
		if dn, err = NewAuthorIdentity(identity); err != nil {
			return errors.Join(ErrMarshaling, err)
		}
		v.identityCache[identity] = dn
	}

	keys, err := v.keyLocator.Locate(ctx, dn)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrPublicKeyLookup, identity)
	}
	if len(keys) == 0 {
		return fmt.Errorf("%w: %s", ErrPublicKeyNotFound, identity)
	}

	// Collect the error for each public key we try, and start with
	// ErrVerificationFailed to represent the failure of the whole operation.
	errs := []error{ErrVerificationFailed}
	hashBytes := hashFunc(event)
	for _, key := range keys {
		err = verifyFunc(key, hash, hashBytes, sigBytes)
		if err == nil {
			return nil
		}
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}

func hashSHA256(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}

func hashSHA384(data []byte) []byte {
	h := sha512.Sum384(data)
	return h[:]
}

func hashSHA512(data []byte) []byte {
	h := sha512.Sum512(data)
	return h[:]
}

func verifyECDSA(pub crypto.PublicKey, hash crypto.Hash, digest []byte, sig []byte) error {
	pubECDSA, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("%w; expected *ecdsa.PublicKey, got %T", ErrKeyTypeMismatch, pub)
	}
	if ecdsa.VerifyASN1(pubECDSA, digest, sig) {
		return nil
	} else {
		return ErrSignatureMismatch
	}
}

func verifyPKCS1v15(pub crypto.PublicKey, hash crypto.Hash, digest []byte, sig []byte) error {
	pubRSA, ok := pub.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("%w; expected *rsa.PublicKey, got %T", ErrKeyTypeMismatch, pub)
	}
	if err := rsa.VerifyPKCS1v15(pubRSA, hash, digest, sig); err != nil {
		return errors.Join(ErrSignatureMismatch, err)
	}
	return nil
}

func verifyPSS(pub crypto.PublicKey, hash crypto.Hash, digest []byte, sig []byte) error {
	pubRSA, ok := pub.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("%w; expected *rsa.PublicKey, got %T", ErrKeyTypeMismatch, pub)
	}
	if err := rsa.VerifyPSS(pubRSA, hash, digest, sig, nil); err != nil {
		return errors.Join(ErrSignatureMismatch, err)
	}
	return nil
}
