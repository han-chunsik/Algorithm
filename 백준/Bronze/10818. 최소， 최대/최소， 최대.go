package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a int

	fmt.Fscanln(reader, &a)

	inputNums, _ := reader.ReadString('\n')

	strSlice := strings.Split(strings.TrimSpace(inputNums), " ")

	maxStr := strSlice[0]
	max, _ := strconv.Atoi(maxStr)

	minStr := strSlice[0]
	min, _ := strconv.Atoi(minStr)

	for _, s := range strSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if min > num {
			min = num
		}

		if max < num {
			max = num
		}
	}

	fmt.Fprintln(writer, min, max)

	writer.Flush()
}
