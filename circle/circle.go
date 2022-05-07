package circle

func New[T comparable](vals ...T) (first *Entry[T], lookup map[T]*Entry[T]) {
	var last *Entry[T]
	lookup = make(map[T]*Entry[T])
	for _, v := range vals {
		e := &Entry[T]{
			value: v,
		}
		lookup[v] = e
		if first == nil {
			first = e
		}
		if last != nil {
			last.next = e
		}
		last = e
	}
	last.next = first
	return first, lookup
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
func (e *Entry[T]) Snip(n int) (snippet *Snippet[T], vals map[T]bool) {
	vals = make(map[T]bool)
	var first, last *Entry[T]
	cur := e
	for i := 0; i < n; i++ {
		cur = cur.Next()
		vals[cur.Value()] = true
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
	}, vals
}

// Insert inserts the given snippet immediately to the right of e, repairing the circle such that e's Next() is the
// first entry in the snippet and e's previous Next() is now the last entry in the Snippet's Next().
func (e *Entry[T]) Insert(snippet *Snippet[T]) {
	snippet.last.next = e.next
	e.next = snippet.first
}

func (e *Entry[T]) Value() T {
	return e.value
}
