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

// NewSourceChangeCreatedV2 creates a new struct pointer that represents
// major version 2 of EiffelSourceChangeCreatedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 2.x.x
// currently known by this SDK.
func NewSourceChangeCreatedV2(modifiers ...Modifier) (*SourceChangeCreatedV2, error) {
	var event SourceChangeCreatedV2
	event.Meta.Type = "EiffelSourceChangeCreatedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][2].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new SourceChangeCreatedV2: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *SourceChangeCreatedV2) MarshalJSON() ([]byte, error) {
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
		links = make([]SCCV2Link, 0)
	}
	s := struct {
		Data  *SCCV2Data  `json:"data"`
		Links []SCCV2Link `json:"links"`
		Meta  *SCCV2Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *SourceChangeCreatedV2) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *SourceChangeCreatedV2) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &SourceChangeCreatedV2{}
var _ MetaTeller = &SourceChangeCreatedV2{}

// ID returns the value of the meta.id field.
func (e SourceChangeCreatedV2) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e SourceChangeCreatedV2) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e SourceChangeCreatedV2) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e SourceChangeCreatedV2) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e SourceChangeCreatedV2) DomainID() string {
	return e.Meta.Source.DomainID
}

type SourceChangeCreatedV2 struct {
	// Mandatory fields
	Data  SCCV2Data  `json:"data"`
	Links SCCV2Links `json:"links"`
	Meta  SCCV2Meta  `json:"meta"`

	// Optional fields

}

type SCCV2Data struct {
	// Mandatory fields

	// Optional fields
	Author                SCCV2DataAuthor                `json:"author,omitempty"`
	CcCompositeIdentifier SCCV2DataCcCompositeIdentifier `json:"ccCompositeIdentifier,omitempty"`
	Change                SCCV2DataChange                `json:"change,omitempty"`
	CustomData            []SCCV2DataCustomDatum         `json:"customData,omitempty"`
	GitIdentifier         SCCV2DataGitIdentifier         `json:"gitIdentifier,omitempty"`
	HgIdentifier          SCCV2DataHgIdentifier          `json:"hgIdentifier,omitempty"`
	SvnIdentifier         SCCV2DataSvnIdentifier         `json:"svnIdentifier,omitempty"`
}

type SCCV2DataAuthor struct {
	// Mandatory fields

	// Optional fields
	Email string `json:"email,omitempty"`
	Group string `json:"group,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}

type SCCV2DataCcCompositeIdentifier struct {
	// Mandatory fields
	Branch     string   `json:"branch"`
	ConfigSpec string   `json:"configSpec"`
	Vobs       []string `json:"vobs"`

	// Optional fields

}

type SCCV2DataChange struct {
	// Mandatory fields

	// Optional fields
	Deletions  int64  `json:"deletions,omitempty"`
	Details    string `json:"details,omitempty"`
	Files      string `json:"files,omitempty"`
	ID         string `json:"id,omitempty"`
	Insertions int64  `json:"insertions,omitempty"`
	Tracker    string `json:"tracker,omitempty"`
}

type SCCV2DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type SCCV2DataGitIdentifier struct {
	// Mandatory fields
	CommitID string `json:"commitId"`
	RepoURI  string `json:"repoUri"`

	// Optional fields
	Branch   string `json:"branch,omitempty"`
	RepoName string `json:"repoName,omitempty"`
}

type SCCV2DataHgIdentifier struct {
	// Mandatory fields
	CommitID string `json:"commitId"`
	RepoURI  string `json:"repoUri"`

	// Optional fields
	Branch   string `json:"branch,omitempty"`
	RepoName string `json:"repoName,omitempty"`
}

type SCCV2DataSvnIdentifier struct {
	// Mandatory fields
	Directory string `json:"directory"`
	RepoURI   string `json:"repoUri"`
	Revision  int64  `json:"revision"`

	// Optional fields
	RepoName string `json:"repoName,omitempty"`
}

// SCCV2Links represents a slice of SCCV2Link values with helper methods
// for adding new links.
type SCCV2Links []SCCV2Link

var _ LinkFinder = &SCCV2Links{}

// Add adds a new link of the specified type to a target event.
func (links *SCCV2Links) Add(linkType string, target MetaTeller) {
	*links = append(*links, SCCV2Link{Target: target.ID(), Type: linkType})
}

// Add adds a new link of the specified type to a target event identified by an ID.
func (links *SCCV2Links) AddByID(linkType string, target string) {
	*links = append(*links, SCCV2Link{Target: target, Type: linkType})
}

// FindAll returns the IDs of all links of the specified type, or an empty
// slice if no such links are found.
func (links SCCV2Links) FindAll(linkType string) []string {
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
func (links SCCV2Links) FindFirst(linkType string) string {
	for _, link := range links {
		if link.Type == linkType {
			return link.Target
		}
	}
	return ""
}

type SCCV2Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type SCCV2Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security SCCV2MetaSecurity `json:"security,omitempty"`
	Source   SCCV2MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type SCCV2MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM SCCV2MetaSecuritySDM `json:"sdm,omitempty"`
}

type SCCV2MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type SCCV2MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string                    `json:"domainId,omitempty"`
	Host       string                    `json:"host,omitempty"`
	Name       string                    `json:"name,omitempty"`
	Serializer SCCV2MetaSourceSerializer `json:"serializer,omitempty"`
	URI        string                    `json:"uri,omitempty"`
}

type SCCV2MetaSourceSerializer struct {
	// Mandatory fields
	ArtifactID string `json:"artifactId"`
	GroupID    string `json:"groupId"`
	Version    string `json:"version"`

	// Optional fields

}
