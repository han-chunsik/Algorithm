package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int

	fmt.Scan(&a)

	repeated := strings.Repeat("long ", a/4)
	fmt.Print(repeated + "int")
}
