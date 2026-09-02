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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	eiffelevents "github.com/eiffel-community/eiffelevents-sdk-go/editions/arica"
)

func TestPackageURLValidator(t *testing.T) {
	testcases := []struct {
		name          string
		event         []byte
		expectedError string
	}{
		{
			name: "Non-ArtC event is ignored",
			event: []byte(`{
				"meta": {"type": "EiffelCompositionDefinedEvent"},
				"data": {"identity": "not a purl"}
			}`),
		},
		{
			name: "ArtC event with valid purl",
			event: []byte(`{
				"meta": {"type": "EiffelArtifactCreatedEvent"},
				"data": {"identity": "pkg:generic/example"}
			}`),
		},
		{
			name: "ArtC event with missing identity",
			event: []byte(`{
				"meta": {"type": "EiffelArtifactCreatedEvent"},
				"data": {}
			}`),
			expectedError: "missing or invalid contents of data.identity",
		},
		{
			name: "ArtC event with malformed purl",
			event: []byte(`{
				"meta": {"type": "EiffelArtifactCreatedEvent"},
				"data": {"identity": "definitely-not-a-purl"}
			}`),
			expectedError: "invalid package URL in data.identity",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			validator := NewPackageURLValidator()
			err := validator.Validate(t.Context(), tc.event)
			if tc.expectedError == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.expectedError)
		})
	}
}

func TestDefaultSetValidatesArtifactCreatedIdentityPURL(t *testing.T) {
	event, err := eiffelevents.NewArtifactCreated()
	require.NoError(t, err)
	event.Data.Identity = "pkg:"

	err = DefaultSet().Validate(context.Background(), []byte(event.String()))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid package URL in data.identity")
}
