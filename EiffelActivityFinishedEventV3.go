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
		links = make([]ActFV3Link, 0)
	}
	s := struct {
		Data  *ActFV3Data  `json:"data"`
		Links []ActFV3Link `json:"links"`
		Meta  *ActFV3Meta  `json:"meta"`
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

var _ FieldSetter = &ActivityFinishedV3{}
var _ MetaTeller = &ActivityFinishedV3{}

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

type ActivityFinishedV3 struct {
	// Mandatory fields
	Data  ActFV3Data  `json:"data"`
	Links ActFV3Links `json:"links"`
	Meta  ActFV3Meta  `json:"meta"`

	// Optional fields

}

type ActFV3Data struct {
	// Mandatory fields
	Outcome ActFV3DataOutcome `json:"outcome"`

	// Optional fields
	CustomData     []ActFV3DataCustomDatum   `json:"customData,omitempty"`
	PersistentLogs []ActFV3DataPersistentLog `json:"persistentLogs,omitempty"`
}

type ActFV3DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

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

// ActFV3Links represents a slice of ActFV3Link values with helper methods
// for adding new links.
type ActFV3Links []ActFV3Link

var _ LinkFinder = &ActFV3Links{}

// Add adds a new link of the specified type to a target event.
func (links *ActFV3Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, ActFV3Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *ActFV3Links) AddByID(linkType string, target string) {
	*links = append(*links, ActFV3Link{Target: target, Type: linkType})
}

// FindAll returns the IDs of all links of the specified type, or an empty
// slice if no such links are found.
func (links ActFV3Links) FindAll(linkType string) []string {
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
func (links ActFV3Links) FindFirst(linkType string) string {
	for _, link := range links {
		if link.Type == linkType {
			return link.Target
		}
	}
	return ""
}

type ActFV3Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields
	DomainID string `json:"domainId,omitempty"`
}

type ActFV3Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security ActFV3MetaSecurity `json:"security,omitempty"`
	Source   ActFV3MetaSource   `json:"source,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

type ActFV3MetaSecurity struct {
	// Mandatory fields
	AuthorIdentity string `json:"authorIdentity"`

	// Optional fields
	IntegrityProtection ActFV3MetaSecurityIntegrityProtection  `json:"integrityProtection,omitempty"`
	SequenceProtection  []ActFV3MetaSecuritySequenceProtection `json:"sequenceProtection,omitempty"`
}

type ActFV3MetaSecurityIntegrityProtection struct {
	// Mandatory fields
	Alg       ActFV3MetaSecurityIntegrityProtectionAlg `json:"alg"`
	Signature string                                   `json:"signature"`

	// Optional fields
	PublicKey string `json:"publicKey,omitempty"`
}

type ActFV3MetaSecurityIntegrityProtectionAlg string

const (
	ActFV3MetaSecurityIntegrityProtectionAlg_HS256 ActFV3MetaSecurityIntegrityProtectionAlg = "HS256"
	ActFV3MetaSecurityIntegrityProtectionAlg_HS384 ActFV3MetaSecurityIntegrityProtectionAlg = "HS384"
	ActFV3MetaSecurityIntegrityProtectionAlg_HS512 ActFV3MetaSecurityIntegrityProtectionAlg = "HS512"
	ActFV3MetaSecurityIntegrityProtectionAlg_RS256 ActFV3MetaSecurityIntegrityProtectionAlg = "RS256"
	ActFV3MetaSecurityIntegrityProtectionAlg_RS384 ActFV3MetaSecurityIntegrityProtectionAlg = "RS384"
	ActFV3MetaSecurityIntegrityProtectionAlg_RS512 ActFV3MetaSecurityIntegrityProtectionAlg = "RS512"
	ActFV3MetaSecurityIntegrityProtectionAlg_ES256 ActFV3MetaSecurityIntegrityProtectionAlg = "ES256"
	ActFV3MetaSecurityIntegrityProtectionAlg_ES384 ActFV3MetaSecurityIntegrityProtectionAlg = "ES384"
	ActFV3MetaSecurityIntegrityProtectionAlg_ES512 ActFV3MetaSecurityIntegrityProtectionAlg = "ES512"
	ActFV3MetaSecurityIntegrityProtectionAlg_PS256 ActFV3MetaSecurityIntegrityProtectionAlg = "PS256"
	ActFV3MetaSecurityIntegrityProtectionAlg_PS384 ActFV3MetaSecurityIntegrityProtectionAlg = "PS384"
	ActFV3MetaSecurityIntegrityProtectionAlg_PS512 ActFV3MetaSecurityIntegrityProtectionAlg = "PS512"
)

type ActFV3MetaSecuritySequenceProtection struct {
	// Mandatory fields
	Position     int64  `json:"position"`
	SequenceName string `json:"sequenceName"`

	// Optional fields

}

type ActFV3MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
