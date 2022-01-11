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

package validator

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// HTTPGetter is capable of downloading an HTTP resource via a GET request.
type HTTPGetter interface {
	Do(req *http.Request) (*http.Response, error)
}

// MetaSchemaLocator is a SchemaLocator implementation that downloads
// an event's schema from its non-empty meta.schemaURI member via HTTP(S).
type MetaSchemaLocator struct {
	getter HTTPGetter
}

func NewMetaSchemaLocator(getter HTTPGetter) *MetaSchemaLocator {
	return &MetaSchemaLocator{
		getter: getter,
	}
}

// GetSchema downloads the event's schema via HTTP(S) if schemaURI is non-empty.
// Returns (nil, nil) if schemaURI is empty or if the URI scheme isn't "http"
// or "https".
func (bss *MetaSchemaLocator) GetSchema(ctx context.Context, eventType string, version string, schemaURI string) (io.ReadCloser, error) {
	if schemaURI == "" {
		return nil, nil
	}

	// We only support HTTP(S) URI schemes.
	u, err := url.Parse(schemaURI)
	if err != nil {
		return nil, fmt.Errorf("error parsing schema URI: %w", err)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, schemaURI, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}
	resp, err := bss.getter.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching schema: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to fetch schema returned status %d", resp.StatusCode)
	}
	return resp.Body, nil
}
