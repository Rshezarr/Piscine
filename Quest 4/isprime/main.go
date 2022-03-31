package main

import "fmt"

func main() {
	fmt.Println(IsPrime(15)) // false
	fmt.Println(IsPrime(17)) // true
}

func IsPrime(nb int) bool {
	prime := 0

	if nb <= 1 {
		return false
	}

	prime = nb / 2
	for i := 2; i <= prime; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
