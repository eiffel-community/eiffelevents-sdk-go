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
	_ "embed"
	"fmt"

	"github.com/eiffel-community/eiffelevents-sdk-go"
	jsschema "github.com/lestrrat-go/jsschema"
)

// goLinkSlice represents the type of a slice of links. It's different from
// normal slices in that it gets its own type, e.g. EventLinksV1 defined as
// []EventLinkV1, which allows us to define new methods on the slice itself.
type goLinkSlice struct {
	TypeName   string
	SlicedType goType // The type that's being sliced
}

func newLinkSlice(parent *goStruct, name string, items *jsschema.ItemSpec) (*goLinkSlice, error) {
	// We expect to be handed an array where the elements are of a predeclared
	// type. Extract that type so that we can turn it into a slice type.
	typ, err := goTypeFromSchema(parent, pluralizeClient.Singular(name), items.Schemas[0])
	if err != nil {
		return nil, err
	}
	predeclType, ok := typ.(*goPredeclaredType)
	if !ok {
		return nil, fmt.Errorf("type %T of link array items was unexpected", predeclType)
	}
	return &goLinkSlice{
		TypeName:   eiffelevents.VersionedEventStructName(pluralizeClient.Plural(predeclType.BaseName), predeclType.Version),
		SlicedType: typ,
	}, nil
}

func (t *goLinkSlice) String() string {
	return t.TypeName
}
