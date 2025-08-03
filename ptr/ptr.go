package ptr

import (
	"errors"
	"fmt"
	"reflect"
)

// EnforcePtr enforces that a value must be a non-nil pointer and returns the value that the pointer references.
func EnforcePtr(obj any) (reflect.Value, error) {
	v := reflect.ValueOf(obj)

	if v.Kind() != reflect.Ptr {
		if !v.IsValid() {
			return reflect.Value{}, errors.New("expected pointer, but got invalid")
		}

		return reflect.Value{}, fmt.Errorf("expected a pointer, but got %v", v.Type())
	}

	if v.IsNil() {
		return reflect.Value{}, errors.New("obj is a nil pointer")
	}

	return v.Elem(), nil
}
