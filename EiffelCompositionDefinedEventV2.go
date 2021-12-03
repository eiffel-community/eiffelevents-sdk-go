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

// NewCompositionDefinedV2 creates a new struct pointer that represents
// major version 2 of EiffelCompositionDefinedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 2.x.x
// currently known by this SDK.
func NewCompositionDefinedV2(modifiers ...Modifier) (*CompositionDefinedV2, error) {
	var event CompositionDefinedV2
	event.Meta.Type = "EiffelCompositionDefinedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][2].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new CompositionDefinedV2: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *CompositionDefinedV2) MarshalJSON() ([]byte, error) {
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
		links = make([]CDV2Link, 0)
	}
	s := struct {
		Data  *CDV2Data  `json:"data"`
		Links []CDV2Link `json:"links"`
		Meta  *CDV2Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *CompositionDefinedV2) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *CompositionDefinedV2) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &CompositionDefinedV2{}

type CompositionDefinedV2 struct {
	// Mandatory fields
	Data  CDV2Data   `json:"data"`
	Links []CDV2Link `json:"links"`
	Meta  CDV2Meta   `json:"meta"`

	// Optional fields

}

type CDV2Data struct {
	// Mandatory fields
	Name string `json:"name"`

	// Optional fields
	CustomData []CDV2DataCustomDatum `json:"customData,omitempty"`
	Version    string                `json:"version,omitempty"`
}

type CDV2DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type CDV2Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type CDV2Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security CDV2MetaSecurity `json:"security,omitempty"`
	Source   CDV2MetaSource   `json:"source,omitempty"`
	Tags     []string         `json:"tags,omitempty"`
}

type CDV2MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM CDV2MetaSecuritySDM `json:"sdm,omitempty"`
}

type CDV2MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type CDV2MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
