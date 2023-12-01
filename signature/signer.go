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
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gowebpki/jcs"
	"github.com/tidwall/sjson"

	"github.com/eiffel-community/eiffelevents-sdk-go"
)

// Algorithm describes the set of algorithms used when signing or verifying a payload.
// It includes both the signing algorithm and the hash algorithm. The latter is
// important since it's the hash that's signed rather than the payload itself.
type Algorithm string

const (
	RS256 = Algorithm(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_RS256) // RSASSA-PKCS1-v1_5 using SHA-256
	RS384 = Algorithm(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_RS384) // RSASSA-PKCS1-v1_5 using SHA-384
	RS512 = Algorithm(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_RS512) // RSASSA-PKCS1-v1_5 using SHA-512
	ES256 = Algorithm(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_ES256) // ECDSA using P-256 and SHA-256
	ES384 = Algorithm(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_ES384) // ECDSA using P-384 and SHA-384
	ES512 = Algorithm(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_ES512) // ECDSA using P-521 and SHA-512
	PS256 = Algorithm(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_PS256) // RSASSA-PSS using SHA-256 and MGF1 with SHA-256
	PS384 = Algorithm(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_PS384) // RSASSA-PSS using SHA-384 and MGF1 with SHA-384
	PS512 = Algorithm(eiffelevents.MetaV3SecurityIntegrityProtectionAlg_PS512) // RSASSA-PSS using SHA-512 and MGF1 with SHA-512
)

const (
	authorIdentityField = "meta.security.authorIdentity"
	algorithmField      = "meta.security.integrityProtection.alg"
	signatureField      = "meta.security.integrityProtection.signature"
)

// Signer signs Eiffel event payloads in the standard way.
type Signer struct {
	identity   string
	alg        Algorithm
	pk         crypto.PrivateKey
	hashFunc   func([]byte) []byte
	signFunc   func(crypto.PrivateKey, crypto.Hash, []byte) ([]byte, error)
	signerOpts crypto.SignerOpts
}

// NewKeySigner initializes a Signer with a private key and an identity.
func NewKeySigner(identity string, alg Algorithm, pk crypto.PrivateKey) (*Signer, error) {
	s := &Signer{
		identity: identity,
		alg:      alg,
		pk:       pk,
	}
	switch s.alg {
	case RS256:
		s.hashFunc = hashSHA256
		s.signerOpts = crypto.SHA256
		s.signFunc = signPKCS1v15
	case ES256:
		s.hashFunc = hashSHA256
		s.signerOpts = crypto.SHA256
		s.signFunc = signECDSA
	case PS256:
		s.hashFunc = hashSHA256
		s.signerOpts = crypto.SHA256
		s.signFunc = signPSS
	case RS384:
		s.hashFunc = hashSHA384
		s.signerOpts = crypto.SHA384
		s.signFunc = signPKCS1v15
	case ES384:
		s.hashFunc = hashSHA384
		s.signerOpts = crypto.SHA384
		s.signFunc = signECDSA
	case PS384:
		s.hashFunc = hashSHA384
		s.signerOpts = crypto.SHA384
		s.signFunc = signPSS
	case RS512:
		s.hashFunc = hashSHA512
		s.signerOpts = crypto.SHA512
		s.signFunc = signPKCS1v15
	case ES512:
		s.hashFunc = hashSHA512
		s.signerOpts = crypto.SHA512
		s.signFunc = signECDSA
	case PS512:
		s.hashFunc = hashSHA512
		s.signerOpts = crypto.SHA512
		s.signFunc = signPSS
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedAlgorithm, s.alg)
	}
	return s, nil
}

// Sign signs the provided event and returns it as a byte slice that includes
// the signature itself and the details needed to verify the signature.
//
//   - If something goes wrong while modifying the event in preparation of
//     the signing, ErrMarshaling is returned.
//   - If the signing itself fails, ErrSigningFailed is returned.
//
// Errors will be returned in wrapped form so make sure you use errors.Is
// rather than direct comparisons.
func (s *Signer) Sign(event json.Marshaler) ([]byte, error) {
	// TODO: Check if meta.security is supported in this event version.

	eventBytes, err := event.MarshalJSON()
	if err != nil {
		return nil, errors.Join(ErrMarshaling, err)
	}

	// Set the signing-related fields, make sure there's an empty signature field,
	// and transform the event to canonical JSON.
	if eventBytes, err = sjson.SetBytes(eventBytes, authorIdentityField, s.identity); err != nil {
		return nil, errors.Join(ErrMarshaling, err)
	}
	if eventBytes, err = sjson.SetBytes(eventBytes, algorithmField, s.alg); err != nil {
		return nil, errors.Join(ErrMarshaling, err)
	}
	if eventBytes, err = sjson.SetBytes(eventBytes, signatureField, ""); err != nil {
		return nil, errors.Join(ErrMarshaling, err)
	}
	if eventBytes, err = jcs.Transform(eventBytes); err != nil {
		return nil, errors.Join(ErrMarshaling, err)
	}

	sig, err := s.signFunc(s.pk, s.signerOpts.HashFunc(), s.hashFunc(eventBytes))
	if err != nil {
		return nil, errors.Join(ErrSigningFailed, err)
	}
	if eventBytes, err = sjson.SetBytes(eventBytes, signatureField, base64.StdEncoding.EncodeToString(sig)); err != nil {
		return nil, errors.Join(ErrMarshaling, err)
	}
	return eventBytes, nil
}

func signECDSA(priv crypto.PrivateKey, hash crypto.Hash, digest []byte) ([]byte, error) {
	privECDSA, ok := priv.(*ecdsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("private key had the wrong type; expected *ecdsa.PrivateKey, got %T", priv)
	}
	return ecdsa.SignASN1(rand.Reader, privECDSA, digest)
}

func signPKCS1v15(priv crypto.PrivateKey, hash crypto.Hash, digest []byte) ([]byte, error) {
	privRSA, ok := priv.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("private key had the wrong type; expected *rsa.PrivateKey, got %T", priv)
	}
	return rsa.SignPKCS1v15(rand.Reader, privRSA, hash, digest)
}

func signPSS(priv crypto.PrivateKey, hash crypto.Hash, digest []byte) ([]byte, error) {
	privRSA, ok := priv.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("private key had the wrong type; expected *rsa.PrivateKey, got %T", priv)
	}
	return rsa.SignPSS(rand.Reader, privRSA, hash, digest, nil)
}
