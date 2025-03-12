package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, h int
	fmt.Fscanf(r, "%d\n", &n)

	height := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d\n", &h)
		height[i] = h
	}

	max := 0
	result := 0

	for j := n - 1; j >= 0; j-- {
		if height[j] > max {
			max = height[j]
			result++
		}
	}

	fmt.Fprintln(w, result)
}
