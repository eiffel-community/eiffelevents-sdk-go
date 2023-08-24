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
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
)

//go:generate go run ./../internal/cmd/dirsync ../protocol/schemas/ ./schemas

//go:embed schemas
var schemas embed.FS

// BundledSchemaLocator locates schemas from those built into the binary via this package.
type BundledSchemaLocator struct{}

func NewBundledSchemaLocator() *BundledSchemaLocator {
	return &BundledSchemaLocator{}
}

// GetSchema returns the event's schema from the built-in official set of schemas.
// Returns (nil, nil) if no schema was found for the provided event type and version.
func (bsl *BundledSchemaLocator) GetSchema(ctx context.Context, eventType string, version string, schemaURI string) (io.ReadCloser, error) {
	schema, err := schemas.ReadFile(filepath.Join("schemas", eventType, version+".json"))
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error reading schema: %w", err)
	}
	return io.NopCloser(bytes.NewReader(schema)), nil
}
