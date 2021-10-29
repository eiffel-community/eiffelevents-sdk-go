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
	"os"
	"path/filepath"
	"testing"

	"github.com/Masterminds/semver"
	"github.com/stretchr/testify/assert"
)

func TestFindSchemas(t *testing.T) {
	// Ideally we would've wanted to use unique tempdirs for each testcase
	// but since the returned eventSchemaFile values contain absolute paths
	// it would've been impossible to predict the expected paths listed in
	// the testcases. Instead use the testcase name an extra per-testcase directory level
	// that we maintain ourselves.
	tempDir := t.TempDir()

	testcases := []struct {
		name          string
		filenames     []string
		expected      map[string][]eventSchemaFile
		expectedError error
	}{
		{
			name:      "Empty schema directory",
			filenames: []string{},
			expected:  map[string][]eventSchemaFile{},
		},
		{
			name: "Ignores non-Eiffel directory",
			filenames: []string{
				"SomeOtherDirectory/1.0.0.json",
			},
			expected: map[string][]eventSchemaFile{},
		},
		{
			name: "Ignores non-JSON file",
			filenames: []string{
				"EiffelCompositionDefinedEvent/this-is-not-json.txt",
			},
			expected: map[string][]eventSchemaFile{},
		},
		{
			name: "Error on bad version",
			filenames: []string{
				"EiffelCompositionDefinedEvent/this-is-not-a-valid-version.json",
			},
			expectedError: semver.ErrInvalidSemVer,
		},
		{
			name: "Multiple events with multiple versions",
			filenames: []string{
				"EiffelArtifactCreatedEvent/1.0.0.json",
				"EiffelArtifactCreatedEvent/2.0.0.json",
				"EiffelCompositionDefinedEvent/3.0.0.json",
				"EiffelCompositionDefinedEvent/4.0.0.json",
			},
			expected: map[string][]eventSchemaFile{
				"EiffelArtifactCreatedEvent": {
					{
						filepath.Join(tempDir, "Multiple events with multiple versions", "EiffelArtifactCreatedEvent/1.0.0.json"),
						"EiffelArtifactCreatedEvent",
						semver.MustParse("1.0.0"),
					},
					{
						filepath.Join(tempDir, "Multiple events with multiple versions", "EiffelArtifactCreatedEvent/2.0.0.json"),
						"EiffelArtifactCreatedEvent",
						semver.MustParse("2.0.0"),
					},
				},
				"EiffelCompositionDefinedEvent": {
					{
						filepath.Join(tempDir, "Multiple events with multiple versions", "EiffelCompositionDefinedEvent/3.0.0.json"),
						"EiffelCompositionDefinedEvent",
						semver.MustParse("3.0.0"),
					},
					{
						filepath.Join(tempDir, "Multiple events with multiple versions", "EiffelCompositionDefinedEvent/4.0.0.json"),
						"EiffelCompositionDefinedEvent",
						semver.MustParse("4.0.0"),
					},
				},
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, filename := range tc.filenames {
				path := filepath.Join(tempDir, tc.name, filename)
				err := os.MkdirAll(filepath.Dir(path), 0700)
				if err != nil {
					t.Fatal(err)
				}
				_, err = os.Create(path)
				if err != nil {
					t.Fatal(err)
				}
			}

			schemas, err := findSchemas(filepath.Join(tempDir, tc.name))
			if tc.expectedError != nil {
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, schemas)
			}
		})
	}
}
