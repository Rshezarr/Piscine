package main

import "fmt"

func main() {
	fmt.Println(IterativePower(2, 4)) // 2^4 == 16
}

func IterativePower(nb, power int) int {
	num := 1

	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else {
		for i := 1; i <= power; i++ {
			num = num * nb
		}
	}
	return num
}
