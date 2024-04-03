package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Scanln(&s)

	var n int
	fmt.Scanln(&n)

	fmt.Println(string(s[n-1]))
}
