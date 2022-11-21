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
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
)

// trimmedTypeNameSuffixes lists the schema name suffixes that'll get trimmed
// when turning the schema into a Go type.
var trimmedTypeNameSuffixes = []string{
	"Event",
	"Property",
}

// StructName returns the non-versioned name of the Go struct used to
// represent a type.
func StructName(eventType string, eventVersion *semver.Version) string {
	s := strings.TrimPrefix(eventType, "Eiffel")
	for _, suffix := range trimmedTypeNameSuffixes {
		// We want to break after the first removed suffix so we can't just
		// call strings.TrimSuffix.
		if strings.HasSuffix(s, suffix) {
			return strings.TrimSuffix(s, suffix)
		}
	}
	return s
}

// VersionedStructName returns the name of the Go struct used to represent
// a particular version of a type.
func VersionedStructName(eventType string, eventVersion *semver.Version) string {
	return fmt.Sprintf("%sV%d", StructName(eventType, eventVersion), eventVersion.Major())
}
