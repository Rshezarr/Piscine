package main

import "fmt"

func main() {
	fmt.Println(Compare("Hello", "e"))
}

func Compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}
