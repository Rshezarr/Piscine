package main

import "fmt"

func main() {
	fmt.Println(FindNextPrime(16)) //output must be 17
}

func FindNextPrime(nb int) int {
	prime := 0
	if nb <= 1 {
		return FindNextPrime(nb + 1)
	}

	prime = nb / 2
	for i := 2; i <= prime; i++ {
		if nb%i == 0 {
			return FindNextPrime(nb + 1)
		}
	}
	return nb
}
