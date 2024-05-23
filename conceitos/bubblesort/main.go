package main

import (
	"fmt"
)

func bubbleSort1(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] < (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

func bubbleSort2(arr *[]int) {
	n := len(*arr)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if (*arr)[i-1] > (*arr)[i] {
				(*arr)[i], (*arr)[i-1] = (*arr)[i-1], (*arr)[i]
				swapped = true
			}
		}
	}
}

func main() {
	arr := []int{10, 5, 3, 8, 4, 2}
	bubbleSort1(&arr)
	fmt.Println(arr)
}
