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

func (p Position) Equals(other Position) bool {
  return p.Row == other.Row && p.Col == other.Col
}

type Guard struct {
  Position Position
  Direction Position
}

func (g Guard) Equals(other Guard) bool {
  return g.Position.Equals(other.Position) && g.Direction.Equals(other.Direction)
}

func (g Guard) Copy() Guard {
  return Guard{
    Position:  g.Position,
    Direction: g.Direction,
  }
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
  pathGrid := copyGrid(input.Grid)
  pathGrid, _ = runGrid(input.Guard.Copy(), pathGrid)
  pathGrid[input.Guard.Position.Row][input.Guard.Position.Col] = '.'

  total := 0

  for i, row := range pathGrid {
    for j, char := range row {
      if char == 'X' {
        tempGrid := copyGrid(input.Grid)
        tempGrid[i][j] = 'O'
        _, looped := runGrid(input.Guard.Copy(), tempGrid)
        if looped {
          total++
          //printGrid(tempGrid)
          //fmt.Println()
        }
      }
    }
  }

  return total
}

func runGrid(guard Guard, grid [][]rune) ([][]rune, bool) {
  next := true
  hit := false
  hits := make([]Guard, 0)

  for next {
    last := guard
    guard, grid, next, hit = step(guard, grid)
    if hit {
      for _, past := range hits {
        if last.Equals(past) {
          return grid, true
        }
      }
      hits = append(hits, last)
    }
  }

  return grid, false
}

func step(guard Guard, grid [][]rune) (Guard, [][]rune, bool, bool) {

  nextRow := guard.Position.Row + guard.Direction.Row
  nextCol := guard.Position.Col + guard.Direction.Col

  if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
    return guard, grid, false, false
  }

  if grid[nextRow][nextCol] == '#' || grid[nextRow][nextCol] == 'O' {
    temp := guard.Direction.Row
    guard.Direction.Row = guard.Direction.Col
    guard.Direction.Col = -temp
    return guard, grid, true, true
  }

  grid[nextRow][nextCol] = 'X'
  guard.Position.Row = nextRow
  guard.Position.Col = nextCol
  return guard, grid, true, false
}

func copyGrid(grid [][]rune) [][]rune {
  copiedGrid := make([][]rune, len(grid))
  for i := range grid {
    copiedGrid[i] = make([]rune, len(grid[i]))
    copy(copiedGrid[i], grid[i])
  }
  return copiedGrid
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
