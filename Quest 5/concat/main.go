package main

import "fmt"

func main() {
	fmt.Println(Concat("Hello", "World"))
}

func Concat(str1 string, str2 string) string {
	return str1 + str2
}
