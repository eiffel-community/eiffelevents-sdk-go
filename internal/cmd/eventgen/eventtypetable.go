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
	"strings"
	"text/template"

	"github.com/Masterminds/semver"
	"github.com/eiffel-community/eiffelevents-sdk-go"
	"github.com/eiffel-community/eiffelevents-sdk-go/internal/codetemplate"
)

type MajorEventVersion struct {
	StructName    string
	LatestVersion string
	latest        *semver.Version
}

//go:embed templates/eventtypetable.tmpl
var eventTableFileTemplate string

// generateEventTypeTable generates a small Go source file containing
// a variable that maps the major version of each event to a type
// reference to the Go type used to represent the event.
func generateEventTypeTable(schemas map[string][]schemaDefinitionRenderer, outputFile string) error {
	table := make(map[string]map[int]MajorEventVersion)
	for _, eventSchemas := range schemas {
		latestVersions := significantVersions(eventSchemas)
		for _, schema := range latestVersions {
			if !strings.HasSuffix(schema.TypeName(), "Event") {
				continue
			}
			if table[schema.TypeName()] == nil {
				table[schema.TypeName()] = make(map[int]MajorEventVersion)
			}

			// For non-experimental event versions this conditional is unnecessary since significantVersions()
			// has already weeded out non-latest versions, but for experimental versions we need to make sure
			// we fill the table with the latest one (and the latest one only).
			major := int(schema.Version().Major())
			if current, exists := table[schema.TypeName()][major]; !exists || current.latest.LessThan(schema.Version()) {
				table[schema.TypeName()][major] = MajorEventVersion{
					eiffelevents.VersionedStructName(schema.TypeName(), schema.Version()),
					schema.Version().String(),
					schema.Version(),
				}
			}
		}
	}
	output := codetemplate.New(outputFile)
	if err := output.ExpandTemplate(eventTableFileTemplate, table, template.FuncMap{}); err != nil {
		return err
	}
	return output.Close()
}
