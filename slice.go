package golinq

// Slice - wrapper over standard go slice for methods.
// Generic type T must be comparable.
type Slice[T comparable] struct {
	items []T
}

// NewSlice - creates new Slice instance.
//
// args:
//
//	items - slice items.
//
// example:
//
//	s := golinq.NewSlice([]int{1, 2, 3})
func NewSlice[T comparable](items []T) *Slice[T] {
	return &Slice[T]{items: items}
}

func (s *Slice[T]) Len() int {
	return len(s.items)
}

// Items - returns all slice items.
func (s *Slice[T]) Items() []T {
	return s.items
}

func (s *Slice[T]) GetItemByIndex(index int) T {
	if index < 0 || index >= len(s.items) {
		return *new(T)
	}
	return s.items[index]
}

func (s *Slice[T]) GetFirstIndex(item T) int {
	for i, v := range s.items {
		if v == item {
			return i
		}
	}
	return -1
}

// GetLastIndex - returns last index of item.
func (s *Slice[T]) GetLastIndex(item T) int {
	for i := len(s.items) - 1; i >= 0; i-- {
		if s.items[i] == item {
			return i
		}
	}
	return -1
}

func (s *Slice[T]) GetAllIndexes(item T) []int {
	var indexes []int
	for i, v := range s.items {
		if v == item {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func (s *Slice[T]) Add(item T) *Slice[T] {
	items := append(s.items, item)
	return &Slice[T]{items: items}
}

func (s *Slice[T]) RemoveFirst(item T) *Slice[T] {
	var items []T
	for i, v := range s.items {
		if v == item {
			items = append(s.items[:i], s.items[i+1:]...)
			return &Slice[T]{items: items}
		}
	}
	return s
}

func (s *Slice[T]) Remove(item T) *Slice[T] {
	var indexesToDelete []int

	for i, v := range s.items {
		if v == item {
			indexesToDelete = append(indexesToDelete, i)
		}
	}

	if len(indexesToDelete) == 0 {
		return s
	}

	items := make([]T, len(s.items)-len(indexesToDelete))

	for oldSliceIndex, newSliceIndex := 0, 0; oldSliceIndex < len(s.items); oldSliceIndex++ {
		if contains(indexesToDelete, oldSliceIndex) {
			continue
		}

		items[newSliceIndex] = s.items[oldSliceIndex]
		newSliceIndex++
	}

	return &Slice[T]{items: items}
}

func (s *Slice[T]) Filter(fn func(T) bool) *Slice[T] {
	items := make([]T, 0)
	for _, item := range s.items {
		if fn(item) {
			items = append(items, item)
		}
	}

	return &Slice[T]{items: items}
}

func (s *Slice[T]) Contains(item T) bool {
	return contains(s.items, item)
}

func contains[T comparable](items []T, item T) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}
