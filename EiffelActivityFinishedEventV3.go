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

// Code generated by eventgen. DO NOT EDIT.

package eiffelevents

import (
	"fmt"
	"reflect"
	"time"

	"github.com/clarketm/json"
	"github.com/google/uuid"
)

// NewActivityFinishedV3 creates a new struct pointer that represents
// major version 3 of EiffelActivityFinishedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 3.x.x
// currently known by this SDK.
func NewActivityFinishedV3(modifiers ...Modifier) (*ActivityFinishedV3, error) {
	var event ActivityFinishedV3
	event.Meta.Type = "EiffelActivityFinishedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][3].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new ActivityFinishedV3: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *ActivityFinishedV3) MarshalJSON() ([]byte, error) {
	// The standard encoding/json package doesn't honor omitempty for
	// non-pointer structs (it doesn't recurse into values, only examines
	// the immediate value). This is a not terribly elegant way of making
	// sure that this struct is marshaled by github.com/clarketm/json
	// without the infinite loop we'd get by just passing the struct to
	// github.com/clarketm/json.Marshal.
	//
	// Make sure the links slice is non-null so that non-initialized slices
	// get serialized as "[]" instead of "null".
	links := e.Links
	if links == nil {
		links = make(EventLinksV1, 0)
	}
	s := struct {
		Data  *ActFV3Data  `json:"data"`
		Links EventLinksV1 `json:"links"`
		Meta  *MetaV3      `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *ActivityFinishedV3) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *ActivityFinishedV3) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var (
	_ CapabilityTeller = &ActivityFinishedV3{}
	_ FieldSetter      = &ActivityFinishedV3{}
	_ MetaTeller       = &ActivityFinishedV3{}
)

// ID returns the value of the meta.id field.
func (e ActivityFinishedV3) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e ActivityFinishedV3) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e ActivityFinishedV3) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e ActivityFinishedV3) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e ActivityFinishedV3) DomainID() string {
	return e.Meta.Source.DomainID
}

// SupportsSigning returns true if the event supports signatures according
// to V3 of the meta field, i.e. events where the signature is found under
// meta.security.integrityProtection.
func (e ActivityFinishedV3) SupportsSigning() bool {
	return true
}

type ActivityFinishedV3 struct {
	// Mandatory fields
	Data  ActFV3Data   `json:"data"`
	Links EventLinksV1 `json:"links"`
	Meta  MetaV3       `json:"meta"`

	// Optional fields

}

type ActFV3Data struct {
	// Mandatory fields
	Outcome ActFV3DataOutcome `json:"outcome"`

	// Optional fields
	CustomData     []CustomDataV1            `json:"customData,omitempty"`
	PersistentLogs []ActFV3DataPersistentLog `json:"persistentLogs,omitempty"`
}

type ActFV3DataOutcome struct {
	// Mandatory fields
	Conclusion ActFV3DataOutcomeConclusion `json:"conclusion"`

	// Optional fields
	Description string `json:"description,omitempty"`
}

type ActFV3DataOutcomeConclusion string

const (
	ActFV3DataOutcomeConclusion_Successful   ActFV3DataOutcomeConclusion = "SUCCESSFUL"
	ActFV3DataOutcomeConclusion_Unsuccessful ActFV3DataOutcomeConclusion = "UNSUCCESSFUL"
	ActFV3DataOutcomeConclusion_Failed       ActFV3DataOutcomeConclusion = "FAILED"
	ActFV3DataOutcomeConclusion_Aborted      ActFV3DataOutcomeConclusion = "ABORTED"
	ActFV3DataOutcomeConclusion_TimedOut     ActFV3DataOutcomeConclusion = "TIMED_OUT"
	ActFV3DataOutcomeConclusion_Inconclusive ActFV3DataOutcomeConclusion = "INCONCLUSIVE"
)

type ActFV3DataPersistentLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields
	MediaType string   `json:"mediaType,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}
