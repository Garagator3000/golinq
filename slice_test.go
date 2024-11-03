package golinq

import (
	"reflect"
	"strings"
	"testing"
)

func TestSlice_NewSlice(t *testing.T) {
	tcInt := []int{1, 2, 3}
	sliceInt := NewSlice(tcInt)
	if !reflect.DeepEqual(sliceInt.items, tcInt) {
		t.Errorf("expected %v, got %v", tcInt, sliceInt.items)
	}

	tcStr := []string{"a", "b", "c"}
	sliceString := NewSlice(tcStr)
	if !reflect.DeepEqual(sliceString.items, tcStr) {
		t.Errorf("expected %v, got %v", tcStr, sliceString.items)
	}

	tcBool := []bool{true, false, true}
	sliceBool := NewSlice(tcBool)
	if !reflect.DeepEqual(sliceBool.items, tcBool) {
		t.Errorf("expected %v, got %v", tcBool, sliceBool.items)
	}

	tcEmpty := []int{}
	sliceEmpty := NewSlice(tcEmpty)
	if !reflect.DeepEqual(sliceEmpty.items, tcEmpty) {
		t.Errorf("expected %v, got %v", tcEmpty, sliceEmpty.items)
	}
}

func TestSlice_Len(t *testing.T) {
	tcs := []struct {
		items  []int
		result int
	}{
		{
			items:  []int{1, 2, 3},
			result: 3,
		},
		{
			items:  []int{},
			result: 0,
		},
	}

	for _, tc := range tcs {
		s := NewSlice(tc.items)
		if s.Len() != tc.result {
			t.Errorf("expected %d, got %d", tc.result, s.Len())
		}
	}
}

func TestSlice_Items(t *testing.T) {
	tcs := []struct {
		items  []int
		result []int
	}{
		{
			items:  []int{1, 2, 3},
			result: []int{1, 2, 3},
		},
		{
			items:  []int{},
			result: []int{},
		},
	}

	for _, tc := range tcs {
		s := NewSlice(tc.items)
		if !reflect.DeepEqual(s.Items(), tc.result) {
			t.Errorf("expected %v, got %v", tc.result, s.Items())
		}
	}
}

func TestSlice_GetItemByIndex(t *testing.T) {
	tcs := []struct {
		items  []int
		index  int
		result int
	}{
		{
			items:  []int{1, 2, 3},
			index:  0,
			result: 1,
		},
		{
			items:  []int{1, 2, 3},
			index:  1,
			result: 2,
		},
		{
			items:  []int{1, 2, 3},
			index:  10,
			result: 0,
		},
	}

	for _, tc := range tcs {
		s := NewSlice(tc.items)
		if s.GetItemByIndex(tc.index) != tc.result {
			t.Errorf("expected %d, got %d", tc.result, s.GetItemByIndex(tc.index))
		}
	}
}

func TestSlice_GetFirstIndex(t *testing.T) {
	tcs := []struct {
		items  []int
		item   int
		result int
	}{
		{
			items:  []int{1, 2, 1},
			item:   1,
			result: 0,
		},
		{
			items:  []int{1, 2, 1},
			item:   3,
			result: -1,
		},
	}

	for _, tc := range tcs {
		s := NewSlice(tc.items)
		if s.GetFirstIndex(tc.item) != tc.result {
			t.Errorf("expected %d, got %d", tc.result, s.GetFirstIndex(tc.item))
		}
	}
}

func TestSlice_GetLastIndex(t *testing.T) {
	tcs := []struct {
		items  []int
		item   int
		result int
	}{
		{
			items:  []int{1, 2, 1},
			item:   1,
			result: 2,
		},
		{
			items:  []int{1, 2, 1},
			item:   3,
			result: -1,
		},
	}

	for _, tc := range tcs {
		s := NewSlice(tc.items)
		if s.GetLastIndex(tc.item) != tc.result {
			t.Errorf("expected %d, got %d", tc.result, s.GetLastIndex(tc.item))
		}
	}
}

func TestSlice_GetAllIndexes(t *testing.T) {
	tcs := []struct {
		items  []int
		item   int
		result []int
	}{
		{
			items:  []int{1, 2, 1},
			item:   1,
			result: []int{0, 2},
		},
	}

	for _, tc := range tcs {
		s := NewSlice(tc.items)
		if !reflect.DeepEqual(s.GetAllIndexes(tc.item), tc.result) {
			t.Errorf("expected %v, got %v", tc.result, s.GetAllIndexes(tc.item))
		}
	}
}

func TestSlice_Add(t *testing.T) {
	tcs := []struct {
		items  []int
		add    int
		result []int
	}{
		{
			items:  []int{1, 2, 3},
			add:    4,
			result: []int{1, 2, 3, 4},
		},
		{
			items:  []int{1, 2, 3},
			add:    0,
			result: []int{1, 2, 3, 0},
		},
	}

	for _, tc := range tcs {
		s := NewSlice(tc.items)

		s = s.Add(tc.add)
		if len(s.items) != len(tc.result) {
			t.Errorf("expected %v, got %v", tc.result, s.items)
		}

		for i, v := range s.Items() {
			if v != tc.result[i] {
				t.Errorf("expected %v, got %v", tc.result, s.items)
			}
		}
	}
}

