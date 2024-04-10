package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var s string
	count := 0
	for i := 0; i < n; i++ {
		fmt.Scanln(&s)
		lastIndex := 0
		for _, c := range s {
			if lastIndex > strings.IndexRune(s, c) {
				count -= 1
				break
			} else {
				lastIndex = strings.IndexRune(s, c)
			}

		}
		count += 1
	}
	fmt.Print(count)
}
