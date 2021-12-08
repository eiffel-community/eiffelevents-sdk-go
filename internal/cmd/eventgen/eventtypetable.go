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

	"github.com/eiffel-community/eiffelevents-sdk-go"
	"github.com/eiffel-community/eiffelevents-sdk-go/internal/codetemplate"
)

type MajorEventVersion struct {
	StructName    string
	LatestVersion string
}

//go:embed templates/eventtypetable.tmpl
var eventTableFileTemplate string

// generateEventTypeTable generates a small Go source file containing
// a variable that maps the major version of each event to a type
// reference to the Go type used to represent the event.
func generateEventTypeTable(schemas map[string][]eventSchemaFile, outputFile string) error {
	table := make(map[string]map[int]MajorEventVersion)
	for _, eventSchemas := range schemas {
		latestVersions := latestMajorVersions(eventSchemas)
		for majorVersion, schema := range latestVersions {
			if table[schema.EventType] == nil {
				table[schema.EventType] = make(map[int]MajorEventVersion)
			}
			table[schema.EventType][int(majorVersion)] = MajorEventVersion{
				eiffelevents.VersionedEventStructName(schema.EventType, schema.Version),
				schema.Version.String(),
			}
		}
	}
	output := codetemplate.New(outputFile)
	if err := output.ExpandTemplate(eventTableFileTemplate, table); err != nil {
		return err
	}
	return output.Close()
}
