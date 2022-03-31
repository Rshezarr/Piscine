package main

import "fmt"

func main() {
	fmt.Print(StrRev("Hello World!"))
}

func StrRev(s string) string {
	strrev := ""
	for _, w := range s {
		strrev = string(w) + strrev
	}

	return strrev
}
