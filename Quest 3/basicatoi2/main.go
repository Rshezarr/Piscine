package main

import "fmt"

func main() {
	fmt.Println(BasicAtoi2("12345"))        // output must be 12345
	fmt.Println(BasicAtoi2("123 45"))       // output must be 0
	fmt.Println(BasicAtoi2("Hello World!")) // output must be 0
}

func BasicAtoi2(s string) int {
	var n int

	for _, w := range s {
		if w >= '0' && w <= '9' {
			n = n*10 + int(w) - 48
		} else {
			return 0
		}
	}
	return n
}
