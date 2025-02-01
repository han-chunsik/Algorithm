package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	/*
		1 2 3 4 5 6 7 8: ascending
		8 7 6 5 4 3 2 1: descending
		그 외: mixed
	*/

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	numbers := make([]int, 8)

	fmt.Fscanf(r, "%d %d %d %d %d %d %d %d", &numbers[0], &numbers[1], &numbers[2], &numbers[3], &numbers[4], &numbers[5], &numbers[6], &numbers[7])

	ascending := []int{1, 2, 3, 4, 5, 6, 7, 8}
	descending := []int{8, 7, 6, 5, 4, 3, 2, 1}

	switch {
	case slices.Equal(numbers, ascending):
		fmt.Fprintln(w, "ascending")
	case slices.Equal(numbers, descending):
		fmt.Fprintln(w, "descending")
	default:
		fmt.Fprintln(w, "mixed")
	}

	w.Flush()
}
