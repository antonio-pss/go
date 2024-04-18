package main

import (
	"bufio"
	"election/utils"
	"fmt"
	"os"
)

func menu() {
	fmt.Print("\n\n[1] - Add Candidate + Vote.")
	fmt.Print("\n[2] - List Candidates.")
	fmt.Print("\n[3] - List Voters.")
	fmt.Print("\n[4] - List Votes.")
	fmt.Print("\n[5] - Vote Curation.")
	fmt.Print("\n[0] - Save and Leave.\n")
}

func main() {
	var (
		option  int
		vote    utils.Vote
		scanner = bufio.NewScanner(os.Stdin)
	)
	utils.LoadCandidates()
	utils.LoadVoters()
	utils.LoadVotes()

	for {
		menu()
		fmt.Print("Option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Print("Enter Voter Name: ")
			scanner.Scan()
			vote.Voter.Name = scanner.Text()

			fmt.Print("Enter Voter Document: ")
			scanner.Scan()
			vote.Voter.Document = scanner.Text()

			fmt.Print("Enter Candidate Number: ")
			fmt.Scan(&vote.Candidate.Number)

			vote.Add()
			break

		case 2:
			utils.ViewCandidates()
			break
		case 3:
			utils.ViewVoters()
			break
		case 4:
			utils.ViewVotes()
			break
		case 5:
			utils.VoteCuration()
			break
		case 0:
			break
		default:
			fmt.Println("Invalid option.")
		}

		if option == 0 {
			break
		}
	}

	utils.SaveVoters()
	utils.SaveVotes()
}
