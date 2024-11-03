package golinq

import (
	"testing"
)

func TestCollection_NewCollection(t *testing.T) {
	collection := NewCollection[string]()

	if collection == nil {
		t.Errorf("expected not nil, got %v", collection)
	}
}

func TestCollection_NewCollectionWithItems(t *testing.T) {
	tcs := []struct {
		name   string
		items  []string
		result map[string]string
	}{
		{
			name:  "a,b,c to map",
			items: []string{"a", "b", "c"},
			result: map[string]string{ // Also check, what hash function returns idempotent values.
				"ac8d8342bbb2362d13f0a559a3621bb407011368895164b628a54f7fc33fc43c": "a",
				"c100f95c1913f9c72fc1f4ef0847e1e723ffe0bde0b36e5f36c13f81fe8c26ed": "b",
				"879923da020d1533f4d8e921ea7bac61e8ba41d3c89d17a4d14e3a89c6780d5d": "c",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			collection := NewCollectionWithItems[string](tc.items)
			if len(collection.items) != len(tc.result) {
				t.Errorf("expected %d, got %d", len(tc.result), len(collection.items))
			}

			for k, v := range collection.items {
				actualValue, ok := tc.result[k]
				if !ok || actualValue != v {
					t.Errorf("expected %s, got %s", tc.result[k], v)
				}
			}
		})
	}
}

func TestCollection_Len(t *testing.T) {
	tcs := []struct {
		name   string
		items  []string
		result int
	}{
		{
			name:   "a,b,c to collection",
			items:  []string{"a", "b", "c"},
			result: 3,
		},
		{
			name:   "empty collection",
			items:  []string{},
			result: 0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			collection := NewCollectionWithItems(tc.items)
			length := collection.Len()
			if length != tc.result {
				t.Errorf("expected %d, got %d", tc.result, length)
			}
		})
	}
}

func TestCollection_Items(t *testing.T) {
	tcs := []struct {
		name   string
		items  []string
		result []string
	}{
		{
			name:   "a,b,c to collection",
			items:  []string{"a", "b", "c"},
			result: []string{"a", "b", "c"},
		},
		{
			name:   "empty collection",
			items:  []string{},
			result: []string{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			collection := NewCollectionWithItems(tc.items)
			items := collection.Items()
			for _, item := range items {
				if !contains(tc.result, item) {
					t.Errorf("expected %v, got %v", tc.result, items)
				}
			}
		})
	}
}

func TestCollection_Add(t *testing.T) {
	tcs := []struct {
		name   string
		items  []string
		add    string
		result []string
	}{
		{
			name:   "a,b,c add d to collection",
			items:  []string{"a", "b", "c"},
			add:    "d",
			result: []string{"a", "b", "c", "d"},
		},
		{
			name:   "a,b,c,d add e to collection",
			items:  []string{"a", "b", "c", "d"},
			add:    "e",
			result: []string{"a", "b", "c", "d", "e"},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			collection := NewCollectionWithItems(tc.items)

			collection = collection.Add(tc.add)
			if collection.Len() != len(tc.result) {
				t.Errorf("expected %d, got %d", len(tc.result), collection.Len())
			}

			for _, item := range collection.Items() {
				if !contains(tc.result, item) {
					t.Errorf("expected %v, got %v", tc.result, collection.Items())
				}
			}
		})
	}
}
