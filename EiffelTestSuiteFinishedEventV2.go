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
func (e *TestSuiteFinishedV2) MarshalJSON() ([]byte, error) {
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
		links = make([]TSFV2Link, 0)
	}
	s := struct {
		Data  *TSFV2Data  `json:"data"`
		Links []TSFV2Link `json:"links"`
		Meta  *TSFV2Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *TestSuiteFinishedV2) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type TestSuiteFinishedV2 struct {
	// Mandatory fields
	Data  TSFV2Data   `json:"data"`
	Links []TSFV2Link `json:"links"`
	Meta  TSFV2Meta   `json:"meta"`

	// Optional fields

}

type TSFV2Data struct {
	// Mandatory fields

	// Optional fields
	CustomData     []TSFV2DataCustomDatum   `json:"customData,omitempty"`
	Outcome        TSFV2DataOutcome         `json:"outcome,omitempty"`
	PersistentLogs []TSFV2DataPersistentLog `json:"persistentLogs,omitempty"`
}

type TSFV2DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type TSFV2DataOutcome struct {
	// Mandatory fields

	// Optional fields
	Conclusion  TSFV2DataOutcomeConclusion `json:"conclusion,omitempty"`
	Description string                     `json:"description,omitempty"`
	Verdict     TSFV2DataOutcomeVerdict    `json:"verdict,omitempty"`
}

type TSFV2DataOutcomeConclusion string

const (
	TSFV2DataOutcomeConclusion_Successful   TSFV2DataOutcomeConclusion = "SUCCESSFUL"
	TSFV2DataOutcomeConclusion_Failed       TSFV2DataOutcomeConclusion = "FAILED"
	TSFV2DataOutcomeConclusion_Aborted      TSFV2DataOutcomeConclusion = "ABORTED"
	TSFV2DataOutcomeConclusion_TimedOut     TSFV2DataOutcomeConclusion = "TIMED_OUT"
	TSFV2DataOutcomeConclusion_Inconclusive TSFV2DataOutcomeConclusion = "INCONCLUSIVE"
)

type TSFV2DataOutcomeVerdict string

const (
	TSFV2DataOutcomeVerdict_Passed       TSFV2DataOutcomeVerdict = "PASSED"
	TSFV2DataOutcomeVerdict_Failed       TSFV2DataOutcomeVerdict = "FAILED"
	TSFV2DataOutcomeVerdict_Inconclusive TSFV2DataOutcomeVerdict = "INCONCLUSIVE"
)

type TSFV2DataPersistentLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields

}

type TSFV2Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type TSFV2Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security TSFV2MetaSecurity `json:"security,omitempty"`
	Source   TSFV2MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type TSFV2MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM TSFV2MetaSecuritySDM `json:"sdm,omitempty"`
}

type TSFV2MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type TSFV2MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
