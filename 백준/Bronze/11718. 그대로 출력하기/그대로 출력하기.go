package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		text := s.Text()
		fmt.Println(text)

		if s.Err() == io.EOF {
			break
		}
	}

}
