package eiffelevents

import (
	"fmt"

	"github.com/Showmax/go-fqdn"
)

// Modifier defines a kind of function that modifies an untyped Eiffel event.
// It can be passed to factory functions to adorn the newly created event with
// additional fields. The main use case is generic field changes that should
// apply to all events emitted from an application, e.g. the setting of
// meta.source.domainId.
//
// Using modifiers together with a factory function allows creation of new
// factory functions that include the modifiers. Those factory function can
// be passed to other parts of the application to reduce duplication.
type Modifier func(fieldSetter FieldSetter) error

// WithAutoSourceHost attempts to figure out the FQDN of the current host and
// stores it in the meta.source.host field.
// WithVersion sets the meta.version field of a newly created event.
func WithAutoSourceHost() Modifier {
	return func(fieldSetter FieldSetter) error {
		hostname, err := fqdn.FqdnHostname()
		if err != nil {
			return fmt.Errorf("error determining the local hostname to store in meta.source.host: %w", err)
		}
		return fieldSetter.SetField("meta.source.host", hostname)
	}
}

// WithSourceDomainID sets the meta.source.domainId field of a newly created event.
func WithSourceDomainID(domainID string) Modifier {
	return func(fieldSetter FieldSetter) error {
		return fieldSetter.SetField("meta.source.domainId", domainID)
	}
}

// WithSourceHost sets the meta.source.host field of a newly created event.
func WithSourceHost(hostname string) Modifier {
	return func(fieldSetter FieldSetter) error {
		return fieldSetter.SetField("meta.source.host", hostname)
	}
}

// WithSourceName sets the meta.source.name field of a newly created event.
func WithSourceName(name string) Modifier {
	return func(fieldSetter FieldSetter) error {
		return fieldSetter.SetField("meta.source.name", name)
	}
}

// WithSourceURI sets the meta.source.uri field of a newly created event.
func WithSourceURI(uri string) Modifier {
	return func(fieldSetter FieldSetter) error {
		return fieldSetter.SetField("meta.source.uri", uri)
	}
}

// WithVersion sets the meta.version field of a newly created event.
// It's typically used to select a different minor or patch version than
// what the default factory function does.
func WithVersion(version string) Modifier {
	return func(fieldSetter FieldSetter) error {
		return fieldSetter.SetField("meta.version", version)
	}
}
