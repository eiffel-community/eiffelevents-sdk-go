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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type setFieldLevel1TestStruct struct {
	Level1String string                   `json:"level_1_string"`
	Level1Int    int                      `json:"level_1_int"`
	Level2       setFieldLevel2TestStruct `json:"level_2"`
}

type setFieldLevel2TestStruct struct {
	Level2String string `json:"level_2_string"`
}

func TestSetField(t *testing.T) {
	testcases := []struct {
		name        string
		input       setFieldLevel1TestStruct
		field       string
		value       interface{}
		expected    setFieldLevel1TestStruct
		expectedErr string
	}{
		{
			name:  "Set first level string field",
			input: setFieldLevel1TestStruct{},
			field: "level_1_string",
			value: "new value",
			expected: setFieldLevel1TestStruct{
				Level1String: "new value",
			},
		},
		{
			name:  "Set first level int field",
			input: setFieldLevel1TestStruct{},
			field: "level_1_int",
			value: 123,
			expected: setFieldLevel1TestStruct{
				Level1Int: 123,
			},
		},
		{
			name:        "Setting non-existent field fails",
			input:       setFieldLevel1TestStruct{},
			field:       "ThisFieldDoesNotExist",
			value:       "any value",
			expected:    setFieldLevel1TestStruct{},
			expectedErr: "struct did not contain a field with the JSON name",
		},
		{
			name:  "Set first level struct field",
			input: setFieldLevel1TestStruct{},
			field: "level_2",
			value: setFieldLevel2TestStruct{
				Level2String: "new value",
			},
			expected: setFieldLevel1TestStruct{
				Level2: setFieldLevel2TestStruct{
					Level2String: "new value",
				},
			},
		},
		{
			name:  "Set second level string field",
			input: setFieldLevel1TestStruct{},
			field: "level_2.level_2_string",
			value: "new value",
			expected: setFieldLevel1TestStruct{
				Level2: setFieldLevel2TestStruct{
					Level2String: "new value",
				},
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := setField(reflect.ValueOf(&tc.input), tc.field, tc.value)
			if tc.expectedErr != "" {
				assert.Contains(t, err.Error(), tc.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expected, tc.input)
		})
	}
}
