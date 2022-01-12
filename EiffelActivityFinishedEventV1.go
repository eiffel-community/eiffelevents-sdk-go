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

// NewActivityFinishedV1 creates a new struct pointer that represents
// major version 1 of EiffelActivityFinishedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 1.x.x
// currently known by this SDK.
func NewActivityFinishedV1(modifiers ...Modifier) (*ActivityFinishedV1, error) {
	var event ActivityFinishedV1
	event.Meta.Type = "EiffelActivityFinishedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][1].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new ActivityFinishedV1: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *ActivityFinishedV1) MarshalJSON() ([]byte, error) {
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
		links = make([]ActFV1Link, 0)
	}
	s := struct {
		Data  *ActFV1Data  `json:"data"`
		Links []ActFV1Link `json:"links"`
		Meta  *ActFV1Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *ActivityFinishedV1) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *ActivityFinishedV1) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &ActivityFinishedV1{}
var _ MetaTeller = &ActivityFinishedV1{}

// ID returns the value of the meta.id field.
func (e ActivityFinishedV1) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e ActivityFinishedV1) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e ActivityFinishedV1) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e ActivityFinishedV1) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e ActivityFinishedV1) DomainID() string {
	return e.Meta.Source.DomainID
}

type ActivityFinishedV1 struct {
	// Mandatory fields
	Data  ActFV1Data  `json:"data"`
	Links ActFV1Links `json:"links"`
	Meta  ActFV1Meta  `json:"meta"`

	// Optional fields

}

type ActFV1Data struct {
	// Mandatory fields
	Outcome ActFV1DataOutcome `json:"outcome"`

	// Optional fields
	CustomData     []ActFV1DataCustomDatum   `json:"customData,omitempty"`
	PersistentLogs []ActFV1DataPersistentLog `json:"persistentLogs,omitempty"`
}

type ActFV1DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type ActFV1DataOutcome struct {
	// Mandatory fields
	Conclusion ActFV1DataOutcomeConclusion `json:"conclusion"`

	// Optional fields
	Description string `json:"description,omitempty"`
}

type ActFV1DataOutcomeConclusion string

const (
	ActFV1DataOutcomeConclusion_Successful   ActFV1DataOutcomeConclusion = "SUCCESSFUL"
	ActFV1DataOutcomeConclusion_Unsuccessful ActFV1DataOutcomeConclusion = "UNSUCCESSFUL"
	ActFV1DataOutcomeConclusion_Failed       ActFV1DataOutcomeConclusion = "FAILED"
	ActFV1DataOutcomeConclusion_Aborted      ActFV1DataOutcomeConclusion = "ABORTED"
	ActFV1DataOutcomeConclusion_TimedOut     ActFV1DataOutcomeConclusion = "TIMED_OUT"
	ActFV1DataOutcomeConclusion_Inconclusive ActFV1DataOutcomeConclusion = "INCONCLUSIVE"
)

type ActFV1DataPersistentLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields

}

// ActFV1Links represents a slice of ActFV1Link values with helper methods
// for adding new links.
type ActFV1Links []ActFV1Link

var _ LinkFinder = &ActFV1Links{}

// Add adds a new link of the specified type to a target event.
func (links *ActFV1Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, ActFV1Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *ActFV1Links) AddByID(linkType string, target string) {
	*links = append(*links, ActFV1Link{Target: target, Type: linkType})
}

// FindAll returns the IDs of all links of the specified type, or an empty
// slice if no such links are found.
func (links ActFV1Links) FindAll(linkType string) []string {
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
func (links ActFV1Links) FindFirst(linkType string) string {
	for _, link := range links {
		if link.Type == linkType {
			return link.Target
		}
	}
	return ""
}

type ActFV1Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type ActFV1Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security ActFV1MetaSecurity `json:"security,omitempty"`
	Source   ActFV1MetaSource   `json:"source,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

type ActFV1MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM ActFV1MetaSecuritySDM `json:"sdm,omitempty"`
}

type ActFV1MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type ActFV1MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string                     `json:"domainId,omitempty"`
	Host       string                     `json:"host,omitempty"`
	Name       string                     `json:"name,omitempty"`
	Serializer ActFV1MetaSourceSerializer `json:"serializer,omitempty"`
	URI        string                     `json:"uri,omitempty"`
}

type ActFV1MetaSourceSerializer struct {
	// Mandatory fields
	ArtifactID string `json:"artifactId"`
	GroupID    string `json:"groupId"`
	Version    string `json:"version"`

	// Optional fields

}
