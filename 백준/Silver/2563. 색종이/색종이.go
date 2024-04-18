package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var mtx [100][100]int

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			mtx[i][j] = 0
		}
	}

	var n int
	fmt.Fscanln(r, &n)

	var y, x int
	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &y, &x)

		for j := x; j < x+10; j++ {
			for k := y; k < y+10; k++ {
				mtx[j][k] = 1
			}
		}
	}

	result := 0

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if mtx[i][j] == 1 {
				result += 1
			}
		}
	}

	fmt.Fprint(w, result)
	w.Flush()
}
