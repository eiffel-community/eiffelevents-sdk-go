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

// CapabilityTeller answers questions about the capabilities of anything that
// resembles an Eiffel event without having to manually track which version of
// which event type introduced this capability.
type CapabilityTeller interface {
	// SupportsSigning returns true if the event supports signatures according
	// to V3 of the meta field, i.e. events where the signature is found under
	// meta.security.integrityProtection.
	SupportsSigning() bool
}
