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

// NewAnnouncementPublishedV3 creates a new struct pointer that represents
// major version 3 of EiffelAnnouncementPublishedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 3.x.x
// currently known by this SDK.
func NewAnnouncementPublishedV3(modifiers ...Modifier) (*AnnouncementPublishedV3, error) {
	var event AnnouncementPublishedV3
	event.Meta.Type = "EiffelAnnouncementPublishedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][3].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new AnnouncementPublishedV3: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *AnnouncementPublishedV3) MarshalJSON() ([]byte, error) {
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
		links = make(EventLinksV1, 0)
	}
	s := struct {
		Data  *AnnPV3Data  `json:"data"`
		Links EventLinksV1 `json:"links"`
		Meta  *MetaV3      `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *AnnouncementPublishedV3) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *AnnouncementPublishedV3) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var _ FieldSetter = &AnnouncementPublishedV3{}
var _ MetaTeller = &AnnouncementPublishedV3{}

// ID returns the value of the meta.id field.
func (e AnnouncementPublishedV3) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e AnnouncementPublishedV3) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e AnnouncementPublishedV3) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e AnnouncementPublishedV3) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e AnnouncementPublishedV3) DomainID() string {
	return e.Meta.Source.DomainID
}

type AnnouncementPublishedV3 struct {
	// Mandatory fields
	Data  AnnPV3Data   `json:"data"`
	Links EventLinksV1 `json:"links"`
	Meta  MetaV3       `json:"meta"`

	// Optional fields

}

type AnnPV3Data struct {
	// Mandatory fields
	Body     string             `json:"body"`
	Heading  string             `json:"heading"`
	Severity AnnPV3DataSeverity `json:"severity"`

	// Optional fields
	CustomData []CustomDataV1 `json:"customData,omitempty"`
	URI        string         `json:"uri,omitempty"`
}

type AnnPV3DataSeverity string

const (
	AnnPV3DataSeverity_Minor    AnnPV3DataSeverity = "MINOR"
	AnnPV3DataSeverity_Major    AnnPV3DataSeverity = "MAJOR"
	AnnPV3DataSeverity_Critical AnnPV3DataSeverity = "CRITICAL"
	AnnPV3DataSeverity_Blocker  AnnPV3DataSeverity = "BLOCKER"
	AnnPV3DataSeverity_Closed   AnnPV3DataSeverity = "CLOSED"
	AnnPV3DataSeverity_Canceled AnnPV3DataSeverity = "CANCELED"
)
