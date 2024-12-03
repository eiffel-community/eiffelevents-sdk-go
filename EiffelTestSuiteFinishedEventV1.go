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

// NewTestSuiteFinishedV1 creates a new struct pointer that represents
// major version 1 of EiffelTestSuiteFinishedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 1.x.x
// currently known by this SDK.
func NewTestSuiteFinishedV1(modifiers ...Modifier) (*TestSuiteFinishedV1, error) {
	var event TestSuiteFinishedV1
	event.Meta.Type = "EiffelTestSuiteFinishedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][1].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new TestSuiteFinishedV1: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *TestSuiteFinishedV1) MarshalJSON() ([]byte, error) {
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
		Data  *TSFV1Data   `json:"data"`
		Links EventLinksV1 `json:"links"`
		Meta  *MetaV1      `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *TestSuiteFinishedV1) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *TestSuiteFinishedV1) String() string {
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
	_ CapabilityTeller = &TestSuiteFinishedV1{}
	_ FieldSetter      = &TestSuiteFinishedV1{}
	_ MetaTeller       = &TestSuiteFinishedV1{}
)

// ID returns the value of the meta.id field.
func (e TestSuiteFinishedV1) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e TestSuiteFinishedV1) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e TestSuiteFinishedV1) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e TestSuiteFinishedV1) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e TestSuiteFinishedV1) DomainID() string {
	return e.Meta.Source.DomainID
}

// SupportsSigning returns true if the event supports signatures according
// to V3 of the meta field, i.e. events where the signature is found under
// meta.security.integrityProtection.
func (e TestSuiteFinishedV1) SupportsSigning() bool {
	return false
}

type TestSuiteFinishedV1 struct {
	// Mandatory fields
	Data  TSFV1Data    `json:"data"`
	Links EventLinksV1 `json:"links"`
	Meta  MetaV1       `json:"meta"`

	// Optional fields

}

type TSFV1Data struct {
	// Mandatory fields

	// Optional fields
	CustomData     []CustomDataV1           `json:"customData,omitempty"`
	Outcome        TSFV1DataOutcome         `json:"outcome,omitempty"`
	PersistentLogs []TSFV1DataPersistentLog `json:"persistentLogs,omitempty"`
}

type TSFV1DataOutcome struct {
	// Mandatory fields

	// Optional fields
	Conclusion  TSFV1DataOutcomeConclusion `json:"conclusion,omitempty"`
	Description string                     `json:"description,omitempty"`
	Verdict     TSFV1DataOutcomeVerdict    `json:"verdict,omitempty"`
}

type TSFV1DataOutcomeConclusion string

const (
	TSFV1DataOutcomeConclusion_Successful   TSFV1DataOutcomeConclusion = "SUCCESSFUL"
	TSFV1DataOutcomeConclusion_Failed       TSFV1DataOutcomeConclusion = "FAILED"
	TSFV1DataOutcomeConclusion_Aborted      TSFV1DataOutcomeConclusion = "ABORTED"
	TSFV1DataOutcomeConclusion_TimedOut     TSFV1DataOutcomeConclusion = "TIMED_OUT"
	TSFV1DataOutcomeConclusion_Inconclusive TSFV1DataOutcomeConclusion = "INCONCLUSIVE"
)

type TSFV1DataOutcomeVerdict string

const (
	TSFV1DataOutcomeVerdict_Passed       TSFV1DataOutcomeVerdict = "PASSED"
	TSFV1DataOutcomeVerdict_Failed       TSFV1DataOutcomeVerdict = "FAILED"
	TSFV1DataOutcomeVerdict_Inconclusive TSFV1DataOutcomeVerdict = "INCONCLUSIVE"
)

type TSFV1DataPersistentLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields

}
