package main

import "fmt"

func main() {
	arr := []string{"One", "Two", "Three"}
	fmt.Println(BasicJoin(arr))
}

func BasicJoin(elems []string) string {
	var str string
	for i := range elems {
		str += elems[i]
	}
	return str
}
