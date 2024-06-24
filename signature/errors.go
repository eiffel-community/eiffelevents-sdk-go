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

import "errors"

var (
	ErrKeyTypeMismatch      = errors.New("key is of the wrong type")
	ErrMarshaling           = errors.New("the marshaling of the event was unsuccessful")
	ErrPublicKeyLookup      = errors.New("an error occurred looking up the public key for this identity")
	ErrPublicKeyNotFound    = errors.New("no public key for verifying events signed by this identify was found")
	ErrSignatureMismatch    = errors.New("the signature couldn't be verified")
	ErrSigningFailed        = errors.New("signing of the event failed")
	ErrUnsupportedAlgorithm = errors.New("unsupported algorithm")
	ErrUnverifiableEvent    = errors.New("event cannot be verified because an essential field is unset or empty")
	ErrVerificationFailed   = errors.New("the signature couldn't be verified by any of the available public keys")
)
