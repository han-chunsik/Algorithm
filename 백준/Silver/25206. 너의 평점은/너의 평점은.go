package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func gToGp(grade string) float64 {
	result := 0.0
	switch grade {
	case "A+":
		result = 4.5
	case "A0":
		result = 4.0
	case "B+":
		result = 3.5
	case "B0":
		result = 3.0
	case "C+":
		result = 2.5
	case "C0":
		result = 2.0
	case "D+":
		result = 1.5
	case "D0":
		result = 1.0
	case "F":
		result = 0.0
	}
	return result
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	totalGp := 0.0
	sumGpTimesG := 0.0

	for i := 0; i < 20; i++ {
		strInput, _ := r.ReadString('\n')
		gradeInfo := strings.Split(strings.TrimSpace(strInput), " ")
		gp, _ := strconv.ParseFloat(gradeInfo[1], 64)
		g := gradeInfo[2]

		if g != "P" {
			sumGpTimesG += (gp * gToGp(g))
			totalGp += gp
		}
	}
	fmt.Fprint(w, sumGpTimesG/totalGp)
	w.Flush()
}
