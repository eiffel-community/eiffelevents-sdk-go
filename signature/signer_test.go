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
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	rooteiffelevents "github.com/eiffel-community/eiffelevents-sdk-go"
)

func generateRSAKey(t *testing.T) crypto.Signer {
	s, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	return s
}

func generateECDSAKey(t *testing.T, c elliptic.Curve) crypto.Signer {
	s, err := ecdsa.GenerateKey(c, rand.Reader)
	require.NoError(t, err)
	return s
}

func TestSigner(t *testing.T) {
	const identity = "CN=test"
	rsaKey := generateRSAKey(t)
	ecdsa256Key := generateECDSAKey(t, elliptic.P256())
	ecdsa384Key := generateECDSAKey(t, elliptic.P384())
	ecdsa521Key := generateECDSAKey(t, elliptic.P521())

	testcases := []struct {
		name          string
		alg           Algorithm
		key           crypto.Signer
		eventFactory  func() (SigningSubject, error)
		expectedError error
	}{
		{
			name:         "Happy path with RS256",
			alg:          RS256,
			key:          rsaKey,
			eventFactory: func() (SigningSubject, error) { return rooteiffelevents.NewCompositionDefinedV3() },
		},
		{
			name:         "Happy path with RS384",
			alg:          RS384,
			key:          rsaKey,
			eventFactory: func() (SigningSubject, error) { return rooteiffelevents.NewCompositionDefinedV3() },
		},
		{
			name:         "Happy path with RS512",
			alg:          RS512,
			key:          rsaKey,
			eventFactory: func() (SigningSubject, error) { return rooteiffelevents.NewCompositionDefinedV3() },
		},
		{
			name:         "Happy path with ES256",
			alg:          ES256,
			key:          ecdsa256Key,
			eventFactory: func() (SigningSubject, error) { return rooteiffelevents.NewCompositionDefinedV3() },
		},
		{
			name:         "Happy path with ES384",
			alg:          ES384,
			key:          ecdsa384Key,
			eventFactory: func() (SigningSubject, error) { return rooteiffelevents.NewCompositionDefinedV3() },
		},
		{
			name:         "Happy path with ES512",
			alg:          ES512,
			key:          ecdsa521Key,
			eventFactory: func() (SigningSubject, error) { return rooteiffelevents.NewCompositionDefinedV3() },
		},
		{
			name:         "Happy path with PS256",
			alg:          PS256,
			key:          rsaKey,
			eventFactory: func() (SigningSubject, error) { return rooteiffelevents.NewCompositionDefinedV3() },
		},
		{
			name:         "Happy path with PS384",
			alg:          PS384,
			key:          rsaKey,
			eventFactory: func() (SigningSubject, error) { return rooteiffelevents.NewCompositionDefinedV3() },
		},
		{
			name:         "Happy path with PS512",
			alg:          PS512,
			key:          rsaKey,
			eventFactory: func() (SigningSubject, error) { return rooteiffelevents.NewCompositionDefinedV3() },
		},
		{
			name:          "Algorithm and key mismatch",
			alg:           RS256,
			key:           ecdsa256Key,
			eventFactory:  func() (SigningSubject, error) { return rooteiffelevents.NewCompositionDefinedV3() },
			expectedError: ErrSigningFailed,
		},
		{
			name:          "Event is too old to support signing",
			alg:           PS512,
			key:           rsaKey,
			eventFactory:  func() (SigningSubject, error) { return rooteiffelevents.NewCompositionDefinedV2() },
			expectedError: ErrSigningUnavailable,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			event, err := tc.eventFactory()
			require.NoError(t, err)

			signer, err := NewKeySigner(identity, tc.alg, tc.key)
			require.NoError(t, err)
			b, err := signer.Sign(event)
			if tc.expectedError == nil {
				require.NoError(t, err)
				assert.Equal(t, identity, gjson.GetBytes(b, authorIdentityField).String())
				assert.Equal(t, string(tc.alg), gjson.GetBytes(b, algorithmField).String())
				// We have other tests that actually try to do something with the signature.
				// We'll just check that it's a non-empty string.
				assert.NotEmpty(t, gjson.GetBytes(b, signatureField).String())
			} else {
				require.ErrorIs(t, err, tc.expectedError)
			}
		})
	}
}
