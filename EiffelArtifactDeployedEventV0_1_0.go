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

// NewArtifactDeployedV0_1_0 creates a new struct pointer that represents version
// 0.1.0 of EiffelArtifactDeployedEvent. The returned struct has all required
// meta members populated.
func NewArtifactDeployedV0_1_0(modifiers ...Modifier) (*ArtifactDeployedV0_1_0, error) {
	var event ArtifactDeployedV0_1_0
	event.Meta.Type = "EiffelArtifactDeployedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][0].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new ArtifactDeployedV0_1_0: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *ArtifactDeployedV0_1_0) MarshalJSON() ([]byte, error) {
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
		Data  *ArtDV0_1_0Data `json:"data"`
		Links EventLinksV1    `json:"links"`
		Meta  *MetaV3         `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *ArtifactDeployedV0_1_0) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *ArtifactDeployedV0_1_0) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &ArtifactDeployedV0_1_0{}
var _ MetaTeller = &ArtifactDeployedV0_1_0{}

// ID returns the value of the meta.id field.
func (e ArtifactDeployedV0_1_0) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e ArtifactDeployedV0_1_0) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e ArtifactDeployedV0_1_0) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e ArtifactDeployedV0_1_0) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e ArtifactDeployedV0_1_0) DomainID() string {
	return e.Meta.Source.DomainID
}

type ArtifactDeployedV0_1_0 struct {
	// Mandatory fields
	Data  ArtDV0_1_0Data `json:"data"`
	Links EventLinksV1   `json:"links"`
	Meta  MetaV3         `json:"meta"`

	// Optional fields

}

type ArtDV0_1_0Data struct {
	// Mandatory fields

	// Optional fields
	CustomData  []CustomDataV1 `json:"customData,omitempty"`
	Description string         `json:"description,omitempty"`
	URI         string         `json:"uri,omitempty"`
}
