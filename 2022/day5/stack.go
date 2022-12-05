package main

import (
	"errors"
	"sync"
)

type CmStack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s    []string
}

func NewStack() *CmStack {
	return &CmStack{sync.Mutex{}, make([]string, 0)}
}

func (s *CmStack) Push(v string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *CmStack) Pop() (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return "", errors.New("empty stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *CmStack) Peek() (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return "", errors.New("empty stack")
	}

	res := s.s[l-1]
	// s.s = s.s[:l-1]
	return res, nil
}

// func main() {
// 	s := NewStack()
// 	s.Push(1)
// 	s.Push(2)
// 	s.Push(3)
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())
// }
