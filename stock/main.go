package main

import (
	"bufio"
	"fmt"
	"os"
)

func menu() {
	fmt.Println("[1] - Add Product.")
	fmt.Println("[2] - List Products.")
	fmt.Println("[3] - Update Product Name.")
	fmt.Println("[4] - Update Product Quantity.")
	fmt.Println("[5] - Update Product Price.")
	fmt.Println("[6] - Delete Product.")
	fmt.Println("[0] - Leave.")
}

func main() {
	var (
		option  int
		product Product
		scanner = bufio.NewScanner(os.Stdin)
	)
	Load()

	for {
		menu()
		fmt.Print("Option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Print("Product Name: ")
			scanner.Scan()
			product.Name = scanner.Text()
			fmt.Print("Product Quantity: ")
			fmt.Scan(&product.Quantity)
			fmt.Print("Product Price: R$")
			fmt.Scan(&product.Price)

			product.Add()
			break
		case 2:
			ListProducts()
			break
		case 3:
			fmt.Print("Product ID: ")
			fmt.Scan(&product.ID)
			fmt.Print("Product New Name: ")
			scanner.Scan()
			product.Name = scanner.Text()

			product.UpdateName()
			break
		case 4:
			fmt.Print("Product ID: ")
			fmt.Scan(&product.ID)
			fmt.Print("Product New Quantity: ")
			fmt.Scan(&product.Quantity)

			product.UpdateQuantity()
			break
		case 5:
			fmt.Print("Product ID: ")
			fmt.Scan(&product.ID)
			fmt.Print("Product New Price: R$")
			fmt.Scan(&product.Price)

			product.UpdatePrice()
			break
		case 6:
			fmt.Print("Product ID: ")
			fmt.Scan(&product.ID)

			product.Delete()
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

	Save()
}
