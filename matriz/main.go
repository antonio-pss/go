package main

import "fmt"

func menu(matriz [3][3]string) {
	for i := 0; i < 3; i++ {
		fmt.Printf(" %s | %s | %s\n", matriz[i][0], matriz[i][1], matriz[i][2])
		if i < 2 {
			fmt.Println("----------")
		}
	}
}

func ganhou(matriz [3][3]string) string {
	for i := 0; i < 3; i++ {
		if matriz[i][0] == matriz[i][1] && matriz[i][0] == matriz[i][2] {
			return quem(matriz[i][0])
		} else if matriz[0][i] == matriz[1][i] && matriz[0][i] == matriz[2][i] {
			return quem(matriz[0][i])
		} else if matriz[0][0] == matriz[1][1] && matriz[0][0] == matriz[2][2] {
			return quem(matriz[i][0])
		}
	}
	return "N"
}

func quem(jogador string) string {
	if jogador == "X" {
		return "X"
	} else if jogador == "O" {
		return "O"
	} else {
		return "N"
	}
}

func main() {
	var matriz = [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	var cond = true
	var coluna int
	var linha int
	var contador = 0

	for contador < 9 {
		menu(matriz)
		fmt.Print("Linha da Matriz (0 a 2): ")
		fmt.Scan(&linha)
		fmt.Print("Coluna da Matriz (0 a 2): ")
		fmt.Scan(&coluna)

		if cond && matriz[linha][coluna] == " " {
			matriz[linha][coluna] = "X"
			contador += 1
			cond = !cond
		} else if matriz[linha][coluna] == " " {
			matriz[linha][coluna] = "O"
			contador += 1
			cond = !cond
		}

		if ganhou(matriz) == "X" {
			menu(matriz)
			fmt.Println("O jogador X ganhou.")
			break
		} else if ganhou(matriz) == "O" {
			menu(matriz)
			fmt.Println("O jogador O ganhou.")
			break
		} else {
			fmt.Println()
		}
	}
}
