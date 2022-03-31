package main

import "github.com/01-edu/z01"

func main() {
	IsNegative(-5) // T
	IsNegative(5)  // F
}

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
