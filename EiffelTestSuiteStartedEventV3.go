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

// NewTestSuiteStartedV3 creates a new struct pointer that represents
// major version 3 of EiffelTestSuiteStartedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 3.x.x
// currently known by this SDK.
func NewTestSuiteStartedV3(modifiers ...Modifier) (*TestSuiteStartedV3, error) {
	var event TestSuiteStartedV3
	event.Meta.Type = "EiffelTestSuiteStartedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][3].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new TestSuiteStartedV3: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *TestSuiteStartedV3) MarshalJSON() ([]byte, error) {
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
		links = make([]TSSV3Link, 0)
	}
	s := struct {
		Data  *TSSV3Data  `json:"data"`
		Links []TSSV3Link `json:"links"`
		Meta  *TSSV3Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *TestSuiteStartedV3) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *TestSuiteStartedV3) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &TestSuiteStartedV3{}
var _ MetaTeller = &TestSuiteStartedV3{}

// ID returns the value of the meta.id field.
func (e TestSuiteStartedV3) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e TestSuiteStartedV3) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e TestSuiteStartedV3) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e TestSuiteStartedV3) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e TestSuiteStartedV3) DomainID() string {
	return e.Meta.Source.DomainID
}

type TestSuiteStartedV3 struct {
	// Mandatory fields
	Data  TSSV3Data  `json:"data"`
	Links TSSV3Links `json:"links"`
	Meta  TSSV3Meta  `json:"meta"`

	// Optional fields

}

type TSSV3Data struct {
	// Mandatory fields
	Name string `json:"name"`

	// Optional fields
	Categories []string               `json:"categories,omitempty"`
	CustomData []TSSV3DataCustomDatum `json:"customData,omitempty"`
	LiveLogs   []TSSV3DataLiveLog     `json:"liveLogs,omitempty"`
	Types      []TSSV3DataType        `json:"types,omitempty"`
}

type TSSV3DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type TSSV3DataLiveLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields
	MediaType string   `json:"mediaType,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

type TSSV3DataType string

const (
	TSSV3DataType_Accessibility    TSSV3DataType = "ACCESSIBILITY"
	TSSV3DataType_BackupRecovery   TSSV3DataType = "BACKUP_RECOVERY"
	TSSV3DataType_Compatibility    TSSV3DataType = "COMPATIBILITY"
	TSSV3DataType_Conversion       TSSV3DataType = "CONVERSION"
	TSSV3DataType_DisasterRecovery TSSV3DataType = "DISASTER_RECOVERY"
	TSSV3DataType_Functional       TSSV3DataType = "FUNCTIONAL"
	TSSV3DataType_Installability   TSSV3DataType = "INSTALLABILITY"
	TSSV3DataType_Interoperability TSSV3DataType = "INTEROPERABILITY"
	TSSV3DataType_Localization     TSSV3DataType = "LOCALIZATION"
	TSSV3DataType_Maintainability  TSSV3DataType = "MAINTAINABILITY"
	TSSV3DataType_Performance      TSSV3DataType = "PERFORMANCE"
	TSSV3DataType_Portability      TSSV3DataType = "PORTABILITY"
	TSSV3DataType_Procedure        TSSV3DataType = "PROCEDURE"
	TSSV3DataType_Reliability      TSSV3DataType = "RELIABILITY"
	TSSV3DataType_Security         TSSV3DataType = "SECURITY"
	TSSV3DataType_Stability        TSSV3DataType = "STABILITY"
	TSSV3DataType_Usability        TSSV3DataType = "USABILITY"
)

// TSSV3Links represents a slice of TSSV3Link values with helper methods
// for adding new links.
type TSSV3Links []TSSV3Link

// Add adds a new link of the specified type to a target event.
func (links *TSSV3Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, TSSV3Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *TSSV3Links) AddByID(linkType string, target string) {
	*links = append(*links, TSSV3Link{Target: target, Type: linkType})
}

type TSSV3Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields
	DomainID string `json:"domainId,omitempty"`
}

type TSSV3Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security TSSV3MetaSecurity `json:"security,omitempty"`
	Source   TSSV3MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type TSSV3MetaSecurity struct {
	// Mandatory fields
	AuthorIdentity string `json:"authorIdentity"`

	// Optional fields
	IntegrityProtection TSSV3MetaSecurityIntegrityProtection  `json:"integrityProtection,omitempty"`
	SequenceProtection  []TSSV3MetaSecuritySequenceProtection `json:"sequenceProtection,omitempty"`
}

type TSSV3MetaSecurityIntegrityProtection struct {
	// Mandatory fields
	Alg       TSSV3MetaSecurityIntegrityProtectionAlg `json:"alg"`
	Signature string                                  `json:"signature"`

	// Optional fields
	PublicKey string `json:"publicKey,omitempty"`
}

type TSSV3MetaSecurityIntegrityProtectionAlg string

const (
	TSSV3MetaSecurityIntegrityProtectionAlg_HS256 TSSV3MetaSecurityIntegrityProtectionAlg = "HS256"
	TSSV3MetaSecurityIntegrityProtectionAlg_HS384 TSSV3MetaSecurityIntegrityProtectionAlg = "HS384"
	TSSV3MetaSecurityIntegrityProtectionAlg_HS512 TSSV3MetaSecurityIntegrityProtectionAlg = "HS512"
	TSSV3MetaSecurityIntegrityProtectionAlg_RS256 TSSV3MetaSecurityIntegrityProtectionAlg = "RS256"
	TSSV3MetaSecurityIntegrityProtectionAlg_RS384 TSSV3MetaSecurityIntegrityProtectionAlg = "RS384"
	TSSV3MetaSecurityIntegrityProtectionAlg_RS512 TSSV3MetaSecurityIntegrityProtectionAlg = "RS512"
	TSSV3MetaSecurityIntegrityProtectionAlg_ES256 TSSV3MetaSecurityIntegrityProtectionAlg = "ES256"
	TSSV3MetaSecurityIntegrityProtectionAlg_ES384 TSSV3MetaSecurityIntegrityProtectionAlg = "ES384"
	TSSV3MetaSecurityIntegrityProtectionAlg_ES512 TSSV3MetaSecurityIntegrityProtectionAlg = "ES512"
	TSSV3MetaSecurityIntegrityProtectionAlg_PS256 TSSV3MetaSecurityIntegrityProtectionAlg = "PS256"
	TSSV3MetaSecurityIntegrityProtectionAlg_PS384 TSSV3MetaSecurityIntegrityProtectionAlg = "PS384"
	TSSV3MetaSecurityIntegrityProtectionAlg_PS512 TSSV3MetaSecurityIntegrityProtectionAlg = "PS512"
)

type TSSV3MetaSecuritySequenceProtection struct {
	// Mandatory fields
	Position     int64  `json:"position"`
	SequenceName string `json:"sequenceName"`

	// Optional fields

}

type TSSV3MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
