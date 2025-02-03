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
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
)

// AuthorIdentity is a representation of the distinguished name
// in the meta.security.authorIdentity field of an Eiffel event.
// It can be compared to other values of the same type.
type AuthorIdentity struct {
	dn       *ldap.DN
	original string
}

func NewAuthorIdentity(s string) (*AuthorIdentity, error) {
	dn, err := ldap.ParseDN(s)
	if err != nil {
		return nil, fmt.Errorf("error parsing author identity %q: %w", s, err)
	}
	return &AuthorIdentity{
		dn:       dn,
		original: s,
	}, nil
}

// Equal returns true if the provided *AuthorIdentity is equal to this one,
// igoring differences in whitespace etc.
func (ai AuthorIdentity) Equal(other *AuthorIdentity) bool {
	return ai.dn.Equal(other.dn)
}

func (ai AuthorIdentity) String() string {
	return ai.original
}

func (ai AuthorIdentity) MarshalText() (text []byte, err error) {
	return []byte(ai.original), nil
}

func (ai *AuthorIdentity) UnmarshalText(text []byte) error {
	i, err := NewAuthorIdentity(string(text))
	if err != nil {
		return err
	}
	ai.dn = i.dn
	ai.original = i.original
	return nil
}
