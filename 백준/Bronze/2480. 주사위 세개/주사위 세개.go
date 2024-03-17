package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a == b && b == c {
		fmt.Print(10000 + a*1000)
	} else if a == b {
		fmt.Print(1000 + a*100)
	} else if a == c {
		fmt.Print(1000 + a*100)
	} else if b == c {
		fmt.Print(1000 + b*100)
	} else {
		max := a
		if max < b {
			max = b
		}
		if max < c {
			max = c
		}
		fmt.Print(100 * max)
	}
}
