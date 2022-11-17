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
	"text/template"

	jsschema "github.com/lestrrat-go/jsschema"

	"github.com/eiffel-community/eiffelevents-sdk-go/internal/codetemplate"
)

// goLinkSlice represents the type of a slice value.
type goLinkSlice struct {
	TypeName   string
	SlicedType goType // The type that's being sliced
}

func newLinkSlice(parent *goStruct, name string, items *jsschema.ItemSpec) (*goLinkSlice, error) {
	// The type created for this array should use the singular form of the noun, i.e. we want
	//     Links XXXLink
	// and not this:
	//     Links XXXLinks
	typ, err := goTypeFromSchema(parent, pluralizeClient.Singular(name), items.Schemas[0])
	if err != nil {
		return nil, err
	}
	return &goLinkSlice{
		TypeName:   parent.SubTypeNamePrefix + fieldTitle(name),
		SlicedType: typ,
	}, nil
}

//go:embed templates/linkslice_decl.tmpl
var linkSliceDeclTemplate string

func (t *goLinkSlice) declare(ct *codetemplate.OutputFile) error {
	// Declare the custom slice type itself.
	if err := ct.ExpandTemplate(linkSliceDeclTemplate, t, template.FuncMap{}); err != nil {
		return err
	}

	if declarer, ok := t.SlicedType.(goTypeDeclarer); ok {
		return declarer.declare(ct)
	}
	return nil
}

func (t *goLinkSlice) String() string {
	return t.TypeName
}
