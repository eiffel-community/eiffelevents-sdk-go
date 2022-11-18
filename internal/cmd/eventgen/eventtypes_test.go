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
	"testing"

	"github.com/Masterminds/semver"
	"github.com/stretchr/testify/assert"
)

func TestLatestMajorVersions(t *testing.T) {
	testcases := []struct {
		name     string
		input    []schemaDefinitionRenderer
		expected map[int64]schemaDefinitionRenderer
	}{
		{
			name:     "Empty input map",
			input:    []schemaDefinitionRenderer{},
			expected: map[int64]schemaDefinitionRenderer{},
		},
		{
			name: "Multiple major versions",
			input: []schemaDefinitionRenderer{
				&eventDefinitionFile{
					definitionFile: definitionFile{"/path/to/EiffelActivityStartedEvent/1.0.0.json", "EiffelActivityStartedEvent", semver.MustParse("1.0.0")},
				},
				&eventDefinitionFile{
					definitionFile: definitionFile{"/path/to/EiffelActivityStartedEvent/1.1.0.json", "EiffelActivityStartedEvent", semver.MustParse("1.1.0")},
				},
				&eventDefinitionFile{
					definitionFile: definitionFile{"/path/to/EiffelActivityStartedEvent/2.0.0.json", "EiffelActivityStartedEvent", semver.MustParse("2.0.0")},
				},
				&eventDefinitionFile{
					definitionFile: definitionFile{"/path/to/EiffelActivityStartedEvent/3.0.0.json", "EiffelActivityStartedEvent", semver.MustParse("3.0.0")},
				},
				&eventDefinitionFile{
					definitionFile: definitionFile{"/path/to/EiffelActivityStartedEvent/4.0.0.json", "EiffelActivityStartedEvent", semver.MustParse("4.0.0")},
				},
				&eventDefinitionFile{
					definitionFile: definitionFile{"/path/to/EiffelActivityStartedEvent/4.1.0.json", "EiffelActivityStartedEvent", semver.MustParse("4.1.0")},
				},
				&eventDefinitionFile{
					definitionFile: definitionFile{"/path/to/EiffelActivityStartedEvent/4.2.0.json", "EiffelActivityStartedEvent", semver.MustParse("4.2.0")},
				},
			},
			expected: map[int64]schemaDefinitionRenderer{
				1: &eventDefinitionFile{
					definitionFile: definitionFile{"/path/to/EiffelActivityStartedEvent/1.1.0.json", "EiffelActivityStartedEvent", semver.MustParse("1.1.0")},
				},
				2: &eventDefinitionFile{
					definitionFile: definitionFile{"/path/to/EiffelActivityStartedEvent/2.0.0.json", "EiffelActivityStartedEvent", semver.MustParse("2.0.0")},
				},
				3: &eventDefinitionFile{
					definitionFile: definitionFile{"/path/to/EiffelActivityStartedEvent/3.0.0.json", "EiffelActivityStartedEvent", semver.MustParse("3.0.0")},
				},
				4: &eventDefinitionFile{
					definitionFile: definitionFile{"/path/to/EiffelActivityStartedEvent/4.2.0.json", "EiffelActivityStartedEvent", semver.MustParse("4.2.0")},
				},
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := latestMajorVersions(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
