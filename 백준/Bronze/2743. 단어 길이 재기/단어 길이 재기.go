package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var s string

	fmt.Fscanln(r, &s)
	fmt.Fprint(w, len(s))
	w.Flush()
}
