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

package eiffelevents

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

//go:generate go run ./internal/cmd/roundtriptestgen protocol/examples/events events_test_roundtripdata.go

// TestRoundtrip reads example events from the Eiffel protocol repository,
// deserializes them into Go structs, verifies that we get the expected struct,
// and finally serializes back to JSON and verifies that the result is
// equivalent to what we started with.
func TestRoundtrip(t *testing.T) {
	for _, tc := range eventRoundTripTestTable {
		t.Run(tc.Filename, func(t *testing.T) {
			f, err := os.Open(tc.Filename)
			if err != nil {
				t.Fatal(err.Error())
			}
			defer f.Close()
			input, err := io.ReadAll(f)
			if err != nil {
				t.Fatal(err.Error())
			}

			event, err := UnmarshalAny(input)
			assert.NoError(t, err)
			assert.IsType(t, tc.ExpectedType, event)

			output, err := json.Marshal(event)
			assert.NoError(t, err)
			assert.JSONEq(t, string(input), string(output))
		})
	}
}

// TestMajorVersionFactory calls the factory for a randomly chosen major
// version of an event type and checks that we get reasonable results.
// All factory functions are generated so we assume that if one is okay
// then so are the rest.
func TestMajorVersionFactory(t *testing.T) {
	// We're deliberately avoiding the most recent major version to avoid having
	// to update this test if there's another minor version of the event.
	event, err := NewActivityTriggeredV3()
	assert.NoError(t, err)

	assert.Equal(t, "EiffelActivityTriggeredEvent", event.Meta.Type)
	_, err = uuid.Parse(event.Meta.ID)
	assert.NoError(t, err)
	assert.Equal(t, "3.0.0", event.Meta.Version)

	// Sanity check that meta.time is within two minutes of the current time.
	eventTime := time.UnixMilli(0).Add(time.Duration(event.Meta.Time) * time.Millisecond)
	assert.WithinDuration(t, time.Now(), eventTime, 2*time.Minute)
}

// TestFactoriesWithModifiers tests that factories interact nicely with
// the optional list of Modifier functions. The functions themselves have
// their own tests.
func TestFactoriesWithModifiers(t *testing.T) {
	// Single modifier
	event, err := NewActivityTriggeredV3(WithSourceName("name"))
	assert.NoError(t, err)
	assert.Equal(t, "name", event.Meta.Source.Name)

	// Modifiers are applied in order
	event, err = NewActivityTriggeredV3(WithSourceName("name1"), WithSourceName("name2"))
	assert.NoError(t, err)
	assert.Equal(t, "name2", event.Meta.Source.Name)

	// Different kinds of modifiers
	event, err = NewActivityTriggeredV3(WithSourceHost("host"), WithSourceName("name"))
	assert.NoError(t, err)
	assert.Equal(t, "host", event.Meta.Source.Host)
	assert.Equal(t, "name", event.Meta.Source.Name)

	// Modifier errors are propagated
	modifierError := errors.New("the modifier failed")
	_, err = NewActivityTriggeredV3(func(fieldSetter FieldSetter) error { return modifierError })
	assert.ErrorIs(t, err, modifierError)
}

// TestLinksFindAll tests the functionaltiy of the generated FindAll method
// for one of the XXXLinks types.
func TestLinksFindAll(t *testing.T) {
	testcases := []struct {
		name     string
		links    ActTV3Links
		wanted   string
		expected []string
	}{
		{
			"Empty list of links",
			ActTV3Links{},
			"CAUSE",
			[]string{},
		},
		{
			"Single matching link",
			ActTV3Links{
				ActTV3Link{"context id", "CONTEXT"},
				ActTV3Link{"cause id", "CAUSE"},
			},
			"CAUSE",
			[]string{"cause id"},
		},
		{
			"No matching links",
			ActTV3Links{
				ActTV3Link{"context id", "CONTEXT"},
			},
			"CAUSE",
			[]string{},
		},
		{
			"Returns all matching links",
			ActTV3Links{
				ActTV3Link{"cause id 1", "CAUSE"},
				ActTV3Link{"context id", "CONTEXT"},
				ActTV3Link{"cause id 2", "CAUSE"},
			},
			"CAUSE",
			[]string{"cause id 1", "cause id 2"},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.links.FindAll(tc.wanted))
		})
	}
}

// TestLinksFindFirst tests the functionaltiy of the generated FindFirst method
// for one of the XXXLinks types.
func TestLinksFindFirst(t *testing.T) {
	testcases := []struct {
		name     string
		links    ActTV3Links
		wanted   string
		expected string
	}{
		{
			"Empty list of links",
			ActTV3Links{},
			"CAUSE",
			"",
		},
		{
			"Single matching link",
			ActTV3Links{
				ActTV3Link{"context id", "CONTEXT"},
				ActTV3Link{"cause id", "CAUSE"},
			},
			"CAUSE",
			"cause id",
		},
		{
			"No matching links",
			ActTV3Links{
				ActTV3Link{"context id", "CONTEXT"},
			},
			"CAUSE",
			"",
		},
		{
			"Returns first of multiple matching links",
			ActTV3Links{
				ActTV3Link{"cause id 1", "CAUSE"},
				ActTV3Link{"context id", "CONTEXT"},
				ActTV3Link{"cause id 2", "CAUSE"},
			},
			"CAUSE",
			"cause id 1",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.links.FindFirst(tc.wanted))
		})
	}
}
