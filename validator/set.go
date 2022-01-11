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

package validator

import (
	"context"
	"net/http"
)

// Validator is capable of applying some set of rules to validate
// the correctness of an event.
type Validator interface {
	Validate(ctx context.Context, event []byte) error
}

// ValidatorSet contains an ordered set of one or more Validator
// instances. The set can validate events against all validators
// in the set and require all validators to give a passing grade.
type ValidatorSet struct {
	validators []Validator
}

// NewSet returns a ValidatorSet containing the specified Validator pointers.
func NewSet(validators ...Validator) *ValidatorSet {
	return &ValidatorSet{
		validators: validators,
	}
}

// Add appends one or more validators to the current set.
func (vs *ValidatorSet) Add(validators ...Validator) {
	vs.validators = append(vs.validators, validators...)
}

// Validate loops over the Validator instances in the set and asks them to
// validate the given event. The loop will terminate upon the first validation
// error, i.e. all validators aren't guaranteed to be called.
func (vs *ValidatorSet) Validate(ctx context.Context, event []byte) error {
	for _, v := range vs.validators {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if err := v.Validate(ctx, event); err != nil {
			return err
		}
	}
	return nil
}

// DefaultSet returns the currently recommended set of validators,
// each with a default configuration.
func DefaultSet() *ValidatorSet {
	return NewSet(
		NewSchemaValidator(
			NewMetaSchemaLocator(http.DefaultClient),
			NewBundledSchemaLocator(),
		),
	)
}
