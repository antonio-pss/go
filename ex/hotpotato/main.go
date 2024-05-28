package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	hotPotato = CircleList{}
	kids      = []string{
		"Antônio", "Luana", "Alex", "Paulo", "Junior", "Gustavo", "Marcelo",
	}
)

func appendKids() {
	for _, kid := range kids {
		hotPotato.Append(kid)
	}
}

func main() {
	appendKids()
	fmt.Print("The player are: ")
	hotPotato.Print()
	time.Sleep(5 * time.Second)

	for hotPotato.size != 1 {
		var number = rand.Intn(len(kids))
		fmt.Print("The potato was catch by ", kids[number], "\n")
		hotPotato.Remove(kids[number])
		kids = append(kids[:number], kids[number+1:]...)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("The winner is", hotPotato.head.value)
}
