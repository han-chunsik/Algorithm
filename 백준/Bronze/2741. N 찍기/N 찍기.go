package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		자연수 N이 주어졌을 때, 1부터 N까지 한 줄에 하나씩 출력
	*/

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanf(r, "%d", &n)

	for i := 1; i <= n; i++ {
		fmt.Fprintln(w, i)
	}
	w.Flush()
}
