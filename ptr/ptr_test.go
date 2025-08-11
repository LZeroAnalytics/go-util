package ptr_test

import (
	"reflect"
	"testing"

	"github.com/lzeroanalytics/go-util/ptr"
)

func TestEnforcePtr(t *testing.T) {
	cases := []struct {
		name     string
		in       any
		expected reflect.Kind
		ok       bool
	}{
		{
			name:     "invalid interface value",
			in:       nil,
			expected: reflect.Invalid,
			ok:       false,
		},
		{
			name:     "non-pointer value",
			in:       1,
			expected: reflect.Invalid,
			ok:       false,
		},
		{
			name:     "nil pointer",
			in:       (*int)(nil),
			expected: reflect.Invalid,
			ok:       false,
		},
		{
			name: "underlying int pointer",
			in: func() *int {
				v := 1
				return &v
			}(),
			expected: reflect.Int,
			ok:       true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			v, err := ptr.EnforcePtr(tc.in)

			if tc.ok {
				if err != nil {
					t.Fatalf("EnforcePtr(%v) unexpected error: %v", tc.in, err)
				}

				if got := v.Kind(); got != tc.expected {
					t.Fatalf("EnforcePtr(%v) kind = %v, expected %v", tc.in, got, tc.expected)
				}

				return
			}

			if err == nil {
				t.Fatalf("EnforcePtr(%v) expected error", tc.in)
			}
		})
	}
}

func TestTo(t *testing.T) {
	v1 := 1
	r1 := ptr.To(v1)
	if *r1 != v1 {
		t.Fatalf("To(%v) returned %v", v1, *r1)
	}

	v2 := "test"
	r2 := ptr.To(v2)
	if *r2 != v2 {
		t.Fatalf("To($%v) returned %v", v2, *r2)
	}
}
