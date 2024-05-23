package main

import (
	"fmt"
)

var (
	children = Queue{}
	priority = Queue{}
	normal   = Queue{}
	person   = Person{}
	option   int
)

func menuStart() {
	fmt.Println("\n[1] - Functionary")
	fmt.Println("[2] - Patient")
	fmt.Println("[0] - Leave")
}

func menuPatient() {
	fmt.Println("\n[1] - Child")
	fmt.Println("[2] - Priority")
	fmt.Println("[3] - Normal")
	fmt.Println("[0] - Leave")
}

func menuFunctionary() {
	fmt.Println("\n[1] - Assist")
	fmt.Println("[2] - List")
	fmt.Println("[3] - List Done")
	fmt.Println("[3] - Leave")
}

func patientOption() {
	switch option {
	case 1:
		children.Enqueue(fmt.Sprint("C", priority.index+1))
		break
	case 2:
		priority.Enqueue(fmt.Sprint("P", priority.index+1))
		break
	case 3:
		normal.Enqueue(fmt.Sprint("N", normal.index+1))
		break
	case 0:
		break
	default:
		fmt.Println("Invalid option")
		break
	}
}

func functionaryOption() {
	switch option {
	case 1:
		menuPatient()
		fmt.Print("Option: ")
		fmt.Scan(&option)
		functionaryAssist()
		break
	case 2:
		menuPatient()
		fmt.Print("Option: ")
		fmt.Scan(&option)
		functionaryList()
		break
	case 3:
		ListPeople()
		break
	case 0:
		break
	default:
		fmt.Println("Invalid Option")
		break
	}
}

func functionaryList() {
	switch option {
	case 1:
		children.Print()
		break
	case 2:
		priority.Print()
		break
	case 3:
		normal.Print()
		break
	case 0:
		break
	default:
		fmt.Println("Invalid option")
		break
	}
}

func functionaryAssist() {
	fmt.Print("Document: ")
	fmt.Scan(&person.document)
	fmt.Print("Name: ")
	fmt.Scan(&person.name)
	fmt.Print("Age: ")
	fmt.Scan(&person.age)

	switch option {
	case 1:
		person.ticket, _ = children.Dequeue()
		break
	case 2:
		person.ticket, _ = priority.Dequeue()
		break
	case 3:
		person.ticket, _ = normal.Dequeue()
		break
	case 0:
		break
	default:
		fmt.Println("Invalid Option.")
		break
	}

}

func main() {

	for {
		menuStart()
		fmt.Print("Option: ")
		fmt.Scan(&option)
		switch option {
		case 1:
			menuFunctionary()
			fmt.Print("Option: ")
			fmt.Scan(&option)
			functionaryOption()
			break
		case 2:
			menuPatient()
			fmt.Print("Option: ")
			fmt.Scan(&option)
			patientOption()
			break
		case 0:
			break
		default:
			fmt.Println("Invalid option")
			break
		}
		if option == 0 {
			break
		}
	}

}
