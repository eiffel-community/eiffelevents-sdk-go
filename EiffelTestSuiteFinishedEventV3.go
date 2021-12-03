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

// NewTestSuiteFinishedV3 creates a new struct pointer that represents
// major version 3 of EiffelTestSuiteFinishedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 3.x.x
// currently known by this SDK.
func NewTestSuiteFinishedV3(modifiers ...Modifier) (*TestSuiteFinishedV3, error) {
	var event TestSuiteFinishedV3
	event.Meta.Type = "EiffelTestSuiteFinishedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][3].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new TestSuiteFinishedV3: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *TestSuiteFinishedV3) MarshalJSON() ([]byte, error) {
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
		links = make([]TSFV3Link, 0)
	}
	s := struct {
		Data  *TSFV3Data  `json:"data"`
		Links []TSFV3Link `json:"links"`
		Meta  *TSFV3Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *TestSuiteFinishedV3) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *TestSuiteFinishedV3) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &TestSuiteFinishedV3{}

type TestSuiteFinishedV3 struct {
	// Mandatory fields
	Data  TSFV3Data   `json:"data"`
	Links []TSFV3Link `json:"links"`
	Meta  TSFV3Meta   `json:"meta"`

	// Optional fields

}

type TSFV3Data struct {
	// Mandatory fields

	// Optional fields
	CustomData     []TSFV3DataCustomDatum   `json:"customData,omitempty"`
	Outcome        TSFV3DataOutcome         `json:"outcome,omitempty"`
	PersistentLogs []TSFV3DataPersistentLog `json:"persistentLogs,omitempty"`
}

type TSFV3DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type TSFV3DataOutcome struct {
	// Mandatory fields

	// Optional fields
	Conclusion  TSFV3DataOutcomeConclusion `json:"conclusion,omitempty"`
	Description string                     `json:"description,omitempty"`
	Verdict     TSFV3DataOutcomeVerdict    `json:"verdict,omitempty"`
}

type TSFV3DataOutcomeConclusion string

const (
	TSFV3DataOutcomeConclusion_Successful   TSFV3DataOutcomeConclusion = "SUCCESSFUL"
	TSFV3DataOutcomeConclusion_Failed       TSFV3DataOutcomeConclusion = "FAILED"
	TSFV3DataOutcomeConclusion_Aborted      TSFV3DataOutcomeConclusion = "ABORTED"
	TSFV3DataOutcomeConclusion_TimedOut     TSFV3DataOutcomeConclusion = "TIMED_OUT"
	TSFV3DataOutcomeConclusion_Inconclusive TSFV3DataOutcomeConclusion = "INCONCLUSIVE"
)

type TSFV3DataOutcomeVerdict string

const (
	TSFV3DataOutcomeVerdict_Passed       TSFV3DataOutcomeVerdict = "PASSED"
	TSFV3DataOutcomeVerdict_Failed       TSFV3DataOutcomeVerdict = "FAILED"
	TSFV3DataOutcomeVerdict_Inconclusive TSFV3DataOutcomeVerdict = "INCONCLUSIVE"
)

type TSFV3DataPersistentLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields
	MediaType string   `json:"mediaType,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

type TSFV3Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields
	DomainID string `json:"domainId,omitempty"`
}

type TSFV3Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security TSFV3MetaSecurity `json:"security,omitempty"`
	Source   TSFV3MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type TSFV3MetaSecurity struct {
	// Mandatory fields
	AuthorIdentity string `json:"authorIdentity"`

	// Optional fields
	IntegrityProtection TSFV3MetaSecurityIntegrityProtection  `json:"integrityProtection,omitempty"`
	SequenceProtection  []TSFV3MetaSecuritySequenceProtection `json:"sequenceProtection,omitempty"`
}

type TSFV3MetaSecurityIntegrityProtection struct {
	// Mandatory fields
	Alg       TSFV3MetaSecurityIntegrityProtectionAlg `json:"alg"`
	Signature string                                  `json:"signature"`

	// Optional fields
	PublicKey string `json:"publicKey,omitempty"`
}

type TSFV3MetaSecurityIntegrityProtectionAlg string

const (
	TSFV3MetaSecurityIntegrityProtectionAlg_HS256 TSFV3MetaSecurityIntegrityProtectionAlg = "HS256"
	TSFV3MetaSecurityIntegrityProtectionAlg_HS384 TSFV3MetaSecurityIntegrityProtectionAlg = "HS384"
	TSFV3MetaSecurityIntegrityProtectionAlg_HS512 TSFV3MetaSecurityIntegrityProtectionAlg = "HS512"
	TSFV3MetaSecurityIntegrityProtectionAlg_RS256 TSFV3MetaSecurityIntegrityProtectionAlg = "RS256"
	TSFV3MetaSecurityIntegrityProtectionAlg_RS384 TSFV3MetaSecurityIntegrityProtectionAlg = "RS384"
	TSFV3MetaSecurityIntegrityProtectionAlg_RS512 TSFV3MetaSecurityIntegrityProtectionAlg = "RS512"
	TSFV3MetaSecurityIntegrityProtectionAlg_ES256 TSFV3MetaSecurityIntegrityProtectionAlg = "ES256"
	TSFV3MetaSecurityIntegrityProtectionAlg_ES384 TSFV3MetaSecurityIntegrityProtectionAlg = "ES384"
	TSFV3MetaSecurityIntegrityProtectionAlg_ES512 TSFV3MetaSecurityIntegrityProtectionAlg = "ES512"
	TSFV3MetaSecurityIntegrityProtectionAlg_PS256 TSFV3MetaSecurityIntegrityProtectionAlg = "PS256"
	TSFV3MetaSecurityIntegrityProtectionAlg_PS384 TSFV3MetaSecurityIntegrityProtectionAlg = "PS384"
	TSFV3MetaSecurityIntegrityProtectionAlg_PS512 TSFV3MetaSecurityIntegrityProtectionAlg = "PS512"
)

type TSFV3MetaSecuritySequenceProtection struct {
	// Mandatory fields
	Position     int64  `json:"position"`
	SequenceName string `json:"sequenceName"`

	// Optional fields

}

type TSFV3MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
