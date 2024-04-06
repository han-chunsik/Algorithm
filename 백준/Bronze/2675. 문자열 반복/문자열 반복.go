package main

import (
	"fmt"
)

func repeatString(char string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += char
	}
	return result
}

func main() {
	var a, r int
	fmt.Scanln(&a)

	var s string

	for i := 0; i < a; i++ {
		fmt.Scanln(&r, &s)
		p := ""
		for _, char := range s {
			p += repeatString(string(char), r)
		}
		fmt.Println(p)
	}
}
