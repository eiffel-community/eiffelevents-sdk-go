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
		links = make([]TSFV1Link, 0)
	}
	s := struct {
		Data  *TSFV1Data  `json:"data"`
		Links []TSFV1Link `json:"links"`
		Meta  *TSFV1Meta  `json:"meta"`
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

var _ FieldSetter = &TestSuiteFinishedV1{}
var _ MetaTeller = &TestSuiteFinishedV1{}

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

type TestSuiteFinishedV1 struct {
	// Mandatory fields
	Data  TSFV1Data  `json:"data"`
	Links TSFV1Links `json:"links"`
	Meta  TSFV1Meta  `json:"meta"`

	// Optional fields

}

type TSFV1Data struct {
	// Mandatory fields

	// Optional fields
	CustomData     []TSFV1DataCustomDatum   `json:"customData,omitempty"`
	Outcome        TSFV1DataOutcome         `json:"outcome,omitempty"`
	PersistentLogs []TSFV1DataPersistentLog `json:"persistentLogs,omitempty"`
}

type TSFV1DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

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

// TSFV1Links represents a slice of TSFV1Link values with helper methods
// for adding new links.
type TSFV1Links []TSFV1Link

// Add adds a new link of the specified type to a target event.
func (links *TSFV1Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, TSFV1Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *TSFV1Links) AddByID(linkType string, target string) {
	*links = append(*links, TSFV1Link{Target: target, Type: linkType})
}

type TSFV1Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type TSFV1Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security TSFV1MetaSecurity `json:"security,omitempty"`
	Source   TSFV1MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type TSFV1MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM TSFV1MetaSecuritySDM `json:"sdm,omitempty"`
}

type TSFV1MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type TSFV1MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string                    `json:"domainId,omitempty"`
	Host       string                    `json:"host,omitempty"`
	Name       string                    `json:"name,omitempty"`
	Serializer TSFV1MetaSourceSerializer `json:"serializer,omitempty"`
	URI        string                    `json:"uri,omitempty"`
}

type TSFV1MetaSourceSerializer struct {
	// Mandatory fields
	ArtifactID string `json:"artifactId"`
	GroupID    string `json:"groupId"`
	Version    string `json:"version"`

	// Optional fields

}
