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

// NewFlowContextDefinedV1 creates a new struct pointer that represents
// major version 1 of EiffelFlowContextDefinedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 1.x.x
// currently known by this SDK.
func NewFlowContextDefinedV1(modifiers ...Modifier) (*FlowContextDefinedV1, error) {
	var event FlowContextDefinedV1
	event.Meta.Type = "EiffelFlowContextDefinedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][1].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new FlowContextDefinedV1: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *FlowContextDefinedV1) MarshalJSON() ([]byte, error) {
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
		links = make([]FCDV1Link, 0)
	}
	s := struct {
		Data  *FCDV1Data  `json:"data"`
		Links []FCDV1Link `json:"links"`
		Meta  *FCDV1Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *FlowContextDefinedV1) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *FlowContextDefinedV1) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &FlowContextDefinedV1{}
var _ MetaTeller = &FlowContextDefinedV1{}

// ID returns the value of the meta.id field.
func (e FlowContextDefinedV1) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e FlowContextDefinedV1) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e FlowContextDefinedV1) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e FlowContextDefinedV1) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e FlowContextDefinedV1) DomainID() string {
	return e.Meta.Source.DomainID
}

type FlowContextDefinedV1 struct {
	// Mandatory fields
	Data  FCDV1Data  `json:"data"`
	Links FCDV1Links `json:"links"`
	Meta  FCDV1Meta  `json:"meta"`

	// Optional fields

}

type FCDV1Data struct {
	// Mandatory fields

	// Optional fields
	CustomData []FCDV1DataCustomDatum `json:"customData,omitempty"`
	Product    string                 `json:"product,omitempty"`
	Program    string                 `json:"program,omitempty"`
	Project    string                 `json:"project,omitempty"`
	Track      string                 `json:"track,omitempty"`
	Version    string                 `json:"version,omitempty"`
}

type FCDV1DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

// FCDV1Links represents a slice of FCDV1Link values with helper methods
// for adding new links.
type FCDV1Links []FCDV1Link

var _ LinkFinder = &FCDV1Links{}

// Add adds a new link of the specified type to a target event.
func (links *FCDV1Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, FCDV1Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *FCDV1Links) AddByID(linkType string, target string) {
	*links = append(*links, FCDV1Link{Target: target, Type: linkType})
}

// FindAll returns the IDs of all links of the specified type, or an empty
// slice if no such links are found.
func (links FCDV1Links) FindAll(linkType string) []string {
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
func (links FCDV1Links) FindFirst(linkType string) string {
	for _, link := range links {
		if link.Type == linkType {
			return link.Target
		}
	}
	return ""
}

type FCDV1Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type FCDV1Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security FCDV1MetaSecurity `json:"security,omitempty"`
	Source   FCDV1MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type FCDV1MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM FCDV1MetaSecuritySDM `json:"sdm,omitempty"`
}

type FCDV1MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type FCDV1MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string                    `json:"domainId,omitempty"`
	Host       string                    `json:"host,omitempty"`
	Name       string                    `json:"name,omitempty"`
	Serializer FCDV1MetaSourceSerializer `json:"serializer,omitempty"`
	URI        string                    `json:"uri,omitempty"`
}

type FCDV1MetaSourceSerializer struct {
	// Mandatory fields
	ArtifactID string `json:"artifactId"`
	GroupID    string `json:"groupId"`
	Version    string `json:"version"`

	// Optional fields

}
