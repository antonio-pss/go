package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Candidate struct {
	Name   string
	Number int
}

var Candidates []Candidate

func LoadCandidates() {
	file, err := os.Open("data\\candidates.txt")
	if err != nil {
		fmt.Println("Load Candidates cannot work properly: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		collums := strings.Split(line, "|")
		if len(collums) > 0 {
			name := collums[0]
			number, _ := strconv.Atoi(collums[1])

			Candidates = append(Candidates, Candidate{name, number})
		}
	}
}

func ViewCandidates() {
	for _, c := range Candidates {
		fmt.Printf("Name: %s | Number: %d \n", c.Name, c.Number)
	}
}
