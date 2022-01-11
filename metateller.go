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

// MetaTeller allow introspection of key metadata fields of anything that
// resembles an Eiffel event. Not all meta fields are included since they
// may vary over time, but is believed to be a time-invariant subset.
type MetaTeller interface {
	// ID returns the value of the meta.id field.
	ID() string

	// Type returns the value of the meta.type field.
	Type() string

	// Version returns the value of the meta.version field.
	Version() string

	// Time returns the value of the meta.time field.
	Time() int64

	// DomainID returns the value of the meta.source.domainId field.
	DomainID() string
}
