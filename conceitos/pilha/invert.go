package main

import "fmt"

type nodeString struct {
	value string
	next  *nodeString
}

type stackString struct {
	top *nodeString
}

func (s *stackString) push(value string) {
	newNode := &nodeString{value: value, next: s.top}
	s.top = newNode
}

func (s *stackString) print() {
	current := s.top
	for current != nil {
		fmt.Print(current.value)
		current = current.next
	}
}

func invert() {
	var stack stackString
	var word string
	fmt.Print("Enter word: ")
	fmt.Scan(&word)

	for _, char := range word {
		stack.push(string(char))
	}

	stack.print()
}
