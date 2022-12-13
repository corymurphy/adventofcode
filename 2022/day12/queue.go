package main

import (
	"container/list"
	"fmt"
)

// TODO: this should allow arbitrary types, but ok for now

func NewQueue() *Queue {
	return &Queue{
		queue: list.New(),
	}
}

type Queue struct {
	queue *list.List
}

func (q *Queue) Enqueue(value *Node) {
	q.queue.PushBack(value)
}

func (q *Queue) Dequeue() (*Node, error) {
	if q.queue.Len() > 0 {
		element := q.queue.Front()
		if val, ok := element.Value.(*Node); ok {
			q.queue.Remove(element)
			return val, nil
		}
		return &Node{}, fmt.Errorf("error: datatype incorrect while dequeuing")
	}
	return &Node{}, fmt.Errorf("error: empty queue")
}

// func (q *Queue) Peek() (interface{}, error) {
// 	if q.queue.Len() > 0 {

// 		return q.queue.Front()
// 		// if val, ok := q.queue.Front().Value.(int); ok {
// 		// 	return val, nil
// 		// }
// 		// return 0, fmt.Errorf("error: datatype incorrect while peeking")
// 	}
// 	return 0, fmt.Errorf("error: empty queue")
// }

func (q *Queue) Size() int {
	return q.queue.Len()
}

func (q *Queue) Empty() bool {
	return q.queue.Len() == 0
}
