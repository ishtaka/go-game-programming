package search

import "sync"

type Queue[T any] struct {
	elems []T
	lock  sync.Mutex
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		elems: make([]T, 0, size),
	}
}

func (q *Queue[T]) Enqueue(elem T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.elems = append(q.elems, elem)
}

func (q *Queue[T]) Dequeue() T {
	q.lock.Lock()
	defer q.lock.Unlock()

	elem := q.elems[0]
	q.elems = q.elems[1:]

	return elem
}

func (q *Queue[T]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.elems) == 0
}
