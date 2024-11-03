package golinq

type Collection[T any] struct {
	items map[string]T
}

func NewCollection[T any]() *Collection[T] {
	return &Collection[T]{items: make(map[string]T)}
}

func NewCollectionWithItems[T any](items []T) *Collection[T] {
	collectionItems := make(map[string]T)

	for _, item := range items {
		collectionItems[Hash(item)] = item
	}

	return &Collection[T]{items: collectionItems}
}

func (c *Collection[T]) Len() int {
	return len(c.items)
}

func (c *Collection[T]) Items() []T {
	items := make([]T, 0, len(c.items))

	for _, item := range c.items {
		items = append(items, item)
	}

	return items
}

func (c *Collection[T]) Add(item T) *Collection[T] {
	collection := c
	collection.items[Hash(item)] = item
	return collection
}
