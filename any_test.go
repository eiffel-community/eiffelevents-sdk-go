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
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleAny() {
	inputData := `
	[
	  {
	    "meta": {
	      "type": "EiffelCompositionDefinedEvent",
	      "version": "3.2.0",
	      "time": 1234567890,
	      "id": "aaaaaaaa-bbbb-5ccc-8ddd-eeeeeeeeeee0"
	    },
	    "data": {
	      "name": "My Composition"
	    },
	    "links": []
	  }
	]`
	var eventSlice []*Any
	_ = json.Unmarshal([]byte(inputData), &eventSlice)

	// Any implements MetaTeller so you can call its methods directly
	// if you e.g. only need the event's ID.
	fmt.Printf("The array contains an event with ID %s.\n", eventSlice[0].ID())

	// Or use a type switch if you want access to the full event.
	switch event := eventSlice[0].Get().(type) {
	case *CompositionDefinedV3:
		fmt.Printf("It's a CompositionDefined event with the name %q.\n", event.Data.Name)
	default:
		fmt.Println("Unsupported event type.")
	}

	// Output: The array contains an event with ID aaaaaaaa-bbbb-5ccc-8ddd-eeeeeeeeeee0.
	// It's a CompositionDefined event with the name "My Composition".
}

//go:embed testdata/any/object_with_event_array.json
var eventArrayObjectTestJSON []byte

//go:embed testdata/any/single_event.json
var singleEventTestJSON []byte

func TestAnyUnmarshalArray(t *testing.T) {
	input, err := os.Open("testdata/any/object_with_event_array.json")
	require.NoError(t, err)
	defer input.Close()

	var data struct {
		Events []*Any `json:"events"`
	}
	require.NoError(t, json.Unmarshal(eventArrayObjectTestJSON, &data))

	assert.Equal(t, "aaaaaaaa-bbbb-5ccc-8ddd-eeeeeeeeeee0", data.Events[0].ID())
	assert.Equal(t, "EiffelCompositionDefinedEvent", data.Events[0].Type())
	assert.Equal(t, "aaaaaaaa-bbbb-5ccc-8ddd-eeeeeeeeeee1", data.Events[1].ID())
	assert.Equal(t, "EiffelArtifactCreatedEvent", data.Events[1].Type())
}

func TestAnyUnmarshalSingleMetaTeller(t *testing.T) {
	var event Any
	require.NoError(t, json.Unmarshal(singleEventTestJSON, &event))

	assert.Equal(t, "aaaaaaaa-bbbb-5ccc-8ddd-eeeeeeeeeee0", event.ID())
	assert.Equal(t, "EiffelCompositionDefinedEvent", event.Type())
}

func TestAnyUnmarshalSingleTypeAssertion(t *testing.T) {
	var event Any
	require.NoError(t, json.Unmarshal(singleEventTestJSON, &event))

	v, ok := event.Get().(*CompositionDefinedV3)
	require.True(t, ok)
	assert.Equal(t, "aaaaaaaa-bbbb-5ccc-8ddd-eeeeeeeeeee0", v.ID())
	assert.Equal(t, "EiffelCompositionDefinedEvent", v.Type())
}

func TestAnyUnmarshalMarshalSymmetric(t *testing.T) {
	var event Any
	require.NoError(t, json.Unmarshal(singleEventTestJSON, &event))

	var buf bytes.Buffer
	require.NoError(t, json.NewEncoder(&buf).Encode(event))

	assert.JSONEq(t, string(singleEventTestJSON), buf.String())
}
