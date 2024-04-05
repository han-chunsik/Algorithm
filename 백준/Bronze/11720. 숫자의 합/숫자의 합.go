package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(r, &n)

	s, _ := r.ReadString('\n')
	strSlice := strings.Split(strings.TrimSpace(s), "")
	intSlice := make([]int, 0, n)

	for _, str := range strSlice {
		num, _ := strconv.Atoi(str)
		intSlice = append(intSlice, num)
	}

	result := 0

	for _, it := range intSlice {
		result += it
	}
	fmt.Fprintln(w, result)

	w.Flush()
}