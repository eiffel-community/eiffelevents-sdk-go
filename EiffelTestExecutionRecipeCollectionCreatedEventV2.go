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

// NewTestExecutionRecipeCollectionCreatedV2 creates a new struct pointer that represents
// major version 2 of EiffelTestExecutionRecipeCollectionCreatedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 2.x.x
// currently known by this SDK.
func NewTestExecutionRecipeCollectionCreatedV2(modifiers ...Modifier) (*TestExecutionRecipeCollectionCreatedV2, error) {
	var event TestExecutionRecipeCollectionCreatedV2
	event.Meta.Type = "EiffelTestExecutionRecipeCollectionCreatedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][2].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	for _, modifier := range modifiers {
		if err := modifier(&event); err != nil {
			return nil, fmt.Errorf("error applying modifier to new TestExecutionRecipeCollectionCreatedV2: %w", err)
		}
	}
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *TestExecutionRecipeCollectionCreatedV2) MarshalJSON() ([]byte, error) {
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
		Data  *TERCCV2Data `json:"data"`
		Links EventLinksV1 `json:"links"`
		Meta  *MetaV1      `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

func (e *TestExecutionRecipeCollectionCreatedV2) SetField(fieldName string, value interface{}) error {
	return setField(reflect.ValueOf(e), fieldName, value)
}

// String returns the JSON encoding of the event.
func (e *TestExecutionRecipeCollectionCreatedV2) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

var (
	_ CapabilityTeller = &TestExecutionRecipeCollectionCreatedV2{}
	_ FieldSetter      = &TestExecutionRecipeCollectionCreatedV2{}
	_ MetaTeller       = &TestExecutionRecipeCollectionCreatedV2{}
)

// ID returns the value of the meta.id field.
func (e TestExecutionRecipeCollectionCreatedV2) ID() string {
	return e.Meta.ID
}

// Type returns the value of the meta.type field.
func (e TestExecutionRecipeCollectionCreatedV2) Type() string {
	return e.Meta.Type
}

// Version returns the value of the meta.version field.
func (e TestExecutionRecipeCollectionCreatedV2) Version() string {
	return e.Meta.Version
}

// Time returns the value of the meta.time field.
func (e TestExecutionRecipeCollectionCreatedV2) Time() int64 {
	return e.Meta.Time
}

// DomainID returns the value of the meta.source.domainId field.
func (e TestExecutionRecipeCollectionCreatedV2) DomainID() string {
	return e.Meta.Source.DomainID
}

// SupportsSigning returns true if the event supports signatures according
// to V3 of the meta field, i.e. events where the signature is found under
// meta.security.integrityProtection.
func (e TestExecutionRecipeCollectionCreatedV2) SupportsSigning() bool {
	return false
}

type TestExecutionRecipeCollectionCreatedV2 struct {
	// Mandatory fields
	Data  TERCCV2Data  `json:"data"`
	Links EventLinksV1 `json:"links"`
	Meta  MetaV1       `json:"meta"`

	// Optional fields

}

type TERCCV2Data struct {
	// Mandatory fields
	SelectionStrategy TERCCV2DataSelectionStrategy `json:"selectionStrategy"`

	// Optional fields
	Batches    []TERCCV2DataBatch `json:"batches,omitempty"`
	BatchesURI string             `json:"batchesUri,omitempty"`
	CustomData []CustomDataV1     `json:"customData,omitempty"`
}

type TERCCV2DataBatch struct {
	// Mandatory fields
	Priority int64                    `json:"priority"`
	Recipes  []TERCCV2DataBatchRecipe `json:"recipes"`

	// Optional fields
	Dependencies []TERCCV2DataBatchDependency `json:"dependencies,omitempty"`
	Name         string                       `json:"name,omitempty"`
}

type TERCCV2DataBatchDependency struct {
	// Mandatory fields
	Dependency string `json:"dependency"`
	Dependent  string `json:"dependent"`

	// Optional fields

}

type TERCCV2DataBatchRecipe struct {
	// Mandatory fields
	ID       string                         `json:"id"`
	TestCase TERCCV2DataBatchRecipeTestCase `json:"testCase"`

	// Optional fields
	Constraints []TERCCV2DataBatchRecipeConstraint `json:"constraints,omitempty"`
}

type TERCCV2DataBatchRecipeConstraint struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type TERCCV2DataBatchRecipeTestCase struct {
	// Mandatory fields
	ID string `json:"id"`

	// Optional fields
	Tracker string `json:"tracker,omitempty"`
	URI     string `json:"uri,omitempty"`
}

type TERCCV2DataSelectionStrategy struct {
	// Mandatory fields
	ID string `json:"id"`

	// Optional fields
	Tracker string `json:"tracker,omitempty"`
	URI     string `json:"uri,omitempty"`
}
