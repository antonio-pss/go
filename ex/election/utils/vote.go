package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vote struct {
	ID        int
	Voter     Voter
	Candidate Candidate
}

var Votes []Vote
var CurrentID int

func LoadVotes() {
	file, err := os.Open("data\\votes.txt")
	if err != nil {
		fmt.Println("Load Votes cannot work properly: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, "|")

		if len(columns) > 0 {
			id, _ := strconv.Atoi(columns[0])
			voter := strings.Split(columns[1], ":")
			candidate := strings.Split(columns[2], ":")
			candidateNumber, _ := strconv.Atoi(candidate[1])

			Votes = append(Votes, Vote{id, Voter{voter[0], voter[1]}, Candidate{Name: candidate[0], Number: candidateNumber}})

			CurrentID = id + 1
		} else {
			CurrentID = 1
		}
	}
}

func SaveVotes() {
	file, err := os.OpenFile("data\\votes.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Save Votes cannot work properly: ", err)
		return
	}
	defer file.Close()

	for _, vote := range Votes {
		line := strconv.Itoa(vote.ID) + "|" + vote.Voter.Name + ":" + vote.Voter.Document + "|" + vote.Candidate.Name + ":" + strconv.Itoa(vote.Candidate.Number) + "\n"
		_, err = file.WriteString(line)
		if err != nil {
			fmt.Println("Save Votes cannot work properly: ", err)
			return
		}
	}
}

func (v Vote) Add() {
	v.ID = CurrentID
	for i, candidate := range Candidates {
		if candidate.Number == v.Candidate.Number {
			v.Candidate.Name = candidate.Name
			if v.Voter.add() {
				Votes = append(Votes, v)
			} else {
				fmt.Println("Voter already voted.")
			}
			break
		} else if i == len(Candidates)-1 {
			fmt.Println("Candidate do not exists")
		}
	}
}

func ViewVotes() {
	for _, vote := range Votes {
		fmt.Printf("Vote ID: %d | Voter Name: %s | Voter Document: %s | Candidate Name: %s | Candidate Number: %d \n", vote.ID, vote.Voter.Name, vote.Voter.Document, vote.Candidate.Name, vote.Candidate.Number)
	}
}

func VoteCuration() {
	var curation = make(map[int]int)
	for _, j := range Votes {
		curation[j.Candidate.Number]++
	}
	for key, value := range curation {
		fmt.Print("Candidate Number: ", key, " | Candidate Votes: ", value, " | Percentage: ", (value*100)/len(Votes), "%\n")
	}
}
