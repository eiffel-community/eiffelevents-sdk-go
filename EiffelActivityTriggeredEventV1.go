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

// NewActivityTriggeredV1 creates a new struct pointer that represents
// major version 1 of EiffelActivityTriggeredEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 1.x.x
// currently known by this SDK.
func NewActivityTriggeredV1(modifiers ...Modifier) (*ActivityTriggeredV1, error) {
	var event ActivityTriggeredV1
	event.Meta.Type = "EiffelActivityTriggeredEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][1].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new ActivityTriggeredV1: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *ActivityTriggeredV1) MarshalJSON() ([]byte, error) {
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
		Data  *ActTV1Data  `json:"data"`
		Links EventLinksV1 `json:"links"`
		Meta  *MetaV1      `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *ActivityTriggeredV1) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *ActivityTriggeredV1) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &ActivityTriggeredV1{}
var _ MetaTeller = &ActivityTriggeredV1{}

// ID returns the value of the meta.id field.
func (e ActivityTriggeredV1) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e ActivityTriggeredV1) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e ActivityTriggeredV1) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e ActivityTriggeredV1) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e ActivityTriggeredV1) DomainID() string {
	return e.Meta.Source.DomainID
}

type ActivityTriggeredV1 struct {
	// Mandatory fields
	Data  ActTV1Data   `json:"data"`
	Links EventLinksV1 `json:"links"`
	Meta  MetaV1       `json:"meta"`

	// Optional fields

}

type ActTV1Data struct {
	// Mandatory fields
	Name string `json:"name"`

	// Optional fields
	Categories    []string                `json:"categories,omitempty"`
	CustomData    []CustomDataV1          `json:"customData,omitempty"`
	ExecutionType ActTV1DataExecutionType `json:"executionType,omitempty"`
	Triggers      []ActTV1DataTrigger     `json:"triggers,omitempty"`
}

type ActTV1DataExecutionType string

const (
	ActTV1DataExecutionType_Manual        ActTV1DataExecutionType = "MANUAL"
	ActTV1DataExecutionType_SemiAutomated ActTV1DataExecutionType = "SEMI_AUTOMATED"
	ActTV1DataExecutionType_Automated     ActTV1DataExecutionType = "AUTOMATED"
	ActTV1DataExecutionType_Other         ActTV1DataExecutionType = "OTHER"
)

type ActTV1DataTrigger struct {
	// Mandatory fields
	Type ActTV1DataTriggerType `json:"type"`

	// Optional fields
	Description string `json:"description,omitempty"`
}

type ActTV1DataTriggerType string

const (
	ActTV1DataTriggerType_Manual       ActTV1DataTriggerType = "MANUAL"
	ActTV1DataTriggerType_EiffelEvent  ActTV1DataTriggerType = "EIFFEL_EVENT"
	ActTV1DataTriggerType_SourceChange ActTV1DataTriggerType = "SOURCE_CHANGE"
	ActTV1DataTriggerType_Timer        ActTV1DataTriggerType = "TIMER"
	ActTV1DataTriggerType_Other        ActTV1DataTriggerType = "OTHER"
)
