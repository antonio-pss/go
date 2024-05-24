package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
}

func (l *List) Append(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (l *List) Remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		return
	}

	current := l.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}

func (l *List) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
}

func main() {
	var list List
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Print()
	fmt.Println()
	list.Remove(30)
	list.Remove(50)
	list.Remove(20)
	list.Print()
}
