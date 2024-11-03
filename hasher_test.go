package golinq

import (
	"testing"
)

type cyclicStruct struct {
	A *cyclicStruct
}

func TestHasher_Hash(t *testing.T) {
	tcs := []struct {
		name   string
		obj1   interface{}
		obj2   interface{}
		equals bool
	}{
		{
			name:   "Equals strings",
			obj1:   "Alex",
			obj2:   "Alex",
			equals: true,
		},
		{
			name:   "Not equals strings",
			obj1:   "Alex",
			obj2:   "Alex1",
			equals: false,
		},
		{
			name:   "Equals ints",
			obj1:   1,
			obj2:   1,
			equals: true,
		},
		{
			name:   "Not equals ints",
			obj1:   1,
			obj2:   2,
			equals: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			hash1 := Hash(tc.obj1)
			hash2 := Hash(tc.obj2)
			if hash1 == hash2 != tc.equals {
				t.Errorf("expected %v, got %v", tc.equals, hash1 == hash2)
			}
		})
	}
}

func TestHasher_HashWithPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	badObj := &cyclicStruct{}
	badObj.A = badObj

	_ = Hash(badObj)
}
