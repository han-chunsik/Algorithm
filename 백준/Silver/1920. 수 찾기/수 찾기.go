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

	var n, m int
	fmt.Fscan(r, &n)

	var nNum int
	a := make(map[int]struct{})
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &nNum)
		a[nNum] = struct{}{}
	}

	fmt.Fscan(r, &m)
	var mNum int

	for j := 0; j < m; j++ {
		fmt.Fscan(r, &mNum)
		if _, exists := a[mNum]; exists {
			fmt.Fprintln(w, "1")
		} else {
			fmt.Fprintln(w, "0")
		}

	}

}
