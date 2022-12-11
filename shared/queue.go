package shared

import (
	"container/list"
	"fmt"
)

// TODO: this should allow arbitrary types, but ok for now

func NewIntQueue() *IntQueue {
	return &IntQueue{
		queue: list.New(),
	}
}

type IntQueue struct {
	queue *list.List
}

func (q *IntQueue) Enqueue(value int) {
	q.queue.PushBack(value)
}

func (q *IntQueue) Dequeue() (int, error) {
	if q.queue.Len() > 0 {
		if val, ok := q.queue.Front().Value.(int); ok {
			return val, nil
		}
		return 0, fmt.Errorf("error: datatype incorrect while dequeuing")
	}
	return 0, fmt.Errorf("error: empty queue")
}

func (q *IntQueue) Peek() (int, error) {
	if q.queue.Len() > 0 {
		if val, ok := q.queue.Front().Value.(int); ok {
			return val, nil
		}
		return 0, fmt.Errorf("error: datatype incorrect while peeking")
	}
	return 0, fmt.Errorf("error: empty queue")
}

func (q *IntQueue) Size() int {
	return q.queue.Len()
}

func (q *IntQueue) Empty() bool {
	return q.queue.Len() == 0
}
