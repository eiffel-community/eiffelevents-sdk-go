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

// NewActivityCanceledV2 creates a new struct pointer that represents
// major version 2 of EiffelActivityCanceledEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 2.x.x
// currently known by this SDK.
func NewActivityCanceledV2(modifiers ...Modifier) (*ActivityCanceledV2, error) {
	var event ActivityCanceledV2
	event.Meta.Type = "EiffelActivityCanceledEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][2].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new ActivityCanceledV2: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *ActivityCanceledV2) MarshalJSON() ([]byte, error) {
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
		links = make([]ActCV2Link, 0)
	}
	s := struct {
		Data  *ActCV2Data  `json:"data"`
		Links []ActCV2Link `json:"links"`
		Meta  *ActCV2Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *ActivityCanceledV2) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *ActivityCanceledV2) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &ActivityCanceledV2{}
var _ MetaTeller = &ActivityCanceledV2{}

// ID returns the value of the meta.id field.
func (e ActivityCanceledV2) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e ActivityCanceledV2) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e ActivityCanceledV2) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e ActivityCanceledV2) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e ActivityCanceledV2) DomainID() string {
	return e.Meta.Source.DomainID
}

type ActivityCanceledV2 struct {
	// Mandatory fields
	Data  ActCV2Data  `json:"data"`
	Links ActCV2Links `json:"links"`
	Meta  ActCV2Meta  `json:"meta"`

	// Optional fields

}

type ActCV2Data struct {
	// Mandatory fields

	// Optional fields
	CustomData []ActCV2DataCustomDatum `json:"customData,omitempty"`
	Reason     string                  `json:"reason,omitempty"`
}

type ActCV2DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

// ActCV2Links represents a slice of ActCV2Link values with helper methods
// for adding new links.
type ActCV2Links []ActCV2Link

// Add adds a new link of the specified type to a target event.
func (links *ActCV2Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, ActCV2Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *ActCV2Links) AddByID(linkType string, target string) {
	*links = append(*links, ActCV2Link{Target: target, Type: linkType})
}

type ActCV2Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type ActCV2Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security ActCV2MetaSecurity `json:"security,omitempty"`
	Source   ActCV2MetaSource   `json:"source,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

type ActCV2MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM ActCV2MetaSecuritySDM `json:"sdm,omitempty"`
}

type ActCV2MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type ActCV2MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
