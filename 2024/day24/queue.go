package main

import (
	"container/list"
	"fmt"
)

func NewQueue() *Queue {
	return &Queue{
		queue: list.New(),
	}
}

type Queue struct {
	queue *list.List
}

func (q *Queue) Print() {
	fmt.Println(q.queue)
}

func (q *Queue) Enqueue(value *Instruction) {
	q.queue.PushBack(value)
}

func (q *Queue) Dequeue() (*Instruction, error) {
	if q.queue.Len() > 0 {
		element := q.queue.Front()
		if val, ok := element.Value.(*Instruction); ok {
			q.queue.Remove(element)
			return val, nil
		}
		return nil, fmt.Errorf("error: datatype incorrect while dequeuing")
	}
	return nil, fmt.Errorf("error: empty queue")
}

func (q *Queue) Size() int {
	return q.queue.Len()
}

func (q *Queue) Empty() bool {
	return q.queue.Len() == 0
}
