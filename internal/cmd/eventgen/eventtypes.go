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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	jsschema "github.com/lestrrat-go/jsschema"
	"gopkg.in/yaml.v3"

	"github.com/eiffel-community/eiffelevents-sdk-go/internal/codetemplate"
)

// goType represents a Go type that's capable of formatting itself
// as a string that can be included in a variable declaration or
// struct field definition.
type goType interface {
	fmt.Stringer
}

// goTypeDeclarer represents a Go type that needs to declare itself in
// the top-most scope of a source file, e.g. a struct or enum.
type goTypeDeclarer interface {
	// declare writes a type declaration as Go code.
	declare(ct *codetemplate.OutputFile) error
}

// goPrimitiveType represents a primitype Go type, e.g. int, bool, or string.
type goPrimitiveType struct {
	name string
}

func (t *goPrimitiveType) String() string {
	return t.name
}

// goInterface represents the empty interface Go type, interface{}.
type goInterface struct{}

func (t *goInterface) String() string {
	return "interface{}"
}

// generateTypes generates Go struct declarations (and types referenced in
// those structs) for the latest version within each major version of each type.
func generateTypes(schemaDefs map[string][]schemaDefinitionRenderer, outputDir string) error {
	for _, defs := range schemaDefs {
		for significantVersion, schema := range significantVersions(defs) {
			schemaFile, err := os.Open(schema.Filename())
			if err != nil {
				return err
			}
			defer schemaFile.Close()

			// Convert input YAML to JSON acceptable to github.com/lestrrat-go/jsschema.
			var def interface{}
			if err := yaml.NewDecoder(schemaFile).Decode(&def); err != nil {
				return err
			}
			var jsonDef bytes.Buffer
			if err := json.NewEncoder(&jsonDef).Encode(def); err != nil {
				return err
			}

			outputFile := filepath.Join(outputDir, fmt.Sprintf("%sV%s.go", schema.TypeName(), strings.ReplaceAll(significantVersion, ".", "_")))
			if err = schema.Render(&jsonDef, outputFile); err != nil {
				return err
			}
		}
	}
	return nil
}

// significantVersions inspects a list of schema definitions for a single type and maps the type
// versions to generate to the corresponding schemaDefinitionRenderers. For non-experimental
// event versions, the type version will be the major version of the most recent major.minor.patch
// version within each major version will be returned, while all experimental event versions will
// be returned with the type version set to the full major.minor.patch version.
//
// For example, given the versions (0.1.0, 0.2.0, 1.0.0, 1.1.0, 2.0.0) a map with the keys
// (0.1.0, 0.2.0, 1, 2) will be returned.
func significantVersions(schemas []schemaDefinitionRenderer) map[string]schemaDefinitionRenderer {
	versions := map[string]schemaDefinitionRenderer{}
	for _, schema := range schemas {
		// Keep all schemas for experimental versions.
		if schema.Version().Major() == 0 {
			versions[schema.Version().String()] = schema
			continue
		}

		// Only keep the latest non-experimental version.
		majorStr := strconv.Itoa(int(schema.Version().Major()))
		if current, exists := versions[majorStr]; !exists || current.Version().LessThan(schema.Version()) {
			versions[majorStr] = schema
		}
	}
	return versions
}

// goTypeFromSchema returns a goType that represents a node in a JSON schema.
func goTypeFromSchema(parent *goStruct, name string, schema *jsschema.Schema) (goType, error) {
	// Are we dealing with a JSON reference, i.e. a $ref key pointing to another schema file?
	if len(schema.Reference) != 0 {
		return newPredeclaredType(schema.Reference)
	}

	// Special case for data.customData.value which has an empty definition.
	if len(schema.Type) == 0 {
		return &goInterface{}, nil
	}

	var typ goType
	var err error

	fullName := parent.qualifiedFieldName(name)

	switch schema.Type[0] { // nolint:exhaustive
	case jsschema.StringType:
		typ = &goPrimitiveType{name: "string"}
	case jsschema.IntegerType:
		typ = &goPrimitiveType{name: "int64"}
	case jsschema.BooleanType:
		typ = &goPrimitiveType{name: "bool"}
	case jsschema.ObjectType:
		// The data.batches.recipes.constraints member of TERCC 1.0.0 has
		// the type "object" but no properties defined. We can't generate
		// a struct type for that so use an interface pointer instead.
		if len(schema.Properties) == 0 {
			typ = &goInterface{}
		} else {
			typ, err = newStruct(parent, name, schema)
		}
	case jsschema.ArrayType:
		// The links slice at the event root should have its own type,
		// e.g. EventLinksV1 instead of []EventLinkV1 like how other slices
		// are represented.
		if fullName == "links" {
			typ, err = newLinkSlice(parent, name, schema.Items)
		} else {
			typ, err = newSlice(parent, name, schema.Items)
		}
	default:
		err = fmt.Errorf("unsupported type: %s", schema.Type[0])
	}
	if err != nil {
		return nil, err
	}

	// Wrap the type in a goEnum if its values are enumerated, but there are
	// couple of exceptions in the schema where a single-value enum has been
	// used instead of a validating regexp or similar. We don't want an enum
	// type to be created for those.
	if fullName != "meta.type" && fullName != "meta.version" && len(schema.Enum) > 0 {
		typ, err = newEnum(parent, name, typ, schema.Enum)
	}

	return typ, err
}
