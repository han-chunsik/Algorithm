package main

import (
	"fmt"
	"strings"
)

func main() {
	croatia := []string{"c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="}

	var s string

	fmt.Scan(&s)

	replaced := s

	for _, i := range croatia {
		replaced = strings.Replace(replaced, i, "@", -1)
	}

	fmt.Println(len(replaced))
}
