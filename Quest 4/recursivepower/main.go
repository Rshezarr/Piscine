package main

import "fmt"

func main() {
	fmt.Println(RecursivePower(2, 4)) // 2^4 == 16
}

func RecursivePower(nb, power int) int {
	num := 1

	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else {
		num = nb * (RecursivePower(nb, power-1))
	}
	return num
}
