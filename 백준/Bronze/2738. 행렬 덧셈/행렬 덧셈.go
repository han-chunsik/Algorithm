package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readMatrix(r *bufio.Reader, n, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		strInput, _ := r.ReadString('\n')
		row := strings.Split(strings.TrimSpace(strInput), " ")
		matrix[i] = make([]int, m)
		for j, rowStr := range row {
			val, _ := strconv.Atoi(rowStr)
			matrix[i][j] = val
		}
	}
	return matrix
}

func addMatrices(a, b [][]int) [][]int {
	n := len(a)
	m := len(a[0])
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}

func printMatrix(w *bufio.Writer, matrix [][]int) {
	for _, row := range matrix {
		for _, value := range row {
			fmt.Fprintf(w, "%d ", value)
		}
		fmt.Fprintln(w)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscanln(r, &n, &m)

	mtxA := readMatrix(r, n, m)
	mtxB := readMatrix(r, n, m)

	mtxR := addMatrices(mtxA, mtxB)

	printMatrix(w, mtxR)

	w.Flush()
}
