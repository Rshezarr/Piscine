package main

import "fmt"

func main() {
	fmt.Println(IterativeFactorial(4)) //1 * 2 * 3 * 4 == 24
}

func IterativeFactorial(nb int) int {
	var factorial int

	if nb < 20 || nb > 0 {
		for i := 1; i <= nb; i++ {
			factorial = factorial * i
		}
		return factorial
	}
	return 0
}
