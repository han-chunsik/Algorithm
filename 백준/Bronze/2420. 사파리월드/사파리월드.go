package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscanf(r, "%d %d", &n, &m)

	fmt.Fprintln(w, int(math.Abs(float64(n-m))))

	w.Flush()
}
