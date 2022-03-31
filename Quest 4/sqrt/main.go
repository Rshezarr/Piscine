package main

import "fmt"

func main() {
	fmt.Println(Sqrt(64)) //√64 == 8
}

func Sqrt(nb int) int {
	for i := 0; i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
