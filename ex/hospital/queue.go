package main

import (
	"fmt"
)

type Node struct {
	value string
	next  *Node
}

type Queue struct {
	front *Node
	back  *Node
	index int
}

func (q *Queue) Enqueue(value string) {
	newNode := &Node{value: value}
	if q.front == nil {
		q.front = newNode // q.front.value = value, q.front.next = nil
		q.back = newNode  // q.back.value = value, q.back.next = nil
		// Em suma, q.front é igual a q.back.
	} else {
		// Na segunda volta, o front e o back são os mesmos, então o front.next também recebe o newNode.
		q.back.next = newNode
		q.back = newNode
	}
	q.index++
}

func (q *Queue) Dequeue() (string, error) {
	if q.front == nil {
		return "0", fmt.Errorf("error")
	}
	value := q.front.value
	q.front = q.front.next
	if q.front == nil {
		q.back = nil
	}
	return value, nil
}

func (q *Queue) Peek() (string, error) {
	if q.front == nil {
		return "0", fmt.Errorf("error")
	}
	return q.front.value, nil
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

func (q *Queue) Print() {
	current := q.front
	id := 1
	for current != nil {
		fmt.Print(id, "º: ", current.value, "\n")
		current = current.next
		id++
	}
}
