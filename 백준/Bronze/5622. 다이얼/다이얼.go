package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	t := 0
	for _, chr := range s {
		n := 0
		if chr >= 65 && chr <= 67 {
			n = 2
		} else if chr >= 68 && chr <= 70 {
			n = 3
		} else if chr >= 71 && chr <= 73 {
			n = 4
		} else if chr >= 74 && chr <= 76 {
			n = 5
		} else if chr >= 77 && chr <= 79 {
			n = 6
		} else if chr >= 80 && chr <= 83 {
			n = 7
		} else if chr >= 84 && chr <= 86 {
			n = 8
		} else if chr >= 87 && chr <= 90 {
			n = 9
		}
		t += n + 1
	}
	fmt.Print(t)
}
