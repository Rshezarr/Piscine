package main

import "fmt"

func main() {
	fmt.Println(BasicAtoi("12345"))     //ouput must be 12345
	fmt.Println(BasicAtoi("000012345")) //ouput must be 12345
}

func BasicAtoi(s string) int {
	var n int
	for _, w := range s {
		n = n*10 + int(w) - 48
	}
	return n
}
