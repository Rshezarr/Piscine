package main

import "fmt"

func main() {
	fmt.Println(Capitalize("One two three 4four"))
}

func Capitalize(s string) string {
	new := []rune(s)

	for i, w := range s {
		if i == 0 || new[i-1] < rune('0') || new[i-1] > rune('9') && new[i-1] < rune('A') || new[i-1] > rune('Z') && new[i-1] < rune('a') || new[i-1] > rune('z') {
			if w >= rune('a') && w <= rune('z') {
				new[i] = w - 32
			}
		} else if w >= rune('A') && w <= rune('Z') && i > 0 {
			new[i] = w + 32
		}
	}
	return string(new)
}
