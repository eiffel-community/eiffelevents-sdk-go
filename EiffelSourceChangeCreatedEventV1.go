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
	"github.com/clarketm/json"
)

// MarshalJSON returns the JSON encoding of the event.
func (e *SourceChangeCreatedV1) MarshalJSON() ([]byte, error) {
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
		links = make([]SCCV1Link, 0)
	}
	s := struct {
		Data  *SCCV1Data  `json:"data"`
		Links []SCCV1Link `json:"links"`
		Meta  *SCCV1Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *SourceChangeCreatedV1) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type SourceChangeCreatedV1 struct {
	// Mandatory fields
	Data  SCCV1Data   `json:"data"`
	Links []SCCV1Link `json:"links"`
	Meta  SCCV1Meta   `json:"meta"`

	// Optional fields

}

type SCCV1Data struct {
	// Mandatory fields

	// Optional fields
	Author                SCCV1DataAuthor                `json:"author,omitempty"`
	CcCompositeIdentifier SCCV1DataCcCompositeIdentifier `json:"ccCompositeIdentifier,omitempty"`
	Change                SCCV1DataChange                `json:"change,omitempty"`
	CustomData            []SCCV1DataCustomDatum         `json:"customData,omitempty"`
	GitIdentifier         SCCV1DataGitIdentifier         `json:"gitIdentifier,omitempty"`
	HgIdentifier          SCCV1DataHgIdentifier          `json:"hgIdentifier,omitempty"`
	Issues                []interface{}                  `json:"issues,omitempty"`
	SvnIdentifier         SCCV1DataSvnIdentifier         `json:"svnIdentifier,omitempty"`
}

type SCCV1DataAuthor struct {
	// Mandatory fields

	// Optional fields
	Email string `json:"email,omitempty"`
	Group string `json:"group,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}

type SCCV1DataCcCompositeIdentifier struct {
	// Mandatory fields
	Branch     string   `json:"branch"`
	ConfigSpec string   `json:"configSpec"`
	Vobs       []string `json:"vobs"`

	// Optional fields

}

type SCCV1DataChange struct {
	// Mandatory fields

	// Optional fields
	Deletions  int64  `json:"deletions,omitempty"`
	Details    string `json:"details,omitempty"`
	Files      string `json:"files,omitempty"`
	ID         string `json:"id,omitempty"`
	Insertions int64  `json:"insertions,omitempty"`
	Tracker    string `json:"tracker,omitempty"`
}

type SCCV1DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type SCCV1DataGitIdentifier struct {
	// Mandatory fields
	CommitID string `json:"commitId"`
	RepoURI  string `json:"repoUri"`

	// Optional fields
	Branch   string `json:"branch,omitempty"`
	RepoName string `json:"repoName,omitempty"`
}

type SCCV1DataHgIdentifier struct {
	// Mandatory fields
	CommitID string `json:"commitId"`
	RepoURI  string `json:"repoUri"`

	// Optional fields
	Branch   string `json:"branch,omitempty"`
	RepoName string `json:"repoName,omitempty"`
}

type SCCV1DataSvnIdentifier struct {
	// Mandatory fields
	Directory string `json:"directory"`
	RepoURI   string `json:"repoUri"`
	Revision  int64  `json:"revision"`

	// Optional fields
	RepoName string `json:"repoName,omitempty"`
}

type SCCV1Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type SCCV1Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security SCCV1MetaSecurity `json:"security,omitempty"`
	Source   SCCV1MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type SCCV1MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM SCCV1MetaSecuritySDM `json:"sdm,omitempty"`
}

type SCCV1MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type SCCV1MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string                    `json:"domainId,omitempty"`
	Host       string                    `json:"host,omitempty"`
	Name       string                    `json:"name,omitempty"`
	Serializer SCCV1MetaSourceSerializer `json:"serializer,omitempty"`
	URI        string                    `json:"uri,omitempty"`
}

type SCCV1MetaSourceSerializer struct {
	// Mandatory fields
	ArtifactID string `json:"artifactId"`
	GroupID    string `json:"groupId"`
	Version    string `json:"version"`

	// Optional fields

}
