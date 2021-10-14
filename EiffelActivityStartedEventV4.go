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
func (e *ActivityStartedV4) MarshalJSON() ([]byte, error) {
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
		links = make([]ActSV4Link, 0)
	}
	s := struct {
		Data  *ActSV4Data  `json:"data"`
		Links []ActSV4Link `json:"links"`
		Meta  *ActSV4Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *ActivityStartedV4) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type ActivityStartedV4 struct {
	// Mandatory fields
	Data  ActSV4Data   `json:"data"`
	Links []ActSV4Link `json:"links"`
	Meta  ActSV4Meta   `json:"meta"`

	// Optional fields

}

type ActSV4Data struct {
	// Mandatory fields

	// Optional fields
	CustomData   []ActSV4DataCustomDatum `json:"customData,omitempty"`
	ExecutionURI string                  `json:"executionUri,omitempty"`
	LiveLogs     []ActSV4DataLiveLog     `json:"liveLogs,omitempty"`
}

type ActSV4DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type ActSV4DataLiveLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields
	MediaType string   `json:"mediaType,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

type ActSV4Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields
	DomainID string `json:"domainId,omitempty"`
}

type ActSV4Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security ActSV4MetaSecurity `json:"security,omitempty"`
	Source   ActSV4MetaSource   `json:"source,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

type ActSV4MetaSecurity struct {
	// Mandatory fields
	AuthorIdentity string `json:"authorIdentity"`

	// Optional fields
	IntegrityProtection ActSV4MetaSecurityIntegrityProtection  `json:"integrityProtection,omitempty"`
	SequenceProtection  []ActSV4MetaSecuritySequenceProtection `json:"sequenceProtection,omitempty"`
}

type ActSV4MetaSecurityIntegrityProtection struct {
	// Mandatory fields
	Alg       ActSV4MetaSecurityIntegrityProtectionAlg `json:"alg"`
	Signature string                                   `json:"signature"`

	// Optional fields
	PublicKey string `json:"publicKey,omitempty"`
}

type ActSV4MetaSecurityIntegrityProtectionAlg string

const (
	ActSV4MetaSecurityIntegrityProtectionAlg_HS256 ActSV4MetaSecurityIntegrityProtectionAlg = "HS256"
	ActSV4MetaSecurityIntegrityProtectionAlg_HS384 ActSV4MetaSecurityIntegrityProtectionAlg = "HS384"
	ActSV4MetaSecurityIntegrityProtectionAlg_HS512 ActSV4MetaSecurityIntegrityProtectionAlg = "HS512"
	ActSV4MetaSecurityIntegrityProtectionAlg_RS256 ActSV4MetaSecurityIntegrityProtectionAlg = "RS256"
	ActSV4MetaSecurityIntegrityProtectionAlg_RS384 ActSV4MetaSecurityIntegrityProtectionAlg = "RS384"
	ActSV4MetaSecurityIntegrityProtectionAlg_RS512 ActSV4MetaSecurityIntegrityProtectionAlg = "RS512"
	ActSV4MetaSecurityIntegrityProtectionAlg_ES256 ActSV4MetaSecurityIntegrityProtectionAlg = "ES256"
	ActSV4MetaSecurityIntegrityProtectionAlg_ES384 ActSV4MetaSecurityIntegrityProtectionAlg = "ES384"
	ActSV4MetaSecurityIntegrityProtectionAlg_ES512 ActSV4MetaSecurityIntegrityProtectionAlg = "ES512"
	ActSV4MetaSecurityIntegrityProtectionAlg_PS256 ActSV4MetaSecurityIntegrityProtectionAlg = "PS256"
	ActSV4MetaSecurityIntegrityProtectionAlg_PS384 ActSV4MetaSecurityIntegrityProtectionAlg = "PS384"
	ActSV4MetaSecurityIntegrityProtectionAlg_PS512 ActSV4MetaSecurityIntegrityProtectionAlg = "PS512"
)

type ActSV4MetaSecuritySequenceProtection struct {
	// Mandatory fields
	Position     int64  `json:"position"`
	SequenceName string `json:"sequenceName"`

	// Optional fields

}

type ActSV4MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
