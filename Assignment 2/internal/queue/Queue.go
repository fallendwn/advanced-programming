package queue

type Queue[T any] struct {
	channel chan T
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		channel: make(chan T, size),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.channel <- item
}

func (q *Queue[T]) Dequeue() <-chan T {
	return q.channel
}

func (q *Queue[T]) Close() {
	close(q.channel)
}
