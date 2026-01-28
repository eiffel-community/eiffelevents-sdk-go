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

package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/eiffel-community/eiffelevents-sdk-go"
)

// TestSignAndVerify is a simple smoke test for the eiffelsignature subcommands
// that generates a keypair and an event, signs the event with the "sign" command
// and verifies it with the "verify" command.
func TestSignAndVerify(t *testing.T) {
	dn := "CN=test"
	privateKeyPath := filepath.Join(t.TempDir(), "private.pem")
	publicKeyDir := t.TempDir()

	createKeypairFiles(t, privateKeyPath, filepath.Join(publicKeyDir, dn+".pem"))

	event, err := eiffelevents.NewCompositionDefinedV3()
	require.NoError(t, err)
	event.Data.Name = "my-composition"

	var signedEvent bytes.Buffer
	require.NoError(t, signCmd([]string{privateKeyPath, dn, "ES512"}, strings.NewReader(event.String()), &signedEvent))
	require.NoError(t, verifyCmd(t.Context(), []string{publicKeyDir}, &signedEvent))
}

func createKeypairFiles(t *testing.T, privateKeyPath string, publicKeyFilePath string) {
	// Generate a new ECDSA private key.
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	require.NoError(t, err)

	// Encode the private to DER and marshal that to a PEM file.
	privBytes, err := x509.MarshalECPrivateKey(privateKey)
	require.NoError(t, err)

	privateKeyFile, err := os.Create(privateKeyPath)
	require.NoError(t, err)
	defer privateKeyFile.Close()
	err = pem.Encode(privateKeyFile, &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privBytes,
	})
	require.NoError(t, err)

	// Extract the public part of the key, encode to DER and marshal to a PEM file.
	pubBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	require.NoError(t, err)

	publicKeyFile, err := os.Create(publicKeyFilePath)
	require.NoError(t, err)
	defer publicKeyFile.Close()
	err = pem.Encode(publicKeyFile, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	})
	require.NoError(t, err)
}
