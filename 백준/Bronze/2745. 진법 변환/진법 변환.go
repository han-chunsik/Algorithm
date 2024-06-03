package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	nums := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var n string
	var b int

	fmt.Fscanf(r, "%s %d", &n, &b)

	sum := 0

	for i, char := range reverseString(n) {
		charIdx := strings.IndexRune(nums, char)
		square := math.Pow(float64(b), float64(i))
		sum += charIdx * int(square)
	}
	fmt.Fprintln(w, sum)

	w.Flush()
}
