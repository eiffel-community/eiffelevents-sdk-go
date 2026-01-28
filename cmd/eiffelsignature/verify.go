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
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/eiffel-community/eiffelevents-sdk-go/signature"
)

func verifyCmd(ctx context.Context, args []string, in io.Reader) error {
	if len(args) < 1 {
		return fmt.Errorf("%w: not enough arguments for verify command", ErrUsage)
	}
	keyDir := args[0]

	locator := signature.NewFSPublicKeyLocator(signature.FSPublicKeyLocatorConfig{KeyDirectory: keyDir})
	verifier := signature.NewVerifier(locator)
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

		if err := verifier.Verify(ctx, []byte(payloadIn)); err != nil {
			return fmt.Errorf("unable to verify event signature: %w", err)
		}
	}
	return nil
}
