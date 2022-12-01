// print hello world
package main

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func readFile(filename string) {
	text := utils.ReadFile(filename)
	elves := strings.Split(text, "\n\n")
	// reduce elves to the sum of their values
	caloriesPerElf := make([]int, len(elves))
	for i, elf := range elves {
		calories := strings.Split(elf, "\n")
		total := 0
		for _, calorie := range calories {
			// cast string to int
			marks, err := strconv.ParseInt(calorie, 10, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			total += int(marks)
		}
		caloriesPerElf[i] = total
	}
	// sort the elves descending
	sort.Ints(caloriesPerElf)

	fmt.Println("The elf with the most calories is:", caloriesPerElf[len(caloriesPerElf)-1])
	fmt.Println("The sum of the elves with the most calories is:", utils.SumInts(caloriesPerElf[len(caloriesPerElf)-3:]))
}

func main() {
	readFile("input_one.txt")
}
