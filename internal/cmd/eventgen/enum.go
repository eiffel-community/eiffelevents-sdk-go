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
	"regexp"
	"strings"

	"github.com/magnusbaeck/eiffelevents-sdk-go/internal/codetemplate"
)

// goEnum represents a Go type that acts as an enumeration. This type has the following form:
//
//     type EnumName UnderlyingType
//
//     const (
//	       EnumName_Value1 EnumName = "VALUE1"
//	       EnumName_Value2 EnumName = "VALUE2"
//	   )
//
// It currently assumes that the underlying type is a string and that that
// string (with limited mangling) can be appended to the name of the enum type
// to form a valid constant name.
type goEnum struct {
	Name   string
	Type   goType
	Values []goEnumValue
}

func newEnum(parent *goStruct, name string, typ goType, values []interface{}) (*goEnum, error) {
	enumTypeName := parent.SubTypeNamePrefix + strings.Title(name)

	var enumValues []goEnumValue
	for _, value := range values {
		enumValue, err := newEnumValue(enumTypeName, value)
		if err != nil {
			return nil, err
		}
		enumValues = append(enumValues, enumValue)
	}
	return &goEnum{
		Name:   enumTypeName,
		Type:   typ,
		Values: enumValues,
	}, nil
}

//go:embed templates/enum_decl.tmpl
var enumDeclTemplate string

func (t *goEnum) declare(ct *codetemplate.OutputFile) error {
	return ct.ExpandTemplate(enumDeclTemplate, t)
}

func (t *goEnum) String() string {
	return t.Name
}

// Regexp used to check if an enum value looks like an abbreviation,
// in which case a different capitalization algorithm applies.
// Abbreviations are notably used for names of crypographic algorithms.
var isConstAbbrevExpr = regexp.MustCompile(`^[A-Z]+\d+$`)

// goEnumValue represents a single valid value for an enum.
type goEnumValue struct {
	ConstName string
	TypeName  string
	Value     string
}

func newEnumValue(typeName string, value interface{}) (goEnumValue, error) {
	strValue, ok := value.(string)
	if !ok {
		return goEnumValue{}, fmt.Errorf("enum value for type %s not a string type: %#v", typeName, value)
	}

	// Use the value as-is if it's an abbreviation. Otherwise use a series of
	// string transforms to turn all-caps snake case strings (ONE_TWO) into
	// Go-style strings (OneTwo).
	constantSuffix := strValue
	if !isConstAbbrevExpr.MatchString(strValue) {
		constantSuffix = strings.Replace(strings.Title(strings.ToLower(strings.Replace(strValue, "_", " ", -1))), " ", "", -1)
	}

	return goEnumValue{
		ConstName: fmt.Sprintf("%s_%s", typeName, constantSuffix),
		TypeName:  typeName,
		Value:     strValue,
	}, nil
}
