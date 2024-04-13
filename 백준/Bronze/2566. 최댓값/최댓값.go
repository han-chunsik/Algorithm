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

	mtxA := make([][]int, 9)

	for i := 0; i < 9; i++ {
		mtxA[i] = make([]int, 9)
	}

	for i := 0; i < 9; i++ {
		strInput, _ := r.ReadString('\n')
		row := strings.Split(strings.TrimSpace(strInput), " ")
		for j, rowStr := range row {
			val, _ := strconv.Atoi(rowStr)
			mtxA[i][j] = val
		}
	}

	maxLine := 0
	maxNum := 0
	maxIndex := 0

	for i, m := range mtxA {
		for j, s := range m {
			if maxNum < s {
				maxNum = s
				maxLine = i
				maxIndex = j
			}
		}
	}

	fmt.Fprintln(w, maxNum)
	fmt.Fprintln(w, maxLine+1, maxIndex+1)

	w.Flush()
}
