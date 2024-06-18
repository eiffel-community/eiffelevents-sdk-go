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
	"encoding/json"
	"errors"
	"fmt"
)

// Any is a struct that represents any Eiffel event value. Its intended use is
// to allow JSON objects with arrays of Eiffel events to be unmarshalled into
// something useful.
//
// The type also implements the MetaTeller interface so if you only need e.g.
// the ID of each event in a slice you don't have to do a type assertion on
// every element.
type Any struct {
	event interface{}
}

var (
	_ MetaTeller       = &Any{}
	_ json.Marshaler   = &Any{}
	_ json.Unmarshaler = &Any{}
)

// Get obtains the event as an interface pointer. Use a type assertion to
// convert it to a concrete event struct pointer.
func (a Any) Get() interface{} {
	return a.event
}

// MarshalJSON converts the event to its JSON representation.
func (a Any) MarshalJSON() ([]byte, error) {
	if v, ok := a.event.(json.Marshaler); ok {
		return v.MarshalJSON()
	}
	return nil, errors.New("value not marshalable as JSON")
}

// UnmarshalJSON parses the byte slice input as JSON and stores it.
func (a *Any) UnmarshalJSON(b []byte) error {
	var err error
	a.event, err = UnmarshalAny(b)
	if err != nil {
		return err
	}

	// These checks are probably unnecessary since UnmarshalAny is pretty
	// picky about what it accepts.
	if _, ok := a.event.(FieldSetter); !ok {
		return fmt.Errorf("unmarshaling did not produce an Eiffel event: %w", err)
	}
	if _, ok := a.event.(MetaTeller); !ok {
		return fmt.Errorf("unmarshaling did not produce an Eiffel event: %w", err)
	}
	return nil
}

// ID returns the value of the meta.id field.
func (a Any) ID() string {
	return a.event.(MetaTeller).ID() // nolint:forcetypeassert
}

// Type returns the value of the meta.type field.
func (a Any) Type() string {
	return a.event.(MetaTeller).Type() // nolint:forcetypeassert
}

// Version returns the value of the meta.version field.
func (a Any) Version() string {
	return a.event.(MetaTeller).Version() // nolint:forcetypeassert
}

// Time returns the value of the meta.time field.
func (a Any) Time() int64 {
	return a.event.(MetaTeller).Time() // nolint:forcetypeassert
}

// DomainID returns the value of the meta.source.domainId field.
func (a Any) DomainID() string {
	return a.event.(MetaTeller).DomainID() // nolint:forcetypeassert
}
