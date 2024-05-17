package main

import "fmt"

type Point struct {
	ID   int
	Name string
}

func (p *Point) Add() { // Criando um método que recebe um ponteiro.
	var other Point
	other = *p // Variável recebe os valores de p.

	fmt.Println(other)

	// Não preciso usar o * para pegar os atributos da struct.
	p.ID++         // Modificando atributo de p.
	p.Name = "Add" //...
}
