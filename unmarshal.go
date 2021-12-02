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
	"fmt"
	"reflect"
	"sort"

	"github.com/Masterminds/semver"
	"github.com/tidwall/gjson"
)

var ErrMalformedInput error = errors.New("malformed JSON input")
var ErrUnsupportedEvent error = errors.New("event unsupported")

// UnmarshalAny unmarshals a JSON string into one of the supported Eiffel
// event type structs, dependent on the meta.type field in the event
// payload. Callers are expected to use type assertions or type switches
// to access concrete event types.
//
// If the input isn't valid JSON, meta.type and meta.version values can't
// be extracted from it, or some other JSON unmarshaling error occurs an
// ErrMalformedInput error is returned. If the event type or version isn't
// supported by this implementation an ErrUnsupportedType error is returned.
func UnmarshalAny(input []byte) (interface{}, error) {
	if !gjson.ValidBytes(input) {
		return nil, fmt.Errorf("%w: not valid JSON", ErrMalformedInput)
	}

	// Extract the event type and version from the JSON payload.
	schemaSelectors := gjson.GetManyBytes(input, "meta.type", "meta.version")
	metaType := schemaSelectors[0].String()
	if metaType == "" {
		return nil, fmt.Errorf("%w: unable to extract meta.type", ErrMalformedInput)
	}
	metaVersion := schemaSelectors[1].String()
	if metaVersion == "" {
		return nil, fmt.Errorf("%w: unable to extract meta.version", ErrMalformedInput)
	}
	version, err := semver.NewVersion(metaVersion)
	if err != nil {
		return nil, fmt.Errorf("%w: unable to parse meta.version: %s", ErrMalformedInput, err)
	}

	// Verify that the event type and version combo is supported.
	if _, ok := eventTypeTable[metaType]; !ok {
		return nil, fmt.Errorf("%w: type: %s", ErrUnsupportedEvent, metaType)
	}
	if _, ok := eventTypeTable[metaType][version.Major()]; !ok {
		var versions []int
		for k := range eventTypeTable[metaType] {
			versions = append(versions, int(k))
		}
		sort.Ints(versions)
		return nil, fmt.Errorf("%w: version of %s unsupported; valid major versions: %v", ErrUnsupportedEvent, metaType, versions)
	}

	// Create an instance of the right struct and unmarshal the payload into it.
	value := reflect.New(eventTypeTable[metaType][version.Major()].structType).Interface()
	if err := json.Unmarshal(input, &value); err != nil {
		// Ideally we should wrap both ErrMalformedInput and err
		// but I couldn't figure out an elegant way of doing that.
		return nil, fmt.Errorf("%w: %s", ErrMalformedInput, err)
	}
	return value, nil
}
