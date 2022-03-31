package main

import "fmt"

func main() {
	fmt.Println(RecursiveFactorial(4)) //1 * 2 * 3 * 4 == 24
}

func RecursiveFactorial(nb int) int {
	if nb > 20 || nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}
