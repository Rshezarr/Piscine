package main

import "fmt"

func main() {
	fmt.Println(FirstRune("Hello World"))
}

func FirstRune(s string) rune {
	runes := []rune(s)
	return runes[0]
}
