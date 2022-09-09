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

package eiffelevents

import (
	"errors"
	"fmt"
	"path"
	"runtime/debug"

	"github.com/Showmax/go-fqdn"
	"github.com/package-url/packageurl-go"
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

// WithAutoSourceSerializer sets the meta.source.serializer field to a purl
// describing the program's main package. It can't figure out the main package
// version so that value needs to be passed via the version parameter.
func WithAutoSourceSerializer(version string) Modifier {
	return func(fieldsetter FieldSetter) error {
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			return errors.New("no build information available")
		}
		purl := purlFromBuildInfo(bi)
		purl.Version = version
		return fieldsetter.SetField("meta.source.serializer", purl.String())
	}
}

func purlFromBuildInfo(bi *debug.BuildInfo) *packageurl.PackageURL {
	namespace := path.Dir(bi.Path)
	if namespace == "." {
		namespace = ""
	}
	return packageurl.NewPackageURL("golang", namespace, path.Base(bi.Path), "", packageurl.Qualifiers{}, "")
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

// WithSourceSerializer sets the meta.source.serializer field of a newly created event.
func WithSourceSerializer(serializer string) Modifier {
	return func(fieldSetter FieldSetter) error {
		return fieldSetter.SetField("meta.source.serializer", serializer)
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
