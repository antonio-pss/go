package main

import "fmt"

func rec(n int) int {
	if n == 1 {
		return n
	} else if n%2 == 0 {
		n = n / 2
	} else {
		n = n + 1
	}
	fmt.Println(n)
	return rec(n)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(10))
}
