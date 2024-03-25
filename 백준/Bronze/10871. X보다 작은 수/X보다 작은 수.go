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

	var a, b int

	fmt.Fscanln(reader, &a, &b)

	inputNums, _ := reader.ReadString('\n')

	strSlice := strings.Split(strings.TrimSpace(inputNums), " ")
	intSlice := make([]int, 0, a)

	for _, s := range strSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		intSlice = append(intSlice, num)
	}

	result := ""

	for i := 0; i < a; i++ {
		if intSlice[i] < b {
			result = result + strconv.Itoa(intSlice[i]) + " "
		}
	}

	fmt.Fprint(writer, result)

	writer.Flush()
}
