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

package eiffelevents

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRoundtrip reads example events from the Eiffel protocol repository,
// deserializes them into Go structs, verifies that we get the expected struct,
// and finally serializes back to JSON and verifies that the result is
// equivalent to what we started with.
func TestRoundtrip(t *testing.T) {
	for _, tc := range eventRoundTripTestTable {
		t.Run(tc.Filename, func(t *testing.T) {
			f, err := os.Open(tc.Filename)
			if err != nil {
				t.Fatal(err.Error())
			}
			defer f.Close()
			input, err := io.ReadAll(f)
			if err != nil {
				t.Fatal(err.Error())
			}

			event, err := UnmarshalAny(input)
			assert.NoError(t, err)
			assert.IsType(t, tc.ExpectedType, event)

			output, err := json.Marshal(event)
			assert.NoError(t, err)
			assert.JSONEq(t, string(input), string(output))
		})
	}
}
