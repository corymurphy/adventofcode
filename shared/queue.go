package shared

import (
	"container/list"
	"fmt"
)

// TODO: this should allow arbitrary types, but ok for now

type IntQueue struct {
	queue *list.List
}

func (q *IntQueue) Enqueue(value int) {
	q.queue.PushBack(value)
}

func (q *IntQueue) Dequeue() error {
	if q.queue.Len() > 0 {
		ele := q.queue.Front()
		q.queue.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}

func (q *IntQueue) Front() (int, error) {
	if q.queue.Len() > 0 {
		if val, ok := q.queue.Front().Value.(int); ok {
			return val, nil
		}
		return 0, fmt.Errorf("Peep Error: Queue Datatype is incorrect")
	}
	return 0, fmt.Errorf("Peep Error: Queue is empty")
}

func (q *IntQueue) Size() int {
	return q.queue.Len()
}

func (q *IntQueue) Empty() bool {
	return q.queue.Len() == 0
}
