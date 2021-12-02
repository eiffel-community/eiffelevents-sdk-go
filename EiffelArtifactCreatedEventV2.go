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
	"time"

	"github.com/clarketm/json"
	"github.com/google/uuid"
)

// NewArtifactCreatedV2 creates a new struct pointer that represents
// major version 2 of EiffelArtifactCreatedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 2.x.x
// currently known by this SDK.
func NewArtifactCreatedV2() (*ArtifactCreatedV2, error) {
	var event ArtifactCreatedV2
	event.Meta.Type = "EiffelArtifactCreatedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][2].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *ArtifactCreatedV2) MarshalJSON() ([]byte, error) {
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
		links = make([]ArtCV2Link, 0)
	}
	s := struct {
		Data  *ArtCV2Data  `json:"data"`
		Links []ArtCV2Link `json:"links"`
		Meta  *ArtCV2Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *ArtifactCreatedV2) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type ArtifactCreatedV2 struct {
	// Mandatory fields
	Data  ArtCV2Data   `json:"data"`
	Links []ArtCV2Link `json:"links"`
	Meta  ArtCV2Meta   `json:"meta"`

	// Optional fields

}

type ArtCV2Data struct {
	// Mandatory fields
	Identity string `json:"identity"`

	// Optional fields
	BuildCommand           string                           `json:"buildCommand,omitempty"`
	CustomData             []ArtCV2DataCustomDatum          `json:"customData,omitempty"`
	DependsOn              []string                         `json:"dependsOn,omitempty"`
	FileInformation        []ArtCV2DataFileInformation      `json:"fileInformation,omitempty"`
	Implements             []string                         `json:"implements,omitempty"`
	Name                   string                           `json:"name,omitempty"`
	RequiresImplementation ArtCV2DataRequiresImplementation `json:"requiresImplementation,omitempty"`
}

type ArtCV2DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type ArtCV2DataFileInformation struct {
	// Mandatory fields
	Name string `json:"name"`

	// Optional fields
	Tags []string `json:"tags,omitempty"`
}

type ArtCV2DataRequiresImplementation string

const (
	ArtCV2DataRequiresImplementation_None       ArtCV2DataRequiresImplementation = "NONE"
	ArtCV2DataRequiresImplementation_Any        ArtCV2DataRequiresImplementation = "ANY"
	ArtCV2DataRequiresImplementation_ExactlyOne ArtCV2DataRequiresImplementation = "EXACTLY_ONE"
	ArtCV2DataRequiresImplementation_AtLeastOne ArtCV2DataRequiresImplementation = "AT_LEAST_ONE"
)

type ArtCV2Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type ArtCV2Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security ArtCV2MetaSecurity `json:"security,omitempty"`
	Source   ArtCV2MetaSource   `json:"source,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

type ArtCV2MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM ArtCV2MetaSecuritySDM `json:"sdm,omitempty"`
}

type ArtCV2MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type ArtCV2MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
