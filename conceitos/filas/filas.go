package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	front *Node
	back  *Node
	size  int
}

func (q *Queue) Enqueue(value int) {
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
	q.size++
}

func (q *Queue) Dequeue() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("error")
	}
	value := q.front.value
	q.front = q.front.next
	if q.front == nil {
		q.back = nil
	}
	q.size--
	return value, nil
}

func (q *Queue) Peek() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("error")
	}
	return q.front.value, nil
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}
