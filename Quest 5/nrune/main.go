package main

import "fmt"

func main() {
	fmt.Println((NRune("Hello World", 5)))
}

func NRune(s string, n int) rune {
	var let rune

	for index, w := range s {
		if n > len(s) || n < 1 {
			return 0
		} else if index+1 == n {
			let = w
			break
		}
	}
	return let
}
