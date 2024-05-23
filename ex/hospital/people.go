package main

import "fmt"

type Person struct {
	document string
	name     string
	age      int
	ticket   string
}

var People []Person

func (p *Person) Add() {
	People = append(People, *p)
}

func ListPeople() {
	for _, person := range People {
		fmt.Println("Document: ", person.document, "| Name: ", person.name, " | Age: ", person.age, " | Ticket: ", person.ticket)
	}
}
