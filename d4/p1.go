package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

type Input struct {
  Grid [][]rune
}

var directions = [][]int{
	{0, 1},   // right
	{1, 0},   // down
	{0, -1},  // left
	{-1, 0},  // up
	{1, 1},   // diagonal down-right
	{-1, -1}, // diagonal up-left
	{-1, 1},  // diagonal up-right
	{1, -1},  // diagonal down-left
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
			for _, dir := range directions {
        if (findWord(input, i, j, dir[1], dir[0], []rune("XMAS"), 0)) {
          total++
        }
      }
		}
  }

  return total
}

func findWord(input Input, row int, col int, xdir int, ydir int, word []rune, index int) bool {
  if row < 0 || row >= len(input.Grid) || col < 0 || col >= len(input.Grid[0]) {
    return false
  }

  if input.Grid[row][col] != word[index] {
    return false
  }

  if index == len(word)-1 {
    return true
  }

  return findWord(input, row+ydir, col+xdir, xdir, ydir, word, index+1)
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
