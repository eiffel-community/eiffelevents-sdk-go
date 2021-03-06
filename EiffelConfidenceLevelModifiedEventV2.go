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
		links = make([]CLMV2Link, 0)
	}
	s := struct {
		Data  *CLMV2Data  `json:"data"`
		Links []CLMV2Link `json:"links"`
		Meta  *CLMV2Meta  `json:"meta"`
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

var _ FieldSetter = &ConfidenceLevelModifiedV2{}
var _ MetaTeller = &ConfidenceLevelModifiedV2{}

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

type ConfidenceLevelModifiedV2 struct {
	// Mandatory fields
	Data  CLMV2Data  `json:"data"`
	Links CLMV2Links `json:"links"`
	Meta  CLMV2Meta  `json:"meta"`

	// Optional fields

}

type CLMV2Data struct {
	// Mandatory fields
	Name  string         `json:"name"`
	Value CLMV2DataValue `json:"value"`

	// Optional fields
	CustomData []CLMV2DataCustomDatum `json:"customData,omitempty"`
	Issuer     CLMV2DataIssuer        `json:"issuer,omitempty"`
}

type CLMV2DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

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

// CLMV2Links represents a slice of CLMV2Link values with helper methods
// for adding new links.
type CLMV2Links []CLMV2Link

var _ LinkFinder = &CLMV2Links{}

// Add adds a new link of the specified type to a target event.
func (links *CLMV2Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, CLMV2Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *CLMV2Links) AddByID(linkType string, target string) {
	*links = append(*links, CLMV2Link{Target: target, Type: linkType})
}

// FindAll returns the IDs of all links of the specified type, or an empty
// slice if no such links are found.
func (links CLMV2Links) FindAll(linkType string) []string {
	result := make([]string, 0, len(links))
	for _, link := range links {
		if link.Type == linkType {
			result = append(result, link.Target)
		}
	}
	return result
}

// FindFirst returns the ID of the first encountered link of the specified
// type, or an empty string if no such link is found.
func (links CLMV2Links) FindFirst(linkType string) string {
	for _, link := range links {
		if link.Type == linkType {
			return link.Target
		}
	}
	return ""
}

type CLMV2Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type CLMV2Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security CLMV2MetaSecurity `json:"security,omitempty"`
	Source   CLMV2MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type CLMV2MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM CLMV2MetaSecuritySDM `json:"sdm,omitempty"`
}

type CLMV2MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type CLMV2MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
