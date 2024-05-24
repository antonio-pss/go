package main

import (
	"fmt"
)

type nodeInt struct {
	value int
	next  *nodeInt
}

type stackInt struct {
	top *nodeInt
}

func (s *stackInt) isEmpty() bool {
	return s.top == nil
}

func (s *stackInt) push(value int) {
	newNode := &nodeInt{value: value, next: s.top}
	s.top = newNode
}

func (s *stackInt) pop() (int, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	value := s.top.value
	s.top = s.top.next
	return value, nil
}

func (s *stackInt) peek() (int, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.top.value, nil
}

func (s *stackInt) print() {
	current := s.top
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func concepts() {
	var stack = stackInt{}

	stack.isEmpty()

	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.print()

	var value, _ = stack.pop()
	fmt.Println("Popped value:", value)
	stack.isEmpty()
	stack.print()
}
