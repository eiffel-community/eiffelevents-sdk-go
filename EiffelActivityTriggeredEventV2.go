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

// NewActivityTriggeredV2 creates a new struct pointer that represents
// major version 2 of EiffelActivityTriggeredEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 2.x.x
// currently known by this SDK.
func NewActivityTriggeredV2(modifiers ...Modifier) (*ActivityTriggeredV2, error) {
	var event ActivityTriggeredV2
	event.Meta.Type = "EiffelActivityTriggeredEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][2].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new ActivityTriggeredV2: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *ActivityTriggeredV2) MarshalJSON() ([]byte, error) {
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
		links = make([]ActTV2Link, 0)
	}
	s := struct {
		Data  *ActTV2Data  `json:"data"`
		Links []ActTV2Link `json:"links"`
		Meta  *ActTV2Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *ActivityTriggeredV2) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *ActivityTriggeredV2) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &ActivityTriggeredV2{}
var _ MetaTeller = &ActivityTriggeredV2{}

// ID returns the value of the meta.id field.
func (e ActivityTriggeredV2) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e ActivityTriggeredV2) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e ActivityTriggeredV2) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e ActivityTriggeredV2) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e ActivityTriggeredV2) DomainID() string {
	return e.Meta.Source.DomainID
}

type ActivityTriggeredV2 struct {
	// Mandatory fields
	Data  ActTV2Data  `json:"data"`
	Links ActTV2Links `json:"links"`
	Meta  ActTV2Meta  `json:"meta"`

	// Optional fields

}

type ActTV2Data struct {
	// Mandatory fields
	Name string `json:"name"`

	// Optional fields
	Categories    []string                `json:"categories,omitempty"`
	CustomData    []ActTV2DataCustomDatum `json:"customData,omitempty"`
	ExecutionType ActTV2DataExecutionType `json:"executionType,omitempty"`
	Triggers      []ActTV2DataTrigger     `json:"triggers,omitempty"`
}

type ActTV2DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type ActTV2DataExecutionType string

const (
	ActTV2DataExecutionType_Manual        ActTV2DataExecutionType = "MANUAL"
	ActTV2DataExecutionType_SemiAutomated ActTV2DataExecutionType = "SEMI_AUTOMATED"
	ActTV2DataExecutionType_Automated     ActTV2DataExecutionType = "AUTOMATED"
	ActTV2DataExecutionType_Other         ActTV2DataExecutionType = "OTHER"
)

type ActTV2DataTrigger struct {
	// Mandatory fields
	Type ActTV2DataTriggerType `json:"type"`

	// Optional fields
	Description string `json:"description,omitempty"`
}

type ActTV2DataTriggerType string

const (
	ActTV2DataTriggerType_Manual       ActTV2DataTriggerType = "MANUAL"
	ActTV2DataTriggerType_EiffelEvent  ActTV2DataTriggerType = "EIFFEL_EVENT"
	ActTV2DataTriggerType_SourceChange ActTV2DataTriggerType = "SOURCE_CHANGE"
	ActTV2DataTriggerType_Timer        ActTV2DataTriggerType = "TIMER"
	ActTV2DataTriggerType_Other        ActTV2DataTriggerType = "OTHER"
)

// ActTV2Links represents a slice of ActTV2Link values with helper methods
// for adding new links.
type ActTV2Links []ActTV2Link

var _ LinkFinder = &ActTV2Links{}

// Add adds a new link of the specified type to a target event.
func (links *ActTV2Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, ActTV2Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *ActTV2Links) AddByID(linkType string, target string) {
	*links = append(*links, ActTV2Link{Target: target, Type: linkType})
}

// FindAll returns the IDs of all links of the specified type, or an empty
// slice if no such links are found.
func (links ActTV2Links) FindAll(linkType string) []string {
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
func (links ActTV2Links) FindFirst(linkType string) string {
	for _, link := range links {
		if link.Type == linkType {
			return link.Target
		}
	}
	return ""
}

type ActTV2Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type ActTV2Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security ActTV2MetaSecurity `json:"security,omitempty"`
	Source   ActTV2MetaSource   `json:"source,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

type ActTV2MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM ActTV2MetaSecuritySDM `json:"sdm,omitempty"`
}

type ActTV2MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type ActTV2MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
