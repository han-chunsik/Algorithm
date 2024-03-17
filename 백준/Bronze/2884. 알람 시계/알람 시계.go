package main

import (
	"fmt"
)

func main() {
	var h, m int

	fmt.Scan(&h, &m)

	if m < 45 {
		m = 60 + (m - 45)
		if h == 0 {
			h = 23
		} else {
			h -= 1
		}
	} else {
		m -= 45
	}

	fmt.Print(h, m)
}
