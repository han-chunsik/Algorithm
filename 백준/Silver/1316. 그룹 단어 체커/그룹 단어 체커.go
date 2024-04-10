package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(r, &n)

	count := 0
	for i := 0; i < n; i++ {
		s, _ := r.ReadString('\n')
		s = strings.TrimSpace(s)

		lastIndex := 0
		for _, c := range s {
			if lastIndex > strings.IndexRune(s, c) {
				count -= 1
				break
			} else {
				lastIndex = strings.IndexRune(s, c)
			}

		}
		count += 1
	}
	fmt.Fprint(w, count)

	w.Flush()
}
