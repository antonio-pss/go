package main

import (
	"fmt"
	"math/rand"
)

var jogo [3][3]string
var certo = [3][3]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", " "},
}

func menu() {
	for i := 0; i < 3; i++ {
		fmt.Print("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(jogo[i][j], " ")
		}
		fmt.Println("]")
	}
}

func tem(num *int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if jogo[i][j] == fmt.Sprint(*num) {
				return true
			}
		}
	}
	return false
}

func preenchida() bool {
	var cont int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if jogo[i][j] != "" {
				cont++
			}
		}
	}
	if cont == 9 {
		return true
	} else {
		return false
	}
}

func mudar(op *string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if jogo[i][j] == " " {
				if i-1 > -1 && jogo[i-1][j] == *op {
					jogo[i-1][j], jogo[i][j] = jogo[i][j], jogo[i-1][j]
					return true
				} else if i+1 < 3 && jogo[i+1][j] == *op {
					jogo[i+1][j], jogo[i][j] = jogo[i][j], jogo[i+1][j]
					return true
				} else if j-1 > -1 && jogo[i][j-1] == *op {
					jogo[i][j-1], jogo[i][j] = jogo[i][j], jogo[i][j-1]
					return true
				} else if j+1 < 3 && jogo[i][j+1] == *op {
					jogo[i][j+1], jogo[i][j] = jogo[i][j], jogo[i][j+1]
					return true
				}
			}
		}
	}
	return false
}

func ganhou() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if jogo[i][j] != certo[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	var aleatorio = rand.Intn(8) + 1
	var op string

	jogo[2][2] = " "

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for {
				if tem(&aleatorio) && !preenchida() {
					aleatorio = rand.Intn(8) + 1
				} else {
					break
				}
			}
			if i != 2 || j != 2 {
				jogo[i][j] = fmt.Sprint(aleatorio)
			}
		}
	}

	for {
		menu()
		fmt.Print("Qual número quer mudar? ")
		fmt.Scan(&op)

		if !mudar(&op) {
			fmt.Println("Não foi possível mudar.")
		}
		if ganhou() {
			fmt.Println("PARABÉNS, VOCÊ CONSEGUIU!")
			menu()
			break
		}
	}

}
