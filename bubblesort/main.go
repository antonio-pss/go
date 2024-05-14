package main

import (
	"fmt"
)

func bubbleSort(arr []int) []int {
	n := len(arr)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}

func main() {
	arr := []int{10, 5, 3, 8, 4, 2}
	arr = bubbleSort(arr)
	fmt.Println(arr)
}
