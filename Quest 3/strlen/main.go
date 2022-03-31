package main

import "fmt"

func main() {
	fmt.Print(StrLen(("Hello World!"))) // output must be 12
}

func StrLen(s string) int {
	var count int

	for range s {
		count++
	}
	return count
}
