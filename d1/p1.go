package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
  "sort"
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

  sort.Ints(input.List1)
  sort.Ints(input.List2)

  total := 0

  for i, _ := range input.List1 {
    total += diff(input.List1[i], input.List2[i])
  }

  return total
}

func diff(a int, b int) int {
  if a < b {
    return b - a
  }
  return a - b
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
