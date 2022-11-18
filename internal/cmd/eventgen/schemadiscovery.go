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
	"path/filepath"
	"sort"
	"strings"

	"github.com/Masterminds/semver"
)

// findSchemas looks for Eiffel event schemas in the specified
// root directory. Schema files are assumed to have
// the path <root>/<type>/<version>.yml. Returns a collection
// of eventSchemaFile structs grouped per type.
func findSchemas(rootDir string) (map[string][]schemaDefinitionRenderer, error) {
	schemaFiles, err := filepath.Glob(filepath.Join(rootDir, "Eiffel*", "*.yml"))
	if err != nil {
		return nil, err
	}

	// Make sure the files have a well-defined order to make it easier to write tests.
	sort.Strings(schemaFiles)

	result := make(map[string][]schemaDefinitionRenderer)
	for _, schemaFile := range schemaFiles {
		eventType := filepath.Base(filepath.Dir(schemaFile))
		stem := strings.TrimSuffix(filepath.Base(schemaFile), filepath.Ext(filepath.Base(schemaFile)))
		version, err := semver.NewVersion(stem)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", stem, err)
		}

		// Use the filename to determine whether this is an event or some other (struct) type.
		var r schemaDefinitionRenderer
		if strings.HasSuffix(eventType, "Event") {
			r = &eventDefinitionFile{
				definitionFile: definitionFile{schemaFile, eventType, version},
			}
		} else if strings.HasSuffix(eventType, "Link") {
			r = &structDefinitionFile{
				definitionFile: definitionFile{schemaFile, eventType, version},
				templateFile:   linkStructFileTemplate,
			}
		} else {
			r = &structDefinitionFile{
				definitionFile: definitionFile{schemaFile, eventType, version},
				templateFile:   structFileTemplate,
			}
		}
		result[eventType] = append(result[eventType], r)
	}
	return result, nil
}
