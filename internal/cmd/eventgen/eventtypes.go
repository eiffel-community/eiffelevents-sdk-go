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
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver"
	jsschema "github.com/lestrrat-go/jsschema"

	"github.com/eiffel-community/eiffelevents-sdk-go"
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

// generateEventTypes generates Go struct declarations (and all types
// referenced in those structs) for the latest event version within
// each major version of each event type.
func generateEventTypes(eventSchemas map[string][]eventSchemaFile, outputDir string) error {
	for _, schemas := range eventSchemas {
		for majorVersion, schema := range latestMajorVersions(schemas) {
			schemaFile, err := os.Open(schema.Filename)
			if err != nil {
				return err
			}
			defer schemaFile.Close()

			outputFile := filepath.Join(outputDir, fmt.Sprintf("%sV%d.go", schema.EventType, majorVersion))
			if err = generateEventFile(schema.EventType, schema.Version, schemaFile, outputFile); err != nil {
				return err
			}
		}
	}
	return nil
}

// latestMajorVersions inspects a list of schemas for a single event type and
// maps each encountered major version to the schemadiscovery.EventSchemaFile
// that represents the most recent minor.patch version within that major
// version.
func latestMajorVersions(schemas []eventSchemaFile) map[int64]eventSchemaFile {
	majorVersions := map[int64]eventSchemaFile{}
	for _, schema := range schemas {
		if current, exists := majorVersions[schema.Version.Major()]; !exists || current.Version.LessThan(schema.Version) {
			majorVersions[schema.Version.Major()] = schema
		}
	}
	return majorVersions
}

// goTypeFromSchema returns a goType that represents a node in a JSON schema.
func goTypeFromSchema(parent *goStruct, name string, schema *jsschema.Schema) (goType, error) {
	// Special case for data.customData.value which has an empty definition.
	if len(schema.Type) == 0 {
		return &goInterface{}, nil
	}

	var typ goType
	var err error

	switch schema.Type[0] {
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
		typ, err = newSlice(parent, name, schema.Items)
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
	fullName := parent.qualifiedFieldName(name)
	if fullName != "meta.type" && fullName != "meta.version" && len(schema.Enum) > 0 {
		typ, err = newEnum(parent, name, typ, schema.Enum)
	}

	return typ, err
}

//go:embed templates/eventfile.tmpl
var eventFileTemplate string

// eventTypeAbbrevMap maps event type name to their abbreviations. For now
// the contents of this map is created by running
//
//     sed -n 's/^# \([^ ]*\) (\([A-Za-z]*\))/"\1": "\2",/p' eiffel-vocubulary/*.md
//
// in the Eiffel protocol repository but when
// https://github.com/eiffel-community/eiffel/issues/282 has been addressed
// we can hopefully do it in a better way.
var eventTypeAbbrevMap = map[string]string{
	"EiffelActivityCanceledEvent":                     "ActC",
	"EiffelActivityFinishedEvent":                     "ActF",
	"EiffelActivityStartedEvent":                      "ActS",
	"EiffelActivityTriggeredEvent":                    "ActT",
	"EiffelAnnouncementPublishedEvent":                "AnnP",
	"EiffelArtifactCreatedEvent":                      "ArtC",
	"EiffelArtifactPublishedEvent":                    "ArtP",
	"EiffelArtifactReusedEvent":                       "ArtR",
	"EiffelCompositionDefinedEvent":                   "CD",
	"EiffelConfidenceLevelModifiedEvent":              "CLM",
	"EiffelEnvironmentDefinedEvent":                   "ED",
	"EiffelFlowContextDefinedEvent":                   "FCD",
	"EiffelIssueDefinedEvent":                         "ID",
	"EiffelIssueVerifiedEvent":                        "IV",
	"EiffelSourceChangeCreatedEvent":                  "SCC",
	"EiffelSourceChangeSubmittedEvent":                "SCS",
	"EiffelTestCaseCanceledEvent":                     "TCC",
	"EiffelTestCaseFinishedEvent":                     "TCF",
	"EiffelTestCaseStartedEvent":                      "TCS",
	"EiffelTestCaseTriggeredEvent":                    "TCT",
	"EiffelTestExecutionRecipeCollectionCreatedEvent": "TERCC",
	"EiffelTestSuiteFinishedEvent":                    "TSF",
	"EiffelTestSuiteStartedEvent":                     "TSS",
}

// generateEventFile generates a Go source file with the Go struct and
// associated types needed to represent a particular major version of
// an Eiffel event, given its JSON schema.
func generateEventFile(eventType string, version *semver.Version, schema io.Reader, outputFile string) error {
	s, err := jsschema.Read(schema)
	if err != nil {
		return err
	}

	eventTypeAbbrev := eventTypeAbbrevMap[eventType]
	if eventTypeAbbrev == "" {
		return fmt.Errorf("the event type %q isn't supported (no known abbreviation)", eventType)
	}

	// Gather some metadata about the event type. This struct is later
	// supplied to the template that generates the event source file.
	eventMeta := struct {
		EventType         string // The name of the event type, e.g. EiffelActivityTriggeredEvent.
		EventTypeAbbrev   string // The abbreviated event type name, e.g. ActT.
		StructName        string // The name of the struct that represents the event type.
		SubTypeNamePrefix string // The prefix that any subtypes of the event type struct gets to their names.
		MajorVersion      int64  // The event type's major version.
	}{
		EventType:         eventType,
		EventTypeAbbrev:   eventTypeAbbrev,
		StructName:        eiffelevents.VersionedEventStructName(eventType, version),
		SubTypeNamePrefix: fmt.Sprintf("%sV%d", eventTypeAbbrev, version.Major()),
		MajorVersion:      version.Major(),
	}

	rootStruct, err := newEventStruct(eventMeta.SubTypeNamePrefix, eventMeta.StructName, s)
	if err != nil {
		return err
	}

	ct := codetemplate.New(outputFile)
	if err := ct.ExpandTemplate(eventFileTemplate, eventMeta); err != nil {
		return err
	}
	if err := rootStruct.declare(ct); err != nil {
		return err
	}
	return ct.Close()
}
