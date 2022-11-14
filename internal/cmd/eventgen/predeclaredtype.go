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

package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/eiffel-community/eiffelevents-sdk-go"
)

// goPredeclaredType represents a versioned type declared by this SDK
// but not in the current file. This is used with schema properties that
// contain a JSON reference to an external schema file.
type goPredeclaredType struct {
	BaseName string // Unversioned name of the type.
	TypeName string // Versioned name of the type.
	Version  *semver.Version
}

func newPredeclaredType(schemaRef string) (*goPredeclaredType, error) {
	name := filepath.Base(filepath.Dir(schemaRef))
	version, err := semver.NewVersion(strings.TrimSuffix(filepath.Base(schemaRef), filepath.Ext(schemaRef)))
	if err != nil {
		return nil, fmt.Errorf("error parsing version number from schema reference %q: %w", schemaRef, err)
	}
	return &goPredeclaredType{
		BaseName: eiffelevents.EventStructName(name, version),
		TypeName: eiffelevents.VersionedEventStructName(name, version),
		Version:  version,
	}, nil
}

func (t *goPredeclaredType) String() string {
	return t.TypeName
}
