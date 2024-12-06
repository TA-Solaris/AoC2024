package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

type Position struct {
  Row int
  Col int
}

type Guard struct {
  Position Position
  Direction Position
}

type Input struct {
  Grid [][]rune
  Guard Guard
}

func parseInput(strInput string) Input {
  lines := strings.Split(strings.TrimSpace(strInput), "\n")

  input := Input{
    Grid: make([][]rune, len(lines)),
  }

  for i, line := range lines {
    input.Grid[i] = make([]rune, len(line))
		for j, char := range line {
      if char == '^' {
        input.Guard = Guard{
          Position: Position{
            Row: i,
            Col: j,
          },
          Direction: Position{
            Row: -1,
            Col: 0,
          },
        }
        input.Grid[i][j] = 'X'
      } else {
        input.Grid[i][j] = char
      }
		}
  }

  return input
}

func calculate(input Input) int {
  next := true

  for next {
    input, next = step(input)
  }

  //printGrid(input.Grid)
  return countX(input.Grid)
}

func step(input Input) (Input, bool) {

  nextRow := input.Guard.Position.Row + input.Guard.Direction.Row
  nextCol := input.Guard.Position.Col + input.Guard.Direction.Col

  if nextRow < 0 || nextRow >= len(input.Grid) || nextCol < 0 || nextCol >= len(input.Grid[0]) {
    return input, false
  }

  if input.Grid[nextRow][nextCol] == '#' {
    temp := input.Guard.Direction.Row
    input.Guard.Direction.Row = input.Guard.Direction.Col
    input.Guard.Direction.Col = -temp
    return input, true
  }

  input.Grid[nextRow][nextCol] = 'X'
  input.Guard.Position.Row = nextRow
  input.Guard.Position.Col = nextCol
  return input, true
}

func countX(grid [][]rune) int {
  count := 0
  for _, row := range grid {
    for _, char := range row {
      if char == 'X' {
        count++
      }
    }
  }
  return count
}

func printGrid(grid [][]rune) {
  for _, row := range grid {
    fmt.Println(string(row))
  }
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
