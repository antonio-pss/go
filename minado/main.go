package main

import (
	"fmt"
	"math/rand"
)

var campo [10][10]bool
var m [10][10]bool
var pontos = 0

func menu() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if m[i][j] {
				fmt.Print("|-|")
			} else if campo[i][j] && !m[i][j] {
				fmt.Print("|+|")
			} else if !campo[i][j] && m[i][j] {
				fmt.Print("|*|")
			}
		}
		fmt.Println()
	}
}

func menuFinal() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if !campo[i][j] {
				fmt.Print("|*|")
			} else if m[i][j] {
				fmt.Print("|-|")
			} else if campo[i][j] && !m[i][j] {
				fmt.Print("|+|")
			}
		}
		fmt.Println()
	}
}

func ganhou(linha, coluna *int) bool {
	if !campo[*linha-1][*coluna-1] {
		return false
	} else {
		pontos += 10
		m[*linha-1][*coluna-1] = false
		return true
	}
}

func main() {

	var linha int
	var coluna int
	var minas int
	var quantidade = 0

	fmt.Print("Quantas minas quer no campo? ")
	fmt.Scan(&minas)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			campo[i][j] = true
			m[i][j] = true
		}
	}
	for {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if !campo[i][j] {
					quantidade++
				}
			}
		}
		if quantidade < minas {
			linha = rand.Intn(9)
			coluna = rand.Intn(9)
			campo[linha][coluna] = false
			quantidade = 0
		} else {
			break
		}
	}

	for {
		menu()
		fmt.Print("Opção Linha (1 a 10): ")
		fmt.Scan(&linha)
		fmt.Print("Opção Coluna (1 a 10): ")
		fmt.Scan(&coluna)

		if !ganhou(&linha, &coluna) {
			fmt.Println("Você acertou a BOMBA!")
			fmt.Println("Pontos:", pontos)
			menuFinal()
			break
		}
		fmt.Println("Pontos:", pontos)
	}

}
