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
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	eiffelevents "github.com/eiffel-community/eiffelevents-sdk-go/editions/lyon"
)

func ExampleSchemaValidator() {
	event, _ := eiffelevents.NewCompositionDefined()
	event.Data.Name = "name-of-composition"

	v := DefaultSet()
	if err := v.Validate(context.Background(), []byte(event.String())); err != nil {
		fmt.Printf("Validation failed: %s", err)
		return
	}
	fmt.Println("Validation passed!")
	// Output: Validation passed!
}

func TestSchemaValidator(t *testing.T) {
	testcases := []struct {
		name           string
		eventFile      string
		schemaLocators []SchemaLocator
		errorIs        error
	}{
		{
			name:      "Successful validation with single schema locator",
			eventFile: "valid_event_without_schemauri.json",
			schemaLocators: []SchemaLocator{
				NewBundledSchemaLocator(),
			},
		},
		{
			name:      "Failed validation with single schema locator",
			eventFile: "invalid_event_without_schemauri.json",
			schemaLocators: []SchemaLocator{
				NewBundledSchemaLocator(),
			},
			errorIs: &SchemaValidationError{},
		},
		{
			// Let the first schema locator return a technically valid but bogus schema
			// and verify that the validation fails.
			name:      "First schema locator match is used",
			eventFile: "valid_event_without_schemauri.json",
			schemaLocators: []SchemaLocator{
				&fixedSchemaLocator{
					schema: `{"$schema": "http://json-schema.org/draft-04/schema#", "type": "string"}`,
				},
				NewBundledSchemaLocator(),
			},
			errorIs: &SchemaValidationError{},
		},
		{
			// Make sure the validation passes even though
			// the first schema locator doesn't return a schema.
			name:      "Unmatched schema locator is ignored",
			eventFile: "valid_event_without_schemauri.json",
			schemaLocators: []SchemaLocator{
				&nullSchemaLocator{},
				NewBundledSchemaLocator(),
			},
		},
		{
			name:      "Return error if no schema found",
			eventFile: "valid_event_without_schemauri.json",
			schemaLocators: []SchemaLocator{
				&nullSchemaLocator{},
			},
			errorIs: ErrSchemaMissing,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			validator := NewSchemaValidator(tc.schemaLocators...)
			event, err := os.ReadFile(filepath.Join("testdata", tc.eventFile))
			require.NoError(t, err)
			err = validator.Validate(t.Context(), event)
			assert.ErrorIs(t, err, tc.errorIs)
		})
	}
}

// fixedSchemaLocator returns a fixed string as the schema every time it's asked.
type fixedSchemaLocator struct {
	schema string
}

func (fsl *fixedSchemaLocator) GetSchema(ctx context.Context, eventType string, version string, schemaURI string) (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader(fsl.schema)), nil
}

// nullSchemaLocator always fails to locate any schemas.
type nullSchemaLocator struct{}

func (nsl *nullSchemaLocator) GetSchema(ctx context.Context, eventType string, version string, schemaURI string) (io.ReadCloser, error) {
	return nil, nil
}
