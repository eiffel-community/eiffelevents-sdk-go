package eiffelevents

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// FieldSetter sets struct field by name using the usual toplevelfield.subfield notation.
// Fields should be expressed with their JSON names, i.e. based on the "json" tag.
type FieldSetter interface {
	SetField(fieldName string, value interface{}) error
}

// setField sets the value of a struct field whose path is expressed
// with dot notation using the field names in the "json" tag. The reflect.Value
// passed as the target must be a pointer to a struct.
func setField(target reflect.Value, fieldName string, value interface{}) error {
	if target.Kind() != reflect.Ptr {
		return errors.New("target value is not a pointer")
	}
	elemVal := target.Elem()
	if elemVal.Kind() != reflect.Struct {
		return fmt.Errorf("the target value must point to a struct but pointed to a %s", elemVal.Kind())
	}

	// Recurse if there's a dot in the field name, otherwise try to set that field.
	dotIndex := strings.Index(fieldName, ".")
	if dotIndex == 0 {
		return fmt.Errorf("invalid field name: %q", fieldName)
	}
	if dotIndex > 0 {
		firstLevelFieldName := fieldName[0:dotIndex]
		remainderFieldName := fieldName[dotIndex+1:]

		// Locate the struct to recurse into
		field, err := getJSONField(elemVal, firstLevelFieldName)
		if err != nil {
			return err
		}
		if !field.CanAddr() {
			return fmt.Errorf("struct field %q is not addressable", firstLevelFieldName)
		}
		return setField(field.Addr(), remainderFieldName, value)
	} else {
		// Locate the field to set
		field, err := getJSONField(elemVal, fieldName)
		if err != nil {
			return err
		}
		if !field.CanSet() {
			return fmt.Errorf("struct field %q cannot be set", fieldName)
		}
		field.Set(reflect.ValueOf(value))
	}

	return nil
}

// getJSONField returns the value of a struct field that has the given JSON name in the "json" tag.
func getJSONField(structVal reflect.Value, fieldName string) (reflect.Value, error) {
	structType := structVal.Type()
	for i := 0; i < structType.NumField(); i++ {
		jsonField := structType.Field(i).Tag.Get("json")
		if jsonField == "" || jsonField == "-" {
			continue
		}

		// Trim any additional options from the end,
		// i.e. turn "fieldname,omitempty" into plain "fieldname".
		if commaIndex := strings.Index(jsonField, ","); commaIndex != -1 {
			jsonField = jsonField[0:commaIndex]
		}

		if jsonField == fieldName {
			return structVal.Field(i), nil
		}
	}
	return reflect.Value{}, fmt.Errorf("the struct did not contain a field with the JSON name %q", fieldName)
}
