package circle

func New[T comparable](vals ...T) *Entry[T] {
	var first, last *Entry[T]
	for _, v := range vals {
		e := &Entry[T]{
			value: v,
		}
		if first == nil {
			first = e
		}
		if last != nil {
			last.next = e
		}
		last = e
	}
	last.next = first
	return first
}

type Snippet[T comparable] struct {
	first, last *Entry[T]
}

type Entry[T comparable] struct {
	value T
	next  *Entry[T]
}

func (e *Entry[T]) Next() *Entry[T] {
	return e.next
}

// Snip removes the n entries immediately to the right of e and returns that Snippet, repairing the circle such that
// e's Next() is now the first entry after the snippet.
func (e *Entry[T]) Snip(n int) *Snippet[T] {
	var first, last *Entry[T]
	cur := e
	for i := 0; i < n; i++ {
		cur = cur.Next()
		if first == nil {
			first = cur
		}
		last = cur
	}
	e.next = last.Next()
	last.next = nil
	return &Snippet[T]{
		first: first,
		last:  last,
	}
}

// Insert inserts the given snippet immediately to the right of e, repairing the circle such that e's Next() is the
// first entry in the snippet and e's previous Next() is now the last entry in the Snippet's Next().
func (e *Entry[T]) Insert(snippet *Snippet[T]) {
	snippet.last.next = e.next
	e.next = snippet.first
}

// Find traverses the circle starting at e, returning the first entry whose value is v, or nil if no such entry exists.
func (e *Entry[T]) Find(v T) *Entry[T] {
	if e.Value() == v {
		return e
	}
	for cur := e.Next(); cur != e; cur = cur.Next() {
		if cur.Value() == v {
			return cur
		}
	}
	return nil
}

func (e *Entry[T]) Value() T {
	return e.value
}
