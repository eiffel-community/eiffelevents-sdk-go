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
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetaSchemaLocator(t *testing.T) {
	testcases := []struct {
		name          string
		schemaURI     string
		httpGetter    *fakeHTTPGetter
		expectedBody  string // If empty, require a nil reader to be returned.
		errorContains string // If non-empty, require that returned error's string representation contains this.
	}{
		{
			name:          "Empty schema URI should return inconclusively",
			schemaURI:     "",
			expectedBody:  "",
			errorContains: "",
		},
		{
			name:          "Non-HTTP(S) scheme should return inconclusively",
			schemaURI:     "file:///tmp/schema.json",
			expectedBody:  "",
			errorContains: "",
		},
		{
			name:      "Happy path",
			schemaURI: "http:///example.com/schema.json",
			httpGetter: &fakeHTTPGetter{
				statusCode: http.StatusOK,
				body:       "event schema",
			},
			expectedBody:  "event schema",
			errorContains: "",
		},
		{
			name:      "HTTP request returns error",
			schemaURI: "http:///example.com/schema.json",
			httpGetter: &fakeHTTPGetter{
				statusCode: http.StatusNotFound,
				body:       "404 Not Found",
			},
			expectedBody:  "",
			errorContains: "returned status 404",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			loc := NewMetaSchemaLocator(tc.httpGetter)
			body, err := loc.GetSchema(context.Background(), "EiffelCompositionDefinedEvent", "1.0.0", tc.schemaURI)
			if tc.errorContains != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errorContains)
			} else {
				require.NoError(t, err)
			}
			if tc.expectedBody != "" {
				b, err := io.ReadAll(body)
				require.NoError(t, err)
				assert.Equal(t, tc.expectedBody, string(b))
			} else {
				assert.Nil(t, body)
			}
		})
	}
}

type fakeHTTPGetter struct {
	statusCode int
	body       string
}

func (fhg *fakeHTTPGetter) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: fhg.statusCode,
		Body:       io.NopCloser(strings.NewReader(fhg.body)),
	}, nil
}
