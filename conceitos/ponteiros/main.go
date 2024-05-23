package main

import "fmt"

func increment(n *int) { // Essa função recebe um endereço e faz um ponteiro.
	*n++ // Modifica o valor do endereço.
}

func main() {
	// Ponteiros armazenam endereços de memória de outra variável.

	var x int = 10     // Não é necessário colocar o tipo, já que já está recebendo o valor certinho.
	var p *int = &x    // Cria um ponteiro que recebe o endereço de x.
	var value int = *p // Recebe o valor de x através do ponteiro.
	*p = 20            // Muda o valor de x com o ponteiro.
	// O p = ... Não pode acontecer, por 'p' guarda o endereço.
	// O * faz a desreferenciação do ponteiro, indicando o valor da variável.
	// Podendo assim modificar x a partir do endereço 'p', com *p.
	// Desreferenciamento de Ponteiros

	fmt.Println(value) // Imprime o valor normalmente.
	fmt.Println(p)     // Imprime o endereço.
	fmt.Println(*p)    // Imprime o valor.
	fmt.Println(&x)    // Imprime o endereço.
	fmt.Println(x)     // Imprime o valor.

	increment(p)  // Manda o endereço.
	increment(&x) // Manda o endereço.
	fmt.Println(x)

	var po = Point{ID: 1, Name: "Antônio"}

	po.Add() // Usando um struct com ponteiro.
	fmt.Println(po)

	// ARRAYS
	// Para trabalhar com arrays é necessário usar (*array)[posição]
	// Exemplo: Arquivo do Bubble Sort.
}
