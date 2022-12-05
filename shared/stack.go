package main

import (
	"errors"
	"sync"
)

type Stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s    []string
}

func NewStack() *Stack {
	return &Stack{sync.Mutex{}, make([]string, 0)}
}

func (s *Stack) Push(v string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *Stack) Pop() (string, error) {
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

func (s *Stack) Peek() (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return "", errors.New("empty stack")
	}

	res := s.s[l-1]
	return res, nil
}
