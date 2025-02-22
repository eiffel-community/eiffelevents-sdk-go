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

// NewActivityTriggeredV4 creates a new struct pointer that represents
// major version 4 of EiffelActivityTriggeredEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 4.x.x
// currently known by this SDK.
func NewActivityTriggeredV4(modifiers ...Modifier) (*ActivityTriggeredV4, error) {
	var event ActivityTriggeredV4
	event.Meta.Type = "EiffelActivityTriggeredEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][4].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new ActivityTriggeredV4: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *ActivityTriggeredV4) MarshalJSON() ([]byte, error) {
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
		Data  *ActTV4Data  `json:"data"`
		Links EventLinksV1 `json:"links"`
		Meta  *MetaV3      `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *ActivityTriggeredV4) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *ActivityTriggeredV4) String() string {
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
	_ CapabilityTeller = &ActivityTriggeredV4{}
	_ FieldSetter      = &ActivityTriggeredV4{}
	_ MetaTeller       = &ActivityTriggeredV4{}
)

// ID returns the value of the meta.id field.
func (e ActivityTriggeredV4) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e ActivityTriggeredV4) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e ActivityTriggeredV4) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e ActivityTriggeredV4) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e ActivityTriggeredV4) DomainID() string {
	return e.Meta.Source.DomainID
}

// SupportsSigning returns true if the event supports signatures according
// to V3 of the meta field, i.e. events where the signature is found under
// meta.security.integrityProtection.
func (e ActivityTriggeredV4) SupportsSigning() bool {
	return true
}

type ActivityTriggeredV4 struct {
	// Mandatory fields
	Data  ActTV4Data   `json:"data"`
	Links EventLinksV1 `json:"links"`
	Meta  MetaV3       `json:"meta"`

	// Optional fields

}

type ActTV4Data struct {
	// Mandatory fields
	Name string `json:"name"`

	// Optional fields
	Categories    []string                `json:"categories,omitempty"`
	CustomData    []CustomDataV1          `json:"customData,omitempty"`
	ExecutionType ActTV4DataExecutionType `json:"executionType,omitempty"`
	Triggers      []ActTV4DataTrigger     `json:"triggers,omitempty"`
}

type ActTV4DataExecutionType string

const (
	ActTV4DataExecutionType_Manual        ActTV4DataExecutionType = "MANUAL"
	ActTV4DataExecutionType_SemiAutomated ActTV4DataExecutionType = "SEMI_AUTOMATED"
	ActTV4DataExecutionType_Automated     ActTV4DataExecutionType = "AUTOMATED"
	ActTV4DataExecutionType_Other         ActTV4DataExecutionType = "OTHER"
)

type ActTV4DataTrigger struct {
	// Mandatory fields
	Type ActTV4DataTriggerType `json:"type"`

	// Optional fields
	Description string `json:"description,omitempty"`
}

type ActTV4DataTriggerType string

const (
	ActTV4DataTriggerType_Manual       ActTV4DataTriggerType = "MANUAL"
	ActTV4DataTriggerType_EiffelEvent  ActTV4DataTriggerType = "EIFFEL_EVENT"
	ActTV4DataTriggerType_SourceChange ActTV4DataTriggerType = "SOURCE_CHANGE"
	ActTV4DataTriggerType_Timer        ActTV4DataTriggerType = "TIMER"
	ActTV4DataTriggerType_Other        ActTV4DataTriggerType = "OTHER"
)
