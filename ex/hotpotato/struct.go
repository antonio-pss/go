package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

type CircleList struct {
	head *Node
	size int
}

func (list *CircleList) Append(value string) {
	newNode := &Node{value: value}
	if list.head == nil {
		list.head = newNode
		newNode.next = list.head
	} else {
		current := list.head
		for current.next != list.head {
			current = current.next
		}
		current.next = newNode
		newNode.next = list.head
	}
	list.size++
}

func (list *CircleList) Remove(value string) {
	if list.head == nil {
		return
	}

	current := list.head
	prev := list.head

	for current.value != value {
		if current.next == list.head {
			return // Valor não foi encontrado.
		}
		prev = current
		current = current.next
	}

	if current == list.head {
		nextNode := list.head.next
		for prev.next != list.head {
			prev = prev.next
		}
		list.head = nextNode
		prev = prev.next
		list.size--
	} else {
		prev.next = current.next
		list.size--
	}
}

func (list *CircleList) Print() {
	if list.head == nil {
		return
	}
	current := list.head
	for {
		fmt.Print(current.value, " -> ")
		current = current.next
		if current == list.head {
			break
		}
	}
	fmt.Println("(first)")
}
