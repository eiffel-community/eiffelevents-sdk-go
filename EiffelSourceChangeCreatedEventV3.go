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

// NewSourceChangeCreatedV3 creates a new struct pointer that represents
// major version 3 of EiffelSourceChangeCreatedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 3.x.x
// currently known by this SDK.
func NewSourceChangeCreatedV3(modifiers ...Modifier) (*SourceChangeCreatedV3, error) {
	var event SourceChangeCreatedV3
	event.Meta.Type = "EiffelSourceChangeCreatedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][3].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new SourceChangeCreatedV3: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *SourceChangeCreatedV3) MarshalJSON() ([]byte, error) {
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
		links = make([]SCCV3Link, 0)
	}
	s := struct {
		Data  *SCCV3Data  `json:"data"`
		Links []SCCV3Link `json:"links"`
		Meta  *SCCV3Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *SourceChangeCreatedV3) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *SourceChangeCreatedV3) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &SourceChangeCreatedV3{}
var _ MetaTeller = &SourceChangeCreatedV3{}

// ID returns the value of the meta.id field.
func (e SourceChangeCreatedV3) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e SourceChangeCreatedV3) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e SourceChangeCreatedV3) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e SourceChangeCreatedV3) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e SourceChangeCreatedV3) DomainID() string {
	return e.Meta.Source.DomainID
}

type SourceChangeCreatedV3 struct {
	// Mandatory fields
	Data  SCCV3Data  `json:"data"`
	Links SCCV3Links `json:"links"`
	Meta  SCCV3Meta  `json:"meta"`

	// Optional fields

}

type SCCV3Data struct {
	// Mandatory fields

	// Optional fields
	Author                SCCV3DataAuthor                `json:"author,omitempty"`
	CcCompositeIdentifier SCCV3DataCcCompositeIdentifier `json:"ccCompositeIdentifier,omitempty"`
	Change                SCCV3DataChange                `json:"change,omitempty"`
	CustomData            []SCCV3DataCustomDatum         `json:"customData,omitempty"`
	GitIdentifier         SCCV3DataGitIdentifier         `json:"gitIdentifier,omitempty"`
	HgIdentifier          SCCV3DataHgIdentifier          `json:"hgIdentifier,omitempty"`
	SvnIdentifier         SCCV3DataSvnIdentifier         `json:"svnIdentifier,omitempty"`
}

type SCCV3DataAuthor struct {
	// Mandatory fields

	// Optional fields
	Email string `json:"email,omitempty"`
	Group string `json:"group,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}

type SCCV3DataCcCompositeIdentifier struct {
	// Mandatory fields
	Branch     string   `json:"branch"`
	ConfigSpec string   `json:"configSpec"`
	Vobs       []string `json:"vobs"`

	// Optional fields

}

type SCCV3DataChange struct {
	// Mandatory fields

	// Optional fields
	Deletions  int64  `json:"deletions,omitempty"`
	Details    string `json:"details,omitempty"`
	Files      string `json:"files,omitempty"`
	ID         string `json:"id,omitempty"`
	Insertions int64  `json:"insertions,omitempty"`
	Tracker    string `json:"tracker,omitempty"`
}

type SCCV3DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type SCCV3DataGitIdentifier struct {
	// Mandatory fields
	CommitID string `json:"commitId"`
	RepoURI  string `json:"repoUri"`

	// Optional fields
	Branch   string `json:"branch,omitempty"`
	RepoName string `json:"repoName,omitempty"`
}

type SCCV3DataHgIdentifier struct {
	// Mandatory fields
	CommitID string `json:"commitId"`
	RepoURI  string `json:"repoUri"`

	// Optional fields
	Branch   string `json:"branch,omitempty"`
	RepoName string `json:"repoName,omitempty"`
}

type SCCV3DataSvnIdentifier struct {
	// Mandatory fields
	Directory string `json:"directory"`
	RepoURI   string `json:"repoUri"`
	Revision  int64  `json:"revision"`

	// Optional fields
	RepoName string `json:"repoName,omitempty"`
}

// SCCV3Links represents a slice of SCCV3Link values with helper methods
// for adding new links.
type SCCV3Links []SCCV3Link

// Add adds a new link of the specified type to a target event.
func (links *SCCV3Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, SCCV3Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *SCCV3Links) AddByID(linkType string, target string) {
	*links = append(*links, SCCV3Link{Target: target, Type: linkType})
}

type SCCV3Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type SCCV3Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security SCCV3MetaSecurity `json:"security,omitempty"`
	Source   SCCV3MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type SCCV3MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM SCCV3MetaSecuritySDM `json:"sdm,omitempty"`
}

type SCCV3MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type SCCV3MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
