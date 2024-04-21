import (
    "strconv"
)

func solution(num_str string) int {
	var numInt int
	result := 0
	for i := 0; i < len(num_str); i++ {
		numInt, _ = strconv.Atoi(string(num_str[i]))
		result += numInt
	}
	return result
}