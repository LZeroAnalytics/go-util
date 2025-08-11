package ptr

import (
	"errors"
	"fmt"
	"reflect"
)

type enforceOptions struct {
	allowNil bool
}

type EnforceOpt func(*enforceOptions)

func AllowNil() EnforceOpt {
	return func(o *enforceOptions) {
		o.allowNil = true
	}
}

// EnforcePtr enforces that a value must be a non-nil pointer and returns the value that the pointer references.
func EnforcePtr(obj any, opts ...EnforceOpt) (reflect.Value, error) {
	base := &enforceOptions{}

	for _, opt := range opts {
		opt(base)
	}

	v := reflect.ValueOf(obj)

	if v.Kind() != reflect.Ptr {
		if !v.IsValid() {
			return reflect.Value{}, errors.New("expected pointer, but got invalid")
		}

		return reflect.Value{}, fmt.Errorf("expected a pointer, but got %v", v.Type())
	}

	if v.IsNil() && !base.allowNil {
		return reflect.Value{}, errors.New("obj is a nil pointer")
	}

	return v.Elem(), nil
}

// To returns a pointer to the given value.
func To[T any](v T) *T {
	return &v
}
