package utils

import (
	"fmt"
	"os"
)

// export sumInts
func SumInts(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	text := string(data)
	return text
}
