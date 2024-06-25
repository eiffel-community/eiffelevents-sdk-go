package signature

import (
	"fmt"

	"github.com/go-ldap/ldap"
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
