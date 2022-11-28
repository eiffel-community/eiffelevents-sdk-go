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
	"text/template"

	"github.com/Masterminds/semver"
	jsschema "github.com/lestrrat-go/jsschema"

	"github.com/eiffel-community/eiffelevents-sdk-go"
	"github.com/eiffel-community/eiffelevents-sdk-go/internal/codetemplate"
)

// schemaDefinitionRenderer renders a schema definition to a Go source file.
type schemaDefinitionRenderer interface {
	TypeName() string
	Filename() string
	Render(schema io.Reader, outputFile string) error
	Version() *semver.Version
}

// definitionFile holds common fields and methods to support rendering
// any kind of schema definition to a Go source file. The struct can
// be embedded in a new struct that implmements the Render method from
// schemaDefinitionRenderer to get a working renderer.
type definitionFile struct {
	filename string
	typeName string
	version  *semver.Version
}

func (edf *definitionFile) TypeName() string {
	return edf.typeName
}

func (edf *definitionFile) Filename() string {
	return edf.filename
}

func (edf *definitionFile) Version() *semver.Version {
	return edf.version
}

//go:embed templates/eventfile.tmpl
var eventFileTemplate string

// eventTypeAbbrevMap maps event type name to their abbreviations. For now
// the contents of this map is created by running
//
//	sed -n 's/^# \([^ ]*\) (\([A-Za-z]*\))/"\1": "\2",/p' eiffel-vocubulary/*.md
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

// eventDefinitionFile holds information about a particular version of
// an Eiffel event and where the event's schema can be located. It also
// implements schemaDefinitionRenderer.
type eventDefinitionFile struct {
	definitionFile
}

func (edf *eventDefinitionFile) Render(schema io.Reader, outputFile string) error {
	s, err := jsschema.Read(schema)
	if err != nil {
		return err
	}

	eventTypeAbbrev := eventTypeAbbrevMap[edf.typeName]
	if eventTypeAbbrev == "" {
		return fmt.Errorf("the event type %q isn't supported (no known abbreviation)", edf.typeName)
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
		EventType:         edf.typeName,
		EventTypeAbbrev:   eventTypeAbbrev,
		StructName:        eiffelevents.VersionedStructName(edf.typeName, edf.version),
		SubTypeNamePrefix: fmt.Sprintf("%sV%d", eventTypeAbbrev, edf.version.Major()),
		MajorVersion:      edf.version.Major(),
	}

	rootStruct, err := newEventStruct(eventMeta.SubTypeNamePrefix, eventMeta.StructName, s)
	if err != nil {
		return err
	}

	ct := codetemplate.New(outputFile)
	funcs := template.FuncMap{
		// The FieldType function allows the template to look up the declared type of any struct member.
		"FieldType": func(name string) (string, error) {
			for _, f := range rootStruct.Fields {
				if f.JSONField == name {
					return f.Type.String(), nil
				}
			}
			return "", fmt.Errorf("no field %q found in struct", name)
		},
	}
	if err := ct.ExpandTemplate(eventFileTemplate, eventMeta, funcs); err != nil {
		return err
	}
	if err := rootStruct.declare(ct); err != nil {
		return err
	}
	return ct.Close()
}

//go:embed templates/linkfile.tmpl
var linkStructFileTemplate string

//go:embed templates/structfile.tmpl
var structFileTemplate string

// structDefinitionFile holds information about a particular version of
// a struct and where its schema can be located. It also implements
// schemaDefinitionRenderer.
type structDefinitionFile struct {
	definitionFile
	templateFile string
}

func (sdf *structDefinitionFile) Render(schema io.Reader, outputFile string) error {
	s, err := jsschema.Read(schema)
	if err != nil {
		return err
	}

	// Gather some metadata about the event type. This struct is later
	// supplied to the template that generates the event source file.
	structName := eiffelevents.VersionedStructName(sdf.typeName, sdf.version)
	typeMeta := struct {
		TypeName          string // The unversioned name of the schema type, e.g. EiffelEventLink.
		BaseStructName    string // The unversioned name of the Go struct that represents the type, e.g. EventLink.
		StructName        string // The versioned name of the Go struct that represents the type, e.g. EventLinkV1.
		SubTypeNamePrefix string // The prefix that any subtypes of the event type struct gets to their names.
		MajorVersion      int64  // The type's major version.
	}{
		TypeName:          sdf.typeName,
		BaseStructName:    eiffelevents.StructName(sdf.typeName, sdf.version),
		StructName:        structName,
		SubTypeNamePrefix: structName,
		MajorVersion:      sdf.version.Major(),
	}

	rootStruct, err := newEventStruct(typeMeta.SubTypeNamePrefix, typeMeta.StructName, s)
	if err != nil {
		return err
	}

	ct := codetemplate.New(outputFile)
	if err := ct.ExpandTemplate(sdf.templateFile, typeMeta, template.FuncMap{}); err != nil {
		return err
	}
	if err := rootStruct.declare(ct); err != nil {
		return err
	}
	return ct.Close()
}
