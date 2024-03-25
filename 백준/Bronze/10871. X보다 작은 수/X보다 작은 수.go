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

	result := ""

	for _, s := range strSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if num < b {
			result += s + " "
		}
	}

	fmt.Fprint(writer, result)

	writer.Flush()
}
