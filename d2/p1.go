package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

type Input struct {
  Reports [][]int
}

func parseInput(strInput string) Input {
  lines := strings.Split(strings.TrimSpace(strInput), "\n")

  input := Input{
    Reports: make([][]int, len(lines)),
  }

  for y, line := range lines {
    parts := strings.Split(line, " ")
    input.Reports[y] = make([]int, len(parts))
    for x, part := range parts {
      input.Reports[y][x], _ = strconv.Atoi(part)
    }
  }

  return input
}

func calculate(input Input) int {

  total := 0

  for _, report := range input.Reports {
    if isSafe(report) {
      total++
    }
  }

  return total
}

func isSafe(report []int) bool {
  if len(report) < 2 {
    return true
  }

  carry := report[0]
  isIncreasing := report[0] < report[1]

  for i := 1; i < len(report); i++ {
    num := report[i]
    diff := num - carry

    if isIncreasing && (diff < 1 || diff > 3) {
      return false
    }
    if !isIncreasing && (diff > -1 || diff < -3) {
      return false
    }

    carry = num
  }

  return true
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
