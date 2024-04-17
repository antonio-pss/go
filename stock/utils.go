package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	ID       int32
	Name     string
	Quantity float32
	Price    float32
}

var (
	ProductList []Product
	CurrentID   int32
)

func Load() {
	file, err := os.Open("products.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, "|")
		if len(columns) > 0 {
			id, _ := strconv.ParseInt(columns[0], 10, 32)
			quantity, _ := strconv.ParseFloat(columns[2], 32)
			price, _ := strconv.ParseFloat(columns[3], 32)

			product := Product{int32(id), columns[1], float32(quantity), float32(price)}

			ProductList = append(ProductList, product)
		}
	}

	if len(ProductList) > 0 {
		CurrentID = ProductList[len(ProductList)-1].ID + 1
	} else {
		CurrentID = 1
	}
}

func Save() {
	file, err := os.OpenFile("products.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, p := range ProductList {
		line := strconv.Itoa(int(p.ID)) + "|" + p.Name + "|" + fmt.Sprintf("%.2f", p.Quantity) + "|" + fmt.Sprintf("%.2f", p.Price) + "\n"
		_, err = file.WriteString(line)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (p Product) Add() {
	p.ID = CurrentID
	ProductList = append(ProductList, p)
	CurrentID++

	file, err := os.OpenFile("products.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	line := strconv.Itoa(int(p.ID)) + "|" + p.Name + "|" + fmt.Sprintf("%.2f", p.Quantity) + "|" + fmt.Sprintf("%.2f", p.Price) + "\n"
	_, err = file.WriteString(line)
	if err != nil {
		log.Fatal(err)
	}
}

func ListProducts() {
	for _, product := range ProductList {
		fmt.Println("| Product ID: ", product.ID, "| Product Name: ", product.Name, "| Product Quantity: ", product.Quantity, "| Product Price: ", product.Price, "|")
	}
}

func (p Product) UpdateName() {
	for index, product := range ProductList {
		if product.ID == p.ID {
			ProductList[index].Name = p.Name
		}
	}
}

func (p Product) UpdateQuantity() {
	for index, product := range ProductList {
		if product.ID == p.ID {
			ProductList[index].Quantity = p.Quantity
		}
	}
}

func (p Product) UpdatePrice() {
	for index, product := range ProductList {
		if product.ID == p.ID {
			ProductList[index].Price = p.Price
		}
	}
}

func (p Product) Delete() {
	for index, product := range ProductList {
		if product.ID == p.ID {
			ProductList = append(ProductList[:index], ProductList[index+1:]...)
		}
	}
}
