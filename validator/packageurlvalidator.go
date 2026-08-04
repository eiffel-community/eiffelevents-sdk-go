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

package validator

import (
	"context"
	"fmt"

	"github.com/package-url/packageurl-go"
	"github.com/tidwall/gjson"
)

// PackageURLValidator validates Package URL fields in selected event types.
type PackageURLValidator struct{}

func NewPackageURLValidator() *PackageURLValidator {
	return &PackageURLValidator{}
}

// Validate validates the data.identity field as a package URL for
// EiffelArtifactCreatedEvent events.
func (pv *PackageURLValidator) Validate(_ context.Context, event []byte) error {
	fields := gjson.GetManyBytes(event, "meta.type", "data.identity")

	eventType := ""
	if fields[0].Type == gjson.String {
		eventType = fields[0].String()
	}
	if eventType != "EiffelArtifactCreatedEvent" {
		return nil
	}

	if fields[1].Type != gjson.String || fields[1].String() == "" {
		return fmt.Errorf("missing or invalid contents of data.identity field for %s", eventType)
	}

	if _, err := packageurl.FromString(fields[1].String()); err != nil {
		return fmt.Errorf("invalid package URL in data.identity field for %s: %w", eventType, err)
	}

	return nil
}
