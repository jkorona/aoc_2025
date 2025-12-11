package utils

type Queue[T any] struct {
	items []T
	head  int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.head >= len(q.items) {
		return zero, false
	}
	v := q.items[q.head]
	q.head++

	// Optional memory cleanup: compact when head is large
	if q.head > 64 && q.head*2 >= len(q.items) {
		q.items = append([]T{}, q.items[q.head:]...)
		q.head = 0
	}
	return v, true
}

func (q *Queue[T]) Len() int {
	return len(q.items) - q.head
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}
