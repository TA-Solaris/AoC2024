package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

type Input struct {
  Grid [][]rune
}

func parseInput(strInput string) Input {
  lines := strings.Split(strings.TrimSpace(strInput), "\n")

  input := Input{
    Grid: make([][]rune, len(lines)),
  }

  for i, line := range lines {
    input.Grid[i] = make([]rune, len(line))
		for j, char := range line {
			input.Grid[i][j] = char
		}
  }

  return input
}

func calculate(input Input) int {
  total := 0

  for i, line := range input.Grid {
		for j, _ := range line {
			if isX(input, i, j) {
        total++
      }
		}
  }

  return total
}

func isX(input Input, row int, col int) bool {
  if !isRune(input, row, col, 'A') {
    return false
  }

  topLeft := isRune(input, row-1, col-1, 'M') && isRune(input, row+1, col+1, 'S')
  topRight := isRune(input, row-1, col+1, 'M') && isRune(input, row+1, col-1, 'S')
  bottomLeft := isRune(input, row+1, col-1, 'M') && isRune(input, row-1, col+1, 'S')
  bottomRight := isRune(input, row+1, col+1, 'M') && isRune(input, row-1, col-1, 'S')

  return (topLeft || bottomRight) && (topRight || bottomLeft)
}

func isRune(input Input, row int, col int, value rune) bool {
  if row < 0 || row >= len(input.Grid) || col < 0 || col >= len(input.Grid[0]) {
    return false
  }

  return input.Grid[row][col] == value
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
