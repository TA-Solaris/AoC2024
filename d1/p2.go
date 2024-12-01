package main

import (
  "fmt"
  "io/ioutil"
	"strings"
	"strconv"
)

type Input struct {
	List1 []int
	List2 []int
}

func parseInput(strInput string) Input {
	lines := strings.Split(strings.TrimSpace(strInput), "\n")

	input := Input{
		List1: make([]int, len(lines)),
		List2: make([]int, len(lines)),
	}

	for i, line := range lines {
		parts := strings.Split(line, "   ")
		input.List1[i], _ = strconv.Atoi(parts[0])
		input.List2[i], _ = strconv.Atoi(parts[1])
	}

	return input
}

func calculate(input Input) int {

	counts := countOccurrences(input.List2)

	total := 0

	for _, num := range input.List1 {
		if occurrences, ok := counts[num]; ok {
			total += num * occurrences
		}
	}

	return total
}

func countOccurrences(nums []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	return counts
}

func main() {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    panic(err)
  }

	input := parseInput(string(data))
	output := calculate(input)

  fmt.Println(output)
}
