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

// Code generated by eventgen. DO NOT EDIT.

package eiffelevents

import (
	"fmt"
	"reflect"
	"time"

	"github.com/clarketm/json"
	"github.com/google/uuid"
)

// NewIssueDefinedV2 creates a new struct pointer that represents
// major version 2 of EiffelIssueDefinedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 2.x.x
// currently known by this SDK.
func NewIssueDefinedV2(modifiers ...Modifier) (*IssueDefinedV2, error) {
	var event IssueDefinedV2
	event.Meta.Type = "EiffelIssueDefinedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][2].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new IssueDefinedV2: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *IssueDefinedV2) MarshalJSON() ([]byte, error) {
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
		Data  *IDV2Data    `json:"data"`
		Links EventLinksV1 `json:"links"`
		Meta  *MetaV2      `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *IssueDefinedV2) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *IssueDefinedV2) String() string {
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
	_ CapabilityTeller = &IssueDefinedV2{}
	_ FieldSetter      = &IssueDefinedV2{}
	_ MetaTeller       = &IssueDefinedV2{}
)

// ID returns the value of the meta.id field.
func (e IssueDefinedV2) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e IssueDefinedV2) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e IssueDefinedV2) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e IssueDefinedV2) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e IssueDefinedV2) DomainID() string {
	return e.Meta.Source.DomainID
}

// SupportsSigning returns true if the event supports signatures according
// to V3 of the meta field, i.e. events where the signature is found under
// meta.security.integrityProtection.
func (e IssueDefinedV2) SupportsSigning() bool {
	return false
}

type IssueDefinedV2 struct {
	// Mandatory fields

	// Optional fields
	Data  IDV2Data     `json:"data,omitempty"`
	Links EventLinksV1 `json:"links,omitempty"`
	Meta  MetaV2       `json:"meta,omitempty"`
}

type IDV2Data struct {
	// Mandatory fields
	ID      string       `json:"id"`
	Tracker string       `json:"tracker"`
	Type    IDV2DataType `json:"type"`
	URI     string       `json:"uri"`

	// Optional fields
	CustomData []CustomDataV1 `json:"customData,omitempty"`
	Title      string         `json:"title,omitempty"`
}

type IDV2DataType string

const (
	IDV2DataType_Bug         IDV2DataType = "BUG"
	IDV2DataType_Improvement IDV2DataType = "IMPROVEMENT"
	IDV2DataType_Feature     IDV2DataType = "FEATURE"
	IDV2DataType_WorkItem    IDV2DataType = "WORK_ITEM"
	IDV2DataType_Requirement IDV2DataType = "REQUIREMENT"
	IDV2DataType_Other       IDV2DataType = "OTHER"
)
