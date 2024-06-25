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
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"io"
	"os"

	eiffelevents "github.com/eiffel-community/eiffelevents-sdk-go"
)

// nolint:gosec
const privKeyPEM = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA4SWVb369VxfJXrkASESVh3RFy+ArrKk8cAc1AEaRks/RhIhY
QQzZDxYtJVYa/K4JmEyxk2jZl16qc9weaSadVL1ZlnukVW5fPejgbgF7+1xyLJ9g
RYTq/hL5vIhpDCR9lxMHZwtZE1cI197n2YFCiYfijBdEzWmh2bdwJzXiZSHC8yry
8+ekBx17Q5c/PhDym8gGIQB/3F1tI8JYQBWfWHDyNPYkvY2fNg+vDObBD3AC56yf
YOsGiIhpdh5z92mDcXo/e5fVsJLL3kJQHGp4a+6/azKPpQRPVYxzzvRox26GGuvY
r2Xj/fDlBpD7E25sVQuxgemzGD/M+7biBeLgEwIDAQABAoIBAQCHQlkAXpfJVtT3
PxVYVTuv4L59uPMEC7fvZaUFwV97X7ZzdKXwjpNoaN4+a/hSjQven1SfRoJSWeD1
MexjJ3uliQvlR+p2GJTHULxj2iht3iAJhsYDfdLfSO8XwKu7S8DXner4kOy2nbcG
WTfYh7s9fJExsFj5PtipP3b1V33nWrwqs0n2+ICejTd0mFRSbQKZL1U513rV+cuu
wT+kofAAqxSfHouWmDQR2PtAxhVk7X4T4dB3JNT5IHqjD5ie9FcVvNGKEy3VfyjI
vVLyL2EymAj1Sx4sJFOOz+ylEHRRV6HKAARgCi0UuNlrUWR4MEZBNbQIih+ikXxq
q6wJtlQhAoGBAONKKifkCgGBQm3bxP7I7w/KV/m+sbJBNVB28lGAyzL32+jqej0Z
MKLZwikA7JSbYDUErw9FAItbLZTEA3kfJGMs63i0+ovY9t8MTrIxI8Fz+qVLgj2Z
Y5sc65E1z1wviz9Ak0YxSoJhLgjzpcYaip6FCAigXetzE6T9rcrPBXM3AoGBAP2W
H+UU5TS4y0c8fikwLPV6Ur9QTjNXWU8aA/m0b1neT/nNO7mFmML5DFbvLWVds6nH
NttlnXhAYiquqGOkTJPB3dILbYMBD1XRAGo+YjxM6it2c7l8QOrRe5e9ZKgV31rr
tTLOvgEM8tTDtINF5Uvq1EOxrTm2lz02z6WMhWAFAoGBALZm5GHS/byrcRYc0oDt
2/w+FFAWmyBEeHa0nk6OH4QtqUvIMIUr2/405z5kwXeZIaIquhp087TiXTgP/gGL
3nXArM/X3WGxopzpkZYrHVi4rKNOb5zjpi3rDZkhJ+IBPaxrNEWWdQcg2gLRFW5g
CnKgrAvQNs8nMNKtynUBoowNAoGABnoxMl7IRAJ8XsNyzYaHf3Wya2SXusP+agDW
HSi4t2jwTgcqAWEiN8i4wfe2ByLPlgSaqBv+W7X5S/HOJ01pD1UiX10fXPtH8v81
rYEObU/ho16RMim0VssnBwc1bP2yCNaAeF3DiK9V/I1LLRc59ih3Z4tAS3sYfd3K
jAX82ikCgYEAp8eeBj2pbtcbWRIlquj68xycx2C06bjR1TeNfbRvnHIBnj7rSe0P
6BleLBsxFxlhnFH9TF84IsfT0+vof/VRTleNU5+em9eN8FUC2xcEfFAMZizXTswY
gESLXERbVkkFCtc1KyZAC6K8/5YRrTcvzzuN2RhRx8prYi8yviqpU8o=
-----END RSA PRIVATE KEY-----
`

func ExampleSigner_Sign() {
	privKeyBlock, _ := pem.Decode([]byte(privKeyPEM))
	privKey, err := x509.ParsePKCS1PrivateKey(privKeyBlock.Bytes)
	if err != nil {
		panic(err.Error())
	}

	// Create an event with some fields set to fixed values so the signed event is kept stable.
	event, err := eiffelevents.NewCompositionDefinedV3()
	if err != nil {
		panic(err.Error())
	}
	event.Meta.ID = "4cd302e1-a636-4c2c-9142-8ec82e39a5f8"
	event.Meta.Time = 1718376599257
	event.Meta.Version = "3.2.0"
	event.Data.Name = "random composition name"

	// Set up the signer and sign the event to a byte slice.
	signer, err := NewKeySigner("CN=test", RS256, privKey)
	if err != nil {
		panic(err.Error())
	}
	eventBytes, err := signer.Sign(event)
	if err != nil {
		panic(err.Error())
	}

	// Pretty-print the event to stdout.
	var prettyEvent bytes.Buffer
	if err := json.Indent(&prettyEvent, eventBytes, "", "  "); err != nil {
		panic(err.Error())
	}
	if _, err := io.Copy(os.Stdout, &prettyEvent); err != nil {
		panic(err.Error())
	}

	// Output: {
	//   "data": {
	//     "name": "random composition name"
	//   },
	//   "links": [],
	//   "meta": {
	//     "id": "4cd302e1-a636-4c2c-9142-8ec82e39a5f8",
	//     "security": {
	//       "authorIdentity": "CN=test",
	//       "integrityProtection": {
	//         "alg": "RS256",
	//         "signature": "LZ0dqqlutep7flccDJuo5I7xxzrKqAeLcZ7aMBHhUGYBi5J+7Rd2dcOEHzN8p+hr6F3dudAb34WYmTgc3/C9dAL0G2RpItnBlPasJkecTJen/AhWfIzjzo8Oji6b4RrSV6LeDC/p7sKr2ocdAEEVV4opHGhavTke9PilPpDndNQCLCssSqRu0ikWkKXwj18lNsCWDcu4phiHkb/BckXJ9ntniDT9evQoBSliOduOa8B8rL0LyRYFC2L++fqJULssCLZ9VoHJZ/FwC1RhFxOofth0kpyhhA+0qGcBMZBX+YqjUmi8PbcuFke4Xlowy+KjQoVWPt2N5Rcg6eSg/716Tw=="
	//       }
	//     },
	//     "time": 1718376599257,
	//     "type": "EiffelCompositionDefinedEvent",
	//     "version": "3.2.0"
	//   }
	// }
}
