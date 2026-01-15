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
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"os"

	"github.com/eiffel-community/eiffelevents-sdk-go"
	"github.com/eiffel-community/eiffelevents-sdk-go/signature"
)

func signCmd(args []string, in io.Reader, out io.Writer) error {
	if len(args) < 3 {
		return fmt.Errorf("%w: not enough arguments for sign command", ErrUsage)
	}
	keyFile := args[0]
	identity := args[1]
	alg := args[2]

	priv, err := loadPrivateKey(keyFile)
	if err != nil {
		return fmt.Errorf("unable to load private key: %w", err)
	}

	signer, err := signature.NewKeySigner(identity, signature.Algorithm(alg), priv)
	if err != nil {
		return fmt.Errorf("unable to create key signer: %w", err)
	}

	decoder := json.NewDecoder(in)
	for {
		var payloadIn json.RawMessage

		err := decoder.Decode(&payloadIn)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("unable to decode input stream: %w", err)
		}

		event, err := eiffelevents.UnmarshalAny(payloadIn)
		if err != nil {
			return fmt.Errorf("unable to unmarshal event: %w", err)
		}

		// We know the event we get from UnmarshalAny implements SigningSubject.
		payloadOut, err := signer.Sign(event.(signature.SigningSubject)) // nolint:forcetypeassert
		if err != nil {
			return fmt.Errorf("unable to sign event: %w", err)
		}

		_, err = out.Write(payloadOut)
		if err != nil {
			return fmt.Errorf("unable to write signed event: %w", err)
		}
	}
	return nil
}

// loadPrivateKey attempts to read and parse a private key file. It supports
// PEM-encoded RSA and ECDSA keys, as well as PKCS#8 (un-encrypted) keys.
func loadPrivateKey(path string) (crypto.PrivateKey, error) {
	pemData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	for block, remaining := pem.Decode(pemData); block != nil; block, remaining = pem.Decode(remaining) {
		switch block.Type {
		case "RSA PRIVATE KEY":
			if k, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
				return k, nil
			}
		case "EC PRIVATE KEY":
			if k, err := x509.ParseECPrivateKey(block.Bytes); err == nil {
				return k, nil
			}
		case "PRIVATE KEY":
			if k, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
				switch pk := k.(type) {
				case *rsa.PrivateKey, *ecdsa.PrivateKey:
					return pk, nil
				default:
					return nil, fmt.Errorf("unsupported key type in PKCS#8: %T", k)
				}
			}
		case "ENCRYPTED PRIVATE KEY":
			return nil, fmt.Errorf("encrypted private keys are not supported")
		}
	}
	return nil, fmt.Errorf("no suitable private key could be found in %s", path)
}
