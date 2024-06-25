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
	_ "embed"
	"sort"
	"strings"
	"text/template"

	jsschema "github.com/lestrrat-go/jsschema"

	"github.com/eiffel-community/eiffelevents-sdk-go/internal/codetemplate"
)

// goStruct represents a struct Go type, including its fields.
type goStruct struct {
	Name              string
	SubTypeNamePrefix string // The prefix of any sub type of this struct, i.e. any nested object.
	JSONField         string // The struct's name in the JSON schema, as given in the parent object.
	Fields            []*goStructField
	required          []string // A list of required properties.
	parent            *goStruct
}

// newEventStruct creates a top-level goStruct, i.e. one that represents an event.
func newEventStruct(subTypeName string, name string, schema *jsschema.Schema) (*goStruct, error) {
	s := &goStruct{
		Name:              fieldTitle(name),
		SubTypeNamePrefix: subTypeName,
		required:          schema.Required,
	}
	return s, s.populate(schema.Properties)
}

// newStruct creates a new goStruct that represents a nested struct type,
// i.e. one nested under a top-level event.
func newStruct(parent *goStruct, name string, schema *jsschema.Schema) (*goStruct, error) {
	s := &goStruct{
		Name:              parent.SubTypeNamePrefix + fieldTitle(name),
		SubTypeNamePrefix: parent.SubTypeNamePrefix + fieldTitle(name),
		JSONField:         name,
		required:          schema.Required,
		parent:            parent,
	}
	return s, s.populate(schema.Properties)
}

//go:embed templates/struct_decl.tmpl
var structDeclTemplate string

func (t *goStruct) declare(ct *codetemplate.OutputFile) error {
	// Declare the struct itself.
	if err := ct.ExpandTemplate(structDeclTemplate, t, template.FuncMap{}); err != nil {
		return err
	}

	// Go through all its members and ask them to declare themselves.
	for _, member := range t.Fields {
		if declarer, ok := member.Type.(goTypeDeclarer); ok {
			if err := declarer.declare(ct); err != nil {
				return err
			}
		}
	}
	return nil
}

// qualifiedFieldName returns the fully qualified JSON name (e.g. meta.version
// or data.name) of a struct field within this struct.
func (t *goStruct) qualifiedFieldName(fieldName string) string {
	qualified := fieldName
	for p := t; p.JSONField != ""; p = p.parent {
		qualified = p.JSONField + "." + qualified
	}
	return qualified
}

// populate loops over the properties described in the schema and adds
// corresponding goStructField values to the struct.
func (t *goStruct) populate(members map[string]*jsschema.Schema) error {
	for name, def := range members {
		typ, err := goTypeFromSchema(t, name, def)
		if err != nil {
			return err
		}

		required := false
		for _, member := range t.required {
			required = required || member == name
		}

		member := &goStructField{
			Name:      fieldTitle(name),
			Type:      typ,
			JSONField: name,
			Required:  required,
		}
		t.Fields = append(t.Fields, member)
	}

	// The range loop aboe will return members in random order so make sure
	// to sort t.Members afterwards so that we can emit the struct members
	// in a consistent order.
	sort.Slice(t.Fields, func(i int, j int) bool {
		return t.Fields[i].Name < t.Fields[j].Name
	})

	return nil
}

func (t *goStruct) String() string {
	return t.Name
}

// goStructField represents a single field within a Go struct.
type goStructField struct {
	Name      string
	Type      goType
	JSONField string
	Required  bool
}

// capitalizationExceptions is a table containing exceptions to the normal
// rule of using the original JSON property with an initial capital letter
// as the name of the Go struct field. This is used to make sure that
// acronyms are all caps, e.g. that we get URI instead of Uri.
var capitalizationExceptions = map[string]string{
	"id":  "ID",
	"sdm": "SDM",
	"uri": "URI",
}

// fieldTitle computes the name of a struct field given the name of the JSON property.
func fieldTitle(propertyName string) string {
	if result, overridden := capitalizationExceptions[propertyName]; overridden {
		return result
	}
	result := initialCapital(propertyName)

	// Also apply overrides when the member ends with the titlecased string,
	// i.e. we want "GroupId" to become "GroupID".
	for orig, repl := range capitalizationExceptions {
		if strings.HasSuffix(result, initialCapital(orig)) {
			result = result[0:len(result)-len(orig)] + repl
		}
	}

	return result
}
