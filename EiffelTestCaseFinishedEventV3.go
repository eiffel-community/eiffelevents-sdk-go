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

// NewTestCaseFinishedV3 creates a new struct pointer that represents
// major version 3 of EiffelTestCaseFinishedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 3.x.x
// currently known by this SDK.
func NewTestCaseFinishedV3(modifiers ...Modifier) (*TestCaseFinishedV3, error) {
	var event TestCaseFinishedV3
	event.Meta.Type = "EiffelTestCaseFinishedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][3].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new TestCaseFinishedV3: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *TestCaseFinishedV3) MarshalJSON() ([]byte, error) {
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
		links = make([]TCFV3Link, 0)
	}
	s := struct {
		Data  *TCFV3Data  `json:"data"`
		Links []TCFV3Link `json:"links"`
		Meta  *TCFV3Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *TestCaseFinishedV3) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *TestCaseFinishedV3) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &TestCaseFinishedV3{}
var _ MetaTeller = &TestCaseFinishedV3{}

// ID returns the value of the meta.id field.
func (e TestCaseFinishedV3) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e TestCaseFinishedV3) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e TestCaseFinishedV3) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e TestCaseFinishedV3) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e TestCaseFinishedV3) DomainID() string {
	return e.Meta.Source.DomainID
}

type TestCaseFinishedV3 struct {
	// Mandatory fields
	Data  TCFV3Data  `json:"data"`
	Links TCFV3Links `json:"links"`
	Meta  TCFV3Meta  `json:"meta"`

	// Optional fields

}

type TCFV3Data struct {
	// Mandatory fields
	Outcome TCFV3DataOutcome `json:"outcome"`

	// Optional fields
	CustomData     []TCFV3DataCustomDatum   `json:"customData,omitempty"`
	PersistentLogs []TCFV3DataPersistentLog `json:"persistentLogs,omitempty"`
}

type TCFV3DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type TCFV3DataOutcome struct {
	// Mandatory fields
	Conclusion TCFV3DataOutcomeConclusion `json:"conclusion"`
	Verdict    TCFV3DataOutcomeVerdict    `json:"verdict"`

	// Optional fields
	Description string                   `json:"description,omitempty"`
	Metrics     []TCFV3DataOutcomeMetric `json:"metrics,omitempty"`
}

type TCFV3DataOutcomeConclusion string

const (
	TCFV3DataOutcomeConclusion_Successful   TCFV3DataOutcomeConclusion = "SUCCESSFUL"
	TCFV3DataOutcomeConclusion_Failed       TCFV3DataOutcomeConclusion = "FAILED"
	TCFV3DataOutcomeConclusion_Aborted      TCFV3DataOutcomeConclusion = "ABORTED"
	TCFV3DataOutcomeConclusion_TimedOut     TCFV3DataOutcomeConclusion = "TIMED_OUT"
	TCFV3DataOutcomeConclusion_Inconclusive TCFV3DataOutcomeConclusion = "INCONCLUSIVE"
)

type TCFV3DataOutcomeMetric struct {
	// Mandatory fields
	Name  string      `json:"name"`
	Value interface{} `json:"value"`

	// Optional fields

}

type TCFV3DataOutcomeVerdict string

const (
	TCFV3DataOutcomeVerdict_Passed       TCFV3DataOutcomeVerdict = "PASSED"
	TCFV3DataOutcomeVerdict_Failed       TCFV3DataOutcomeVerdict = "FAILED"
	TCFV3DataOutcomeVerdict_Inconclusive TCFV3DataOutcomeVerdict = "INCONCLUSIVE"
)

type TCFV3DataPersistentLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields
	MediaType string   `json:"mediaType,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

// TCFV3Links represents a slice of TCFV3Link values with helper methods
// for adding new links.
type TCFV3Links []TCFV3Link

// Add adds a new link of the specified type to a target event.
func (links *TCFV3Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, TCFV3Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *TCFV3Links) AddByID(linkType string, target string) {
	*links = append(*links, TCFV3Link{Target: target, Type: linkType})
}

type TCFV3Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields
	DomainID string `json:"domainId,omitempty"`
}

type TCFV3Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security TCFV3MetaSecurity `json:"security,omitempty"`
	Source   TCFV3MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type TCFV3MetaSecurity struct {
	// Mandatory fields
	AuthorIdentity string `json:"authorIdentity"`

	// Optional fields
	IntegrityProtection TCFV3MetaSecurityIntegrityProtection  `json:"integrityProtection,omitempty"`
	SequenceProtection  []TCFV3MetaSecuritySequenceProtection `json:"sequenceProtection,omitempty"`
}

type TCFV3MetaSecurityIntegrityProtection struct {
	// Mandatory fields
	Alg       TCFV3MetaSecurityIntegrityProtectionAlg `json:"alg"`
	Signature string                                  `json:"signature"`

	// Optional fields
	PublicKey string `json:"publicKey,omitempty"`
}

type TCFV3MetaSecurityIntegrityProtectionAlg string

const (
	TCFV3MetaSecurityIntegrityProtectionAlg_HS256 TCFV3MetaSecurityIntegrityProtectionAlg = "HS256"
	TCFV3MetaSecurityIntegrityProtectionAlg_HS384 TCFV3MetaSecurityIntegrityProtectionAlg = "HS384"
	TCFV3MetaSecurityIntegrityProtectionAlg_HS512 TCFV3MetaSecurityIntegrityProtectionAlg = "HS512"
	TCFV3MetaSecurityIntegrityProtectionAlg_RS256 TCFV3MetaSecurityIntegrityProtectionAlg = "RS256"
	TCFV3MetaSecurityIntegrityProtectionAlg_RS384 TCFV3MetaSecurityIntegrityProtectionAlg = "RS384"
	TCFV3MetaSecurityIntegrityProtectionAlg_RS512 TCFV3MetaSecurityIntegrityProtectionAlg = "RS512"
	TCFV3MetaSecurityIntegrityProtectionAlg_ES256 TCFV3MetaSecurityIntegrityProtectionAlg = "ES256"
	TCFV3MetaSecurityIntegrityProtectionAlg_ES384 TCFV3MetaSecurityIntegrityProtectionAlg = "ES384"
	TCFV3MetaSecurityIntegrityProtectionAlg_ES512 TCFV3MetaSecurityIntegrityProtectionAlg = "ES512"
	TCFV3MetaSecurityIntegrityProtectionAlg_PS256 TCFV3MetaSecurityIntegrityProtectionAlg = "PS256"
	TCFV3MetaSecurityIntegrityProtectionAlg_PS384 TCFV3MetaSecurityIntegrityProtectionAlg = "PS384"
	TCFV3MetaSecurityIntegrityProtectionAlg_PS512 TCFV3MetaSecurityIntegrityProtectionAlg = "PS512"
)

type TCFV3MetaSecuritySequenceProtection struct {
	// Mandatory fields
	Position     int64  `json:"position"`
	SequenceName string `json:"sequenceName"`

	// Optional fields

}

type TCFV3MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
