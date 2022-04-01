package main

import "fmt"

func main() {
	fmt.Println(AlphaCount("Hello World!"))
}

func AlphaCount(s string) int {
	var count int
	for _, w := range s {
		if w >= 'a' && w <= 'z' || w >= 'A' && w <= 'Z' {
			count++
		}
	}
	return count
}
