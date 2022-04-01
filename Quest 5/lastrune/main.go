package main

import "fmt"

func main() {
	fmt.Println(LastRune("Hello World"))
}

func LastRune(s string) rune {
	runes := []rune(s)
	return runes[len(runes)-1]
}
