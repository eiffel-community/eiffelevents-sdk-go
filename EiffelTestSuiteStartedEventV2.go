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

// NewTestSuiteStartedV2 creates a new struct pointer that represents
// major version 2 of EiffelTestSuiteStartedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 2.x.x
// currently known by this SDK.
func NewTestSuiteStartedV2(modifiers ...Modifier) (*TestSuiteStartedV2, error) {
	var event TestSuiteStartedV2
	event.Meta.Type = "EiffelTestSuiteStartedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][2].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new TestSuiteStartedV2: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *TestSuiteStartedV2) MarshalJSON() ([]byte, error) {
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
		links = make([]TSSV2Link, 0)
	}
	s := struct {
		Data  *TSSV2Data  `json:"data"`
		Links []TSSV2Link `json:"links"`
		Meta  *TSSV2Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *TestSuiteStartedV2) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *TestSuiteStartedV2) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &TestSuiteStartedV2{}
var _ MetaTeller = &TestSuiteStartedV2{}

// ID returns the value of the meta.id field.
func (e TestSuiteStartedV2) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e TestSuiteStartedV2) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e TestSuiteStartedV2) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e TestSuiteStartedV2) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e TestSuiteStartedV2) DomainID() string {
	return e.Meta.Source.DomainID
}

type TestSuiteStartedV2 struct {
	// Mandatory fields
	Data  TSSV2Data  `json:"data"`
	Links TSSV2Links `json:"links"`
	Meta  TSSV2Meta  `json:"meta"`

	// Optional fields

}

type TSSV2Data struct {
	// Mandatory fields
	Name string `json:"name"`

	// Optional fields
	Categories []string               `json:"categories,omitempty"`
	CustomData []TSSV2DataCustomDatum `json:"customData,omitempty"`
	LiveLogs   []TSSV2DataLiveLog     `json:"liveLogs,omitempty"`
	Types      []TSSV2DataType        `json:"types,omitempty"`
}

type TSSV2DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type TSSV2DataLiveLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields

}

type TSSV2DataType string

const (
	TSSV2DataType_Accessibility    TSSV2DataType = "ACCESSIBILITY"
	TSSV2DataType_BackupRecovery   TSSV2DataType = "BACKUP_RECOVERY"
	TSSV2DataType_Compatibility    TSSV2DataType = "COMPATIBILITY"
	TSSV2DataType_Conversion       TSSV2DataType = "CONVERSION"
	TSSV2DataType_DisasterRecovery TSSV2DataType = "DISASTER_RECOVERY"
	TSSV2DataType_Functional       TSSV2DataType = "FUNCTIONAL"
	TSSV2DataType_Installability   TSSV2DataType = "INSTALLABILITY"
	TSSV2DataType_Interoperability TSSV2DataType = "INTEROPERABILITY"
	TSSV2DataType_Localization     TSSV2DataType = "LOCALIZATION"
	TSSV2DataType_Maintainability  TSSV2DataType = "MAINTAINABILITY"
	TSSV2DataType_Performance      TSSV2DataType = "PERFORMANCE"
	TSSV2DataType_Portability      TSSV2DataType = "PORTABILITY"
	TSSV2DataType_Procedure        TSSV2DataType = "PROCEDURE"
	TSSV2DataType_Reliability      TSSV2DataType = "RELIABILITY"
	TSSV2DataType_Security         TSSV2DataType = "SECURITY"
	TSSV2DataType_Stability        TSSV2DataType = "STABILITY"
	TSSV2DataType_Usability        TSSV2DataType = "USABILITY"
)

// TSSV2Links represents a slice of TSSV2Link values with helper methods
// for adding new links.
type TSSV2Links []TSSV2Link

var _ LinkFinder = &TSSV2Links{}

// Add adds a new link of the specified type to a target event.
func (links *TSSV2Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, TSSV2Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *TSSV2Links) AddByID(linkType string, target string) {
	*links = append(*links, TSSV2Link{Target: target, Type: linkType})
}

// FindAll returns the IDs of all links of the specified type, or an empty
// slice if no such links are found.
func (links TSSV2Links) FindAll(linkType string) []string {
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
func (links TSSV2Links) FindFirst(linkType string) string {
	for _, link := range links {
		if link.Type == linkType {
			return link.Target
		}
	}
	return ""
}

type TSSV2Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type TSSV2Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security TSSV2MetaSecurity `json:"security,omitempty"`
	Source   TSSV2MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type TSSV2MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM TSSV2MetaSecuritySDM `json:"sdm,omitempty"`
}

type TSSV2MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type TSSV2MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
