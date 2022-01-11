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
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/tidwall/gjson"
	"github.com/xeipuuv/gojsonschema"
)

var ErrSchemaMissing = errors.New("no schema found by any of the configured schema locators")

// SchemaLocator attempts to find the schema for a particular event based
// on the event's type, its version, and (if available) the schema URI from
// the meta.schemaURI member. Returns (nil, nil) if no schema could be found
// (and no error occurred while reaching that conclusion).
type SchemaLocator interface {
	GetSchema(ctx context.Context, eventType string, version string, schemaURI string) (io.ReadCloser, error)
}

// SchemaValidator is a Validator instance that locates a suitable
// JSON schema for an event and validates the event. Loaded schemas
// are cached indefinitely with the event type, event version, and
// schema URI (from the meta.schemaURI member) as the cache key.
type SchemaValidator struct {
	schemaCache    map[string]*gojsonschema.Schema
	schemaCacheMu  sync.RWMutex
	schemaLocators []SchemaLocator
}

func NewSchemaValidator(schemaLocators ...SchemaLocator) *SchemaValidator {
	return &SchemaValidator{
		schemaCache:    make(map[string]*gojsonschema.Schema),
		schemaLocators: schemaLocators,
	}
}

// Validate runs through the configured schema locators to find one that has
// a schema for the provided event, and proceeds to validate the event against
// the schema. Returns an ErrSchemaMissing error if no schema could be located
// and SchemaValidationError if the validation fails.
func (sv *SchemaValidator) Validate(ctx context.Context, event []byte) error {
	var (
		typ       string
		version   string
		schemaURI string
	)
	fields := gjson.GetManyBytes(event, "meta.type", "meta.version", "meta.schemaUri")
	if fields[0].Type == gjson.String {
		typ = fields[0].String()
	}
	if typ == "" {
		return fmt.Errorf("missing or invalid contents of meta.type field: %q", typ)
	}
	if fields[1].Type == gjson.String {
		version = fields[1].String()
	}
	if version == "" {
		return fmt.Errorf("missing or invalid contents of meta.version field: %q", version)
	}
	if fields[2].Type == gjson.String {
		schemaURI = fields[2].String()
	}

	schema, err := sv.getSchema(ctx, typ, version, schemaURI)
	if err != nil {
		return err
	}

	result, err := schema.Validate(gojsonschema.NewBytesLoader(event))
	if err != nil {
		return fmt.Errorf("error validating event: %w", err)
	}
	if !result.Valid() {
		return &SchemaValidationError{errors: result.Errors()}
	}
	return nil
}

func (sv *SchemaValidator) cacheKey(eventType string, version string, schemaURI string) string {
	return eventType + "\n" + version + "\n" + schemaURI
}

func (sv *SchemaValidator) getSchema(ctx context.Context, eventType string, version string, schemaURI string) (*gojsonschema.Schema, error) {
	// Use cached schema loader if available.
	cacheKey := sv.cacheKey(eventType, version, schemaURI)
	sv.schemaCacheMu.RLock()
	cachedSchema, exists := sv.schemaCache[cacheKey]
	sv.schemaCacheMu.RUnlock()
	if exists {
		return cachedSchema, nil
	}

	// Iterate over schema locators and return the first available schema.
	for _, src := range sv.schemaLocators {
		schemaReader, err := src.GetSchema(ctx, eventType, version, schemaURI)
		if err != nil {
			return nil, fmt.Errorf("error finding schema for (%q, %q, %q): %w", eventType, version, schemaURI, err)
		}
		if schemaReader == nil {
			continue
		}
		defer schemaReader.Close()

		schemaBytes, err := io.ReadAll(schemaReader)
		if err != nil {
			return nil, fmt.Errorf("error reading schema for (%q, %q, %q): %w", eventType, version, schemaURI, err)
		}

		// It initially seemed like a good idea to use NewReaderLoader instead and
		// let gojsonschema do the I/O, but that function has a weird interface.
		schema, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(schemaBytes))
		if err != nil {
			return nil, fmt.Errorf("error compiling schema for (%q, %q, %q): %w", eventType, version, schemaURI, err)
		}
		sv.schemaCacheMu.Lock()
		sv.schemaCache[cacheKey] = schema
		sv.schemaCacheMu.Unlock()
		return schema, nil
	}
	return nil, fmt.Errorf("error finding schema for (%q, %q, %q): %w", eventType, version, schemaURI, ErrSchemaMissing)
}

// SchemaValidationError indicates that the event failed validation against the JSON schema.
type SchemaValidationError struct {
	errors []gojsonschema.ResultError
}

func (ve *SchemaValidationError) Error() string {
	var s strings.Builder
	s.WriteString("The event failed the schema validation with the following error(s):")
	for _, re := range ve.errors {
		s.WriteByte('\n')
		s.WriteString(re.String())
	}
	return s.String()
}

func (ve *SchemaValidationError) Is(target error) bool {
	_, ok := target.(*SchemaValidationError)
	return ok
}
