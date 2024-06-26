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

// THIS FILE IS AUTOMATICALLY GENERATED AND MUST NOT BE EDITED BY HAND.

package eiffelevents

import (
	"fmt"
	"reflect"
	"time"

	"github.com/clarketm/json"
	"github.com/google/uuid"
)

// NewConfidenceLevelModifiedV2 creates a new struct pointer that represents
// major version 2 of EiffelConfidenceLevelModifiedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 2.x.x
// currently known by this SDK.
func NewConfidenceLevelModifiedV2(modifiers ...Modifier) (*ConfidenceLevelModifiedV2, error) {
	var event ConfidenceLevelModifiedV2
	event.Meta.Type = "EiffelConfidenceLevelModifiedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][2].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new ConfidenceLevelModifiedV2: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *ConfidenceLevelModifiedV2) MarshalJSON() ([]byte, error) {
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
		Data  *CLMV2Data   `json:"data"`
		Links EventLinksV1 `json:"links"`
		Meta  *MetaV2      `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *ConfidenceLevelModifiedV2) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *ConfidenceLevelModifiedV2) String() string {
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
	_ CapabilityTeller = &ConfidenceLevelModifiedV2{}
	_ FieldSetter      = &ConfidenceLevelModifiedV2{}
	_ MetaTeller       = &ConfidenceLevelModifiedV2{}
)

// ID returns the value of the meta.id field.
func (e ConfidenceLevelModifiedV2) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e ConfidenceLevelModifiedV2) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e ConfidenceLevelModifiedV2) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e ConfidenceLevelModifiedV2) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e ConfidenceLevelModifiedV2) DomainID() string {
	return e.Meta.Source.DomainID
}

// SupportsSigning returns true if the event supports signatures according
// to V3 of the meta field, i.e. events where the signature is found under
// meta.security.integrityProtection.
func (e ConfidenceLevelModifiedV2) SupportsSigning() bool {
	return false
}

type ConfidenceLevelModifiedV2 struct {
	// Mandatory fields
	Data  CLMV2Data    `json:"data"`
	Links EventLinksV1 `json:"links"`
	Meta  MetaV2       `json:"meta"`

	// Optional fields

}

type CLMV2Data struct {
	// Mandatory fields
	Name  string         `json:"name"`
	Value CLMV2DataValue `json:"value"`

	// Optional fields
	CustomData []CustomDataV1  `json:"customData,omitempty"`
	Issuer     CLMV2DataIssuer `json:"issuer,omitempty"`
}

type CLMV2DataIssuer struct {
	// Mandatory fields

	// Optional fields
	Email string `json:"email,omitempty"`
	Group string `json:"group,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}

type CLMV2DataValue string

const (
	CLMV2DataValue_Success      CLMV2DataValue = "SUCCESS"
	CLMV2DataValue_Failure      CLMV2DataValue = "FAILURE"
	CLMV2DataValue_Inconclusive CLMV2DataValue = "INCONCLUSIVE"
)
