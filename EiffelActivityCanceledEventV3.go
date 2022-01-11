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

// NewActivityCanceledV3 creates a new struct pointer that represents
// major version 3 of EiffelActivityCanceledEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 3.x.x
// currently known by this SDK.
func NewActivityCanceledV3(modifiers ...Modifier) (*ActivityCanceledV3, error) {
	var event ActivityCanceledV3
	event.Meta.Type = "EiffelActivityCanceledEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][3].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new ActivityCanceledV3: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *ActivityCanceledV3) MarshalJSON() ([]byte, error) {
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
		links = make([]ActCV3Link, 0)
	}
	s := struct {
		Data  *ActCV3Data  `json:"data"`
		Links []ActCV3Link `json:"links"`
		Meta  *ActCV3Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *ActivityCanceledV3) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *ActivityCanceledV3) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &ActivityCanceledV3{}
var _ MetaTeller = &ActivityCanceledV3{}

// ID returns the value of the meta.id field.
func (e ActivityCanceledV3) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e ActivityCanceledV3) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e ActivityCanceledV3) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e ActivityCanceledV3) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e ActivityCanceledV3) DomainID() string {
	return e.Meta.Source.DomainID
}

type ActivityCanceledV3 struct {
	// Mandatory fields
	Data  ActCV3Data  `json:"data"`
	Links ActCV3Links `json:"links"`
	Meta  ActCV3Meta  `json:"meta"`

	// Optional fields

}

type ActCV3Data struct {
	// Mandatory fields

	// Optional fields
	CustomData []ActCV3DataCustomDatum `json:"customData,omitempty"`
	Reason     string                  `json:"reason,omitempty"`
}

type ActCV3DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

// ActCV3Links represents a slice of ActCV3Link values with helper methods
// for adding new links.
type ActCV3Links []ActCV3Link

// Add adds a new link of the specified type to a target event.
func (links *ActCV3Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, ActCV3Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *ActCV3Links) AddByID(linkType string, target string) {
	*links = append(*links, ActCV3Link{Target: target, Type: linkType})
}

type ActCV3Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields
	DomainID string `json:"domainId,omitempty"`
}

type ActCV3Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security ActCV3MetaSecurity `json:"security,omitempty"`
	Source   ActCV3MetaSource   `json:"source,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

type ActCV3MetaSecurity struct {
	// Mandatory fields
	AuthorIdentity string `json:"authorIdentity"`

	// Optional fields
	IntegrityProtection ActCV3MetaSecurityIntegrityProtection  `json:"integrityProtection,omitempty"`
	SequenceProtection  []ActCV3MetaSecuritySequenceProtection `json:"sequenceProtection,omitempty"`
}

type ActCV3MetaSecurityIntegrityProtection struct {
	// Mandatory fields
	Alg       ActCV3MetaSecurityIntegrityProtectionAlg `json:"alg"`
	Signature string                                   `json:"signature"`

	// Optional fields
	PublicKey string `json:"publicKey,omitempty"`
}

type ActCV3MetaSecurityIntegrityProtectionAlg string

const (
	ActCV3MetaSecurityIntegrityProtectionAlg_HS256 ActCV3MetaSecurityIntegrityProtectionAlg = "HS256"
	ActCV3MetaSecurityIntegrityProtectionAlg_HS384 ActCV3MetaSecurityIntegrityProtectionAlg = "HS384"
	ActCV3MetaSecurityIntegrityProtectionAlg_HS512 ActCV3MetaSecurityIntegrityProtectionAlg = "HS512"
	ActCV3MetaSecurityIntegrityProtectionAlg_RS256 ActCV3MetaSecurityIntegrityProtectionAlg = "RS256"
	ActCV3MetaSecurityIntegrityProtectionAlg_RS384 ActCV3MetaSecurityIntegrityProtectionAlg = "RS384"
	ActCV3MetaSecurityIntegrityProtectionAlg_RS512 ActCV3MetaSecurityIntegrityProtectionAlg = "RS512"
	ActCV3MetaSecurityIntegrityProtectionAlg_ES256 ActCV3MetaSecurityIntegrityProtectionAlg = "ES256"
	ActCV3MetaSecurityIntegrityProtectionAlg_ES384 ActCV3MetaSecurityIntegrityProtectionAlg = "ES384"
	ActCV3MetaSecurityIntegrityProtectionAlg_ES512 ActCV3MetaSecurityIntegrityProtectionAlg = "ES512"
	ActCV3MetaSecurityIntegrityProtectionAlg_PS256 ActCV3MetaSecurityIntegrityProtectionAlg = "PS256"
	ActCV3MetaSecurityIntegrityProtectionAlg_PS384 ActCV3MetaSecurityIntegrityProtectionAlg = "PS384"
	ActCV3MetaSecurityIntegrityProtectionAlg_PS512 ActCV3MetaSecurityIntegrityProtectionAlg = "PS512"
)

type ActCV3MetaSecuritySequenceProtection struct {
	// Mandatory fields
	Position     int64  `json:"position"`
	SequenceName string `json:"sequenceName"`

	// Optional fields

}

type ActCV3MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
