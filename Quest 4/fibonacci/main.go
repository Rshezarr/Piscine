package main

import "fmt"

func main() {
	fmt.Println(Fibonacci(15))
}

func Fibonacci(index int) int {
	if index == 1 {
		return 1
	} else if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
