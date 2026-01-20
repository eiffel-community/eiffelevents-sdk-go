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
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFSPublicKeyLocator_Locate(t *testing.T) {
	// Rather than relying on generating proper PEM files with keys that we have
	// to compare in cumbersome ways, we replace the PEM parsing function with
	// one that just returns each line of the data blob as a mockPublicKey,
	// a custom type based on string that implements crypto.PublicKey.
	testcases := []struct {
		name     string
		files    map[string]string // filename => contents
		lookup   string
		expected []crypto.PublicKey
	}{
		{
			name: "Single file with matching identity",
			files: map[string]string{
				"CN=test.pem": "A",
			},
			lookup:   "CN=test",
			expected: []crypto.PublicKey{mockPublicKey("A")},
		},
		{
			name: "Single file with equivalent identity",
			files: map[string]string{
				"CN=test.pem": "A",
			},
			lookup:   "cn=test",
			expected: []crypto.PublicKey{mockPublicKey("A")},
		},
		{
			name: "Two files with equivalent identities and multiple keys",
			files: map[string]string{
				"CN=test.pem": "A\nB",
				"cn=test.pem": "C\nD",
			},
			lookup:   "CN=test",
			expected: []crypto.PublicKey{mockPublicKey("A"), mockPublicKey("B"), mockPublicKey("C"), mockPublicKey("D")},
		},
		{
			name: "Two files with just one matching",
			files: map[string]string{
				"CN=test.pem":  "A",
				"CN=test2.pem": "B",
			},
			lookup:   "CN=test",
			expected: []crypto.PublicKey{mockPublicKey("A")},
		},
		{
			name: "Ignores files that don't end with .pem",
			files: map[string]string{
				"CN=test": "A",
				"README":  "B",
			},
			lookup:   "CN=test",
			expected: []crypto.PublicKey{},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			tempDir := t.TempDir()
			pkl := NewFSPublicKeyLocator(FSPublicKeyLocatorConfig{
				KeyDirectory: tempDir,
			})
			pkl.keyLoader = mockPublicKeysFromLines

			// Create the files declared in the testcase.
			for filename, contents := range tc.files {
				require.NoError(t, os.WriteFile(filepath.Join(tempDir, filename), []byte(contents), 0600))
			}

			result, err := pkl.Locate(t.Context(), mustParseAuthorIdentity(t, tc.lookup))
			require.NoError(t, err)

			assert.ElementsMatch(t, tc.expected, result)
		})
	}
}

func mustParseAuthorIdentity(t *testing.T, s string) *AuthorIdentity {
	ai, err := NewAuthorIdentity(s)
	require.NoError(t, err)
	return ai
}

// mockPublicKeysFromLines produces a slice of mockPublicKey values,
// one for each line in the input data.
func mockPublicKeysFromLines(data []byte) ([]crypto.PublicKey, error) {
	var result []crypto.PublicKey
	for _, line := range strings.Split(string(data), "\n") {
		result = append(result, mockPublicKey(line))
	}
	return result, nil
}

// mockPublicKey implements crypto.PublicKey and simply holds an arbitrary
// string value that it can compare to that of another struct of the same type.
type mockPublicKey string

func (mpk *mockPublicKey) Equal(x crypto.PublicKey) bool {
	other, ok := x.(*mockPublicKey)
	if !ok {
		return false
	}
	return mpk == other
}

func TestPublicKeysInPEMFile(t *testing.T) {
	testDataDir := filepath.Join("testdata", t.Name())
	testcases := []struct {
		name             string
		pemFile          string
		expectedKeyTypes []crypto.PublicKey
	}{
		{
			name:    "RSA key",
			pemFile: "rsa.pem",
			expectedKeyTypes: []crypto.PublicKey{
				&rsa.PublicKey{},
			},
		},
		{
			name:    "ECDSA key",
			pemFile: "ecdsa.pem",
			expectedKeyTypes: []crypto.PublicKey{
				&ecdsa.PublicKey{},
			},
		},
		{
			name:    "RSA and ECDSA keys",
			pemFile: "rsa_ecdsa.pem",
			expectedKeyTypes: []crypto.PublicKey{
				&rsa.PublicKey{},
				&ecdsa.PublicKey{},
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			pemData, err := os.ReadFile(filepath.Join(testDataDir, tc.pemFile))
			require.NoError(t, err)

			keys, err := publicKeysFromPEMData(pemData)
			require.NoError(t, err)
			require.Len(t, keys, len(tc.expectedKeyTypes))
			for i := range tc.expectedKeyTypes {
				assert.IsType(t, tc.expectedKeyTypes[i], keys[i])
			}
		})
	}
}
