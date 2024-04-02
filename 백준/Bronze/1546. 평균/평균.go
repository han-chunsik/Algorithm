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

	var s int

	fmt.Fscanln(r, &s)

	scores, _ := r.ReadString('\n')

	strSlice := strings.Split(strings.TrimSpace(scores), " ")
	intSlice := make([]float64, 0, s)

	for _, sts := range strSlice {
		num, err := strconv.ParseFloat(sts, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		intSlice = append(intSlice, num)
	}

	max := intSlice[0]

	for _, num := range intSlice {
		if max < num {
			max = num
		}
	}

	var result float64

	for _, num := range intSlice {
		result += num / max * 100
	}

	fmt.Fprintln(w, result/float64(s))

	w.Flush()
}
