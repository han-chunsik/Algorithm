package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a int
	fmt.Fscanln(reader, &a)

	var b, c int

	for i := 0; i < a; i++ {
		fmt.Fscanln(reader, &b, &c)
		fmt.Fprintln(writer, b+c)
	}

	writer.Flush()
}
