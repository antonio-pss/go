package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Voter struct {
	Name     string
	Document string
}

var Voters []Voter

func LoadVoters() {
	file, err := os.Open("data\\voters.txt")
	if err != nil {
		fmt.Println("Load Voters cannot work properly: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columms := strings.Split(line, "|")
		if len(columms) > 0 {
			Voters = append(Voters, Voter{Name: columms[0], Document: columms[1]})
		}
	}
}

func SaveVoters() {
	file, err := os.OpenFile("data\\voters.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Save cannot work properly: ", err)
		return
	}
	defer file.Close()

	for _, voter := range Voters {
		line := voter.Name + "|" + voter.Document + "\n"
		_, err = file.WriteString(line)
		if err != nil {
			fmt.Println("Save cannot work properly: ", err)
			return
		}
	}
}

func (v Voter) add() bool {
	for _, voter := range Voters {
		if voter.Document == v.Document {
			return false
		}
	}
	Voters = append(Voters, v)
	return true
}

func ViewVoters() {
	for _, voter := range Voters {
		fmt.Println("Voter Name:", voter.Name, "| Document:", voter.Document)
	}
}

func (v Voter) Remove() {
	for index, Voter := range Voters {
		if Voter.Name == v.Name {
			Voters = append(Voters[:index], Voters[index+1:]...)
		}
	}
}
