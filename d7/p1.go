package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

type Equation struct {
  Terms []int
  Result int
}

type Input struct {
  Equations []Equation
}

func parseInput(strInput string) Input {
  lines := strings.Split(strings.TrimSpace(strInput), "\n")

  input := Input{
    Equations: make([]Equation, len(lines)),
  }

  for i, line := range lines {
    parts := strings.Split(line, ": ")
    num, err := strconv.Atoi(parts[0])
    if err != nil {
      panic(err)
    }
    input.Equations[i].Result = num

    terms := strings.Split(parts[1], " ")
    input.Equations[i].Terms = make([]int, len(terms))

		for j, term := range terms {
      num, err := strconv.Atoi(term)
      if err != nil {
        panic(err)
      }
      input.Equations[i].Terms[j] = num
		}
  }

  return input
}

func calculate(input Input) int {

  total := 0

  for _, equation := range input.Equations {
    if valid(equation.Result, equation.Terms[0], equation.Terms[1:]) {
      total += equation.Result
    }
  }
  
  return total
}

func valid(result int, current int, terms []int) bool {
  if len(terms) == 0 {
    return result == current
  }
  if valid(result, current+terms[0], terms[1:]) {
    return true
  }
  return valid(result, current*terms[0], terms[1:])
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
