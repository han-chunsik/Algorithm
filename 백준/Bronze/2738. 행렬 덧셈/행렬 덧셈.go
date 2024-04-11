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

	var n, m int

	fmt.Fscanln(r, &n, &m)

	mtxA := make([][]int, n)
	mtxB := make([][]int, n)
	mtxR := make([][]int, n)

	for i := 0; i < n; i++ {
		mtxA[i] = make([]int, m)
		mtxB[i] = make([]int, m)
		mtxR[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		strInput, _ := r.ReadString('\n')
		row := strings.Split(strings.TrimSpace(strInput), " ")
		for j, rowStr := range row {
			val, _ := strconv.Atoi(rowStr)
			mtxA[i][j] = val
		}
	}

	for i := 0; i < n; i++ {
		strInput, _ := r.ReadString('\n')
		row := strings.Split(strings.TrimSpace(strInput), " ")
		for j, rowStr := range row {
			val, _ := strconv.Atoi(rowStr)
			mtxB[i][j] = val
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			mtxR[i][j] = mtxA[i][j] + mtxB[i][j]
		}
	}

	for _, row := range mtxR {
		for _, value := range row {
			fmt.Fprintf(w, "%d ", value)
		}
		fmt.Fprintln(w)
	}

	w.Flush()
}
