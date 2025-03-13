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

	var n int
	fmt.Fscanf(r, "%d", &n)

	cards := make([]int, n)
	for i := 0; i < n; i++ {
		cards[i] = i + 1
	}

	for len(cards) > 1 {
		cards = cards[1:]

		first := cards[0]
		cards = cards[1:]
		cards = append(cards, first)
	}

	fmt.Fprintln(w, cards[0])
}
