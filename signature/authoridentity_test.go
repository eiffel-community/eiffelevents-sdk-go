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

package signature

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthorIdentity_UnmarshalText(t *testing.T) {
	var data struct {
		Identity *AuthorIdentity `json:"identity"`
	}
	require.NoError(t, json.Unmarshal([]byte(`{"identity": "CN=foo,DC=example,DC=com"}`), &data))
	assert.Equal(t, "CN=foo,DC=example,DC=com", data.Identity.String())
}