func TestSlice_RemoveFirst(t *testing.T) {
	tcs := []struct {
		items  []int
		remove int
		result []int
	}{
		{
			items:  []int{1, 2, 3},
			remove: 1,
			result: []int{2, 3},
		},
		{
			items:  []int{1, 2, 3},
			remove: 0,
			result: []int{1, 2, 3},
		},
		{
			items:  []int{1, 2, 3},
			remove: 3,
			result: []int{1, 2},
		},
	}

	for _, tc := range tcs {
		s := NewSlice(tc.items)

		s = s.RemoveFirst(tc.remove)
		if len(s.items) != len(tc.result) {
			t.Errorf("expected %v, got %v", tc.result, s.items)
		}

		for i, v := range s.Items() {
			if v != tc.result[i] {
				t.Errorf("expected %v, got %v", tc.result, s.items)
			}
		}
	}
}

func TestSlice_Remove(t *testing.T) {
	tcs := []struct {
		items  []int
		remove int
		result []int
	}{
		{
			items:  []int{1, 2, 3},
			remove: 1,
			result: []int{2, 3},
		},
		{
			items:  []int{1, 2, 3},
			remove: 0,
			result: []int{1, 2, 3},
		},
		{
			items:  []int{1, 3, 2, 3},
			remove: 3,
			result: []int{1, 2},
		},
		{
			items:  []int{1, 1, 1, 1, 1, 1, 1},
			remove: 1,
			result: []int{},
		},
	}

	for _, tc := range tcs {
		s := NewSlice(tc.items)

		s = s.Remove(tc.remove)
		if len(s.items) != len(tc.result) {
			t.Errorf("expected %v, got %v", tc.result, s.items)
		}

		for i, v := range s.Items() {
			if v != tc.result[i] {
				t.Errorf("expected %v, got %v", tc.result, s.items)
			}
		}
	}
}

func TestSlice_Filter(t *testing.T) {
	tcs := []struct {
		items  []int
		filter func(int) bool
		result []int
	}{
		{
			items:  []int{1, 2, 3},
			filter: func(i int) bool { return i%2 == 0 },
			result: []int{2},
		},
		{
			items:  []int{1, 2, 3},
			filter: func(i int) bool { return true },
			result: []int{1, 2, 3},
		},
		{
			items:  []int{1, 2, 3},
			filter: func(i int) bool { return false },
			result: []int{},
		},
		{
			items:  []int{1, 2, 3},
			filter: func(i int) bool { return i%2 == 1 },
			result: []int{1, 3},
		},
		{
			items:  []int{1, 2, 3},
			filter: func(i int) bool { return i == 1 },
			result: []int{1},
		},
		{
			items:  []int{1, 2, 3},
			filter: func(i int) bool { return i == 4 },
			result: []int{},
		},
		{
			items:  []int{1, 2, 3},
			filter: func(i int) bool { return i > 1 },
			result: []int{2, 3},
		},
	}

	for _, tc := range tcs {
		s := NewSlice(tc.items)

		s = s.Filter(tc.filter)
		if len(s.items) != len(tc.result) {
			t.Errorf("expected %v, got %v", tc.result, s.items)
		}

		for i, v := range s.Items() {
			if v != tc.result[i] {
				t.Errorf("expected %v, got %v", tc.result, s.items)
			}
		}
	}
}

func TestSlice_Contains(t *testing.T) {
	tcs := []struct {
		items  []int
		item   int
		result bool
	}{
		{
			items:  []int{1, 2, 3},
			item:   2,
			result: true,
		},
		{
			items:  []int{1, 2, 3},
			item:   4,
			result: false,
		},
	}

	for _, tc := range tcs {
		s := NewSlice(tc.items)
		if s.Contains(tc.item) != tc.result {
			t.Errorf("expected %v, got %v", tc.result, s.Contains(tc.item))
		}
	}
}

func TestSlice_contains(t *testing.T) {
	tcs := []struct {
		items  []int
		item   int
		result bool
	}{
		{
			items:  []int{1, 2, 3},
			item:   2,
			result: true,
		},
		{
			items:  []int{1, 2, 3},
			item:   4,
			result: false,
		},
	}

	for _, tc := range tcs {
		if contains(tc.items, tc.item) != tc.result {
			t.Errorf("expected %v, got %v", tc.result, contains(tc.items, tc.item))
		}
	}
}

func TestIntegration_Int(t *testing.T) {
	slice := NewSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	expected := []int{3, 9, 12, 13}

	filter := func(i int) bool {
		return i%3 == 0
	}

	newSlice := slice.
		Add(11).        // 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
		Add(12).        // 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12
		RemoveFirst(2). // 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12
		Filter(filter). // 3, 6, 9, 12
		Add(13).        // 3, 6, 9, 12, 13
		Add(6).         // 3, 6, 9, 12, 13, 6
		Add(6).         // 3, 6, 9, 12, 13, 6, 6
		Remove(6)       // 3, 9, 12, 13

	if len(newSlice.items) != len(expected) {
		t.Errorf("expected %d items, got %v", len(expected), newSlice.items)
	}

	if !reflect.DeepEqual(expected, newSlice.items) {
		t.Errorf("expected %v, got %v", expected, newSlice.items)
	}

}

func TestIntegration_String(t *testing.T) {
	slice := NewSlice([]string{"abc", "zxc", "qwe"})
	expected := []string{"qwe", "def"}

	filter := func(s string) bool {
		return strings.Contains(s, "e")
	}

	newSlice := slice.Add("def"). // abc, zxc, qwe, def
					Add("rty").    // abc, zxc, qwe, def, rty
					Remove("zxc"). // abc, qwe, def, rty
					Filter(filter) // qwe, def

	if len(newSlice.items) != len(expected) {
		t.Errorf("expected %d items, got %v", len(expected), newSlice.items)
	}

	if !reflect.DeepEqual(expected, newSlice.items) {
		t.Errorf("expected %v, got %v", expected, newSlice.items)
	}
}