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
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

const (
	pubKeyPEM = `
-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEA4SWVb369VxfJXrkASESVh3RFy+ArrKk8cAc1AEaRks/RhIhYQQzZ
DxYtJVYa/K4JmEyxk2jZl16qc9weaSadVL1ZlnukVW5fPejgbgF7+1xyLJ9gRYTq
/hL5vIhpDCR9lxMHZwtZE1cI197n2YFCiYfijBdEzWmh2bdwJzXiZSHC8yry8+ek
Bx17Q5c/PhDym8gGIQB/3F1tI8JYQBWfWHDyNPYkvY2fNg+vDObBD3AC56yfYOsG
iIhpdh5z92mDcXo/e5fVsJLL3kJQHGp4a+6/azKPpQRPVYxzzvRox26GGuvYr2Xj
/fDlBpD7E25sVQuxgemzGD/M+7biBeLgEwIDAQAB
-----END RSA PUBLIC KEY-----`
	signedEvent = `
{
  "data": {
    "name": "random composition name"
  },
  "links": [],
  "meta": {
    "id": "4cd302e1-a636-4c2c-9142-8ec82e39a5f8",
    "security": {
      "authorIdentity": "CN=test",
      "integrityProtection": {
        "alg": "RS256",
        "signature": "LZ0dqqlutep7flccDJuo5I7xxzrKqAeLcZ7aMBHhUGYBi5J+7Rd2dcOEHzN8p+hr6F3dudAb34WYmTgc3/C9dAL0G2RpItnBlPasJkecTJen/AhWfIzjzo8Oji6b4RrSV6LeDC/p7sKr2ocdAEEVV4opHGhavTke9PilPpDndNQCLCssSqRu0ikWkKXwj18lNsCWDcu4phiHkb/BckXJ9ntniDT9evQoBSliOduOa8B8rL0LyRYFC2L++fqJULssCLZ9VoHJZ/FwC1RhFxOofth0kpyhhA+0qGcBMZBX+YqjUmi8PbcuFke4Xlowy+KjQoVWPt2N5Rcg6eSg/716Tw=="
      }
    },
    "time": 1718376599257,
    "type": "EiffelCompositionDefinedEvent",
    "version": "3.2.0"
  }
}`
)

// keyLocator holds sets of public keys belonging to one or more identities.
// It implements the Locator interface and can therefore be used together
// with the Verifier type.
type keyLocator struct {
	keySets []keySet
}

// keySet contains an identity and one or more public keys.
type keySet struct {
	subject *AuthorIdentity
	keys    []crypto.PublicKey
}

// AddKey adds a public key to the set of keys owned by the identified entity.
func (kl *keyLocator) AddKey(identity string, key crypto.PublicKey) error {
	ai, err := NewAuthorIdentity(identity)
	if err != nil {
		return err
	}

	// Is there an existing keySet we can append to?
	for i, k := range kl.keySets {
		if k.subject.Equal(ai) {
			k.keys = append(k.keys, key)
			kl.keySets[i] = k
			return nil
		}
	}

	// No, create a new keySet.
	kl.keySets = []keySet{
		{
			subject: ai,
			keys:    []crypto.PublicKey{key},
		},
	}
	return nil
}

// Locate returns the public keys owned by an entity, or a nil slice
// if no keys are known for that identity.
func (kl *keyLocator) Locate(ctx context.Context, identity *AuthorIdentity) ([]crypto.PublicKey, error) {
	for _, k := range kl.keySets {
		if k.subject.Equal(identity) {
			return k.keys, nil
		}
	}
	return nil, nil
}

func ExampleVerifier_Verify() {
	pubKeyBlock, _ := pem.Decode([]byte(pubKeyPEM))
	pubKey, err := x509.ParsePKCS1PublicKey(pubKeyBlock.Bytes)
	if err != nil {
		panic(err.Error())
	}

	locator := &keyLocator{}
	if err := locator.AddKey("CN=test", pubKey); err != nil {
		panic(err.Error())
	}
	verifier := NewVerifier(locator)

	err = verifier.Verify(context.Background(), []byte(signedEvent))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Signature checked out")

	// Output: Signature checked out
}
