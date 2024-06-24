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
	"context"
	"crypto"
	"crypto/elliptic"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	eiffelevents "github.com/eiffel-community/eiffelevents-sdk-go/editions/lyon"
)

func TestSignAndVerify(t *testing.T) {
	rsaKey := generateRSAKey(t)
	ecdsa256Key := generateECDSAKey(t, elliptic.P256())
	ecdsa384Key := generateECDSAKey(t, elliptic.P384())
	ecdsa521Key := generateECDSAKey(t, elliptic.P521())

	testcases := []struct {
		name             string
		alg              Algorithm
		key              crypto.Signer
		lookupPublicKeys []crypto.PublicKey
		lookupError      error
		expectedError    error
	}{
		{
			name:             "Happy path with RS256",
			alg:              RS256,
			key:              rsaKey,
			lookupPublicKeys: []crypto.PublicKey{rsaKey.Public()},
		},
		{
			name:             "Happy path with RS384",
			alg:              RS384,
			key:              rsaKey,
			lookupPublicKeys: []crypto.PublicKey{rsaKey.Public()},
		},
		{
			name:             "Happy path with RS512",
			alg:              RS512,
			key:              rsaKey,
			lookupPublicKeys: []crypto.PublicKey{rsaKey.Public()},
		},
		{
			name:             "Happy path with ES256",
			alg:              ES256,
			key:              ecdsa256Key,
			lookupPublicKeys: []crypto.PublicKey{ecdsa256Key.Public()},
		},
		{
			name:             "Happy path with ES384",
			alg:              ES384,
			key:              ecdsa384Key,
			lookupPublicKeys: []crypto.PublicKey{ecdsa384Key.Public()},
		},
		{
			name:             "Happy path with ES512",
			alg:              ES512,
			key:              ecdsa521Key,
			lookupPublicKeys: []crypto.PublicKey{ecdsa521Key.Public()},
		},
		{
			name:             "Happy path with PS256",
			alg:              PS256,
			key:              rsaKey,
			lookupPublicKeys: []crypto.PublicKey{rsaKey.Public()},
		},
		{
			name:             "Happy path with PS384",
			alg:              PS384,
			key:              rsaKey,
			lookupPublicKeys: []crypto.PublicKey{rsaKey.Public()},
		},
		{
			name:             "Happy path with PS512",
			alg:              PS512,
			key:              rsaKey,
			lookupPublicKeys: []crypto.PublicKey{rsaKey.Public()},
		},
		{
			name:             "No matching public keys",
			alg:              PS512,
			key:              rsaKey,
			lookupPublicKeys: []crypto.PublicKey{},
			expectedError:    ErrPublicKeyNotFound,
		},
		{
			name:             "Public key lookup error",
			alg:              PS512,
			key:              rsaKey,
			lookupPublicKeys: []crypto.PublicKey{},
			lookupError:      errors.New("random error"),
			expectedError:    ErrPublicKeyLookup,
		},
		{
			name:             "Multiple matching public keys",
			alg:              PS512,
			key:              rsaKey,
			lookupPublicKeys: []crypto.PublicKey{ecdsa521Key.Public(), rsaKey.Public()},
		},
	}
	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%s: %s", tc.alg, tc.name), func(t *testing.T) {
			event, err := eiffelevents.NewCompositionDefined()
			require.NoError(t, err)

			signer, err := NewKeySigner("CN=test", tc.alg, tc.key)
			require.NoError(t, err)
			b, err := signer.Sign(event)
			require.NoError(t, err)

			err = NewVerifier(&constantPublicKeyLocator{tc.lookupPublicKeys, tc.lookupError}).Verify(context.Background(), b)
			if tc.expectedError == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tc.expectedError)
			}
		})
	}
}

type constantPublicKeyLocator struct {
	keys []crypto.PublicKey
	err  error
}

func (cpkl *constantPublicKeyLocator) Locate(ctx context.Context, identity *AuthorIdentity) ([]crypto.PublicKey, error) {
	return cpkl.keys, cpkl.err
}
