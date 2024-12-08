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

type Input struct {
  Antennas map[rune][]Position
  Bounds Position
}

func parseInput(strInput string) Input {
  lines := strings.Split(strings.TrimSpace(strInput), "\n")

  input := Input{
    Antennas: make(map[rune][]Position),
    Bounds: Position{
      Row: len(lines),
      Col: len(lines[0]),
    },
  }

  for i, line := range lines {
    for j, char := range line {
      if char != '.' {
        if input.Antennas[char] == nil {
          input.Antennas[char] = make([]Position, 0)
        }
    
        input.Antennas[char] = append(input.Antennas[char], Position{
          Row: i,
          Col: j,
        })
      }
    }
  }

  return input
}

func calculate(input Input) int {
  return len(getAntiNodes(input))
}

func getLine(bounds Position, pos1 Position, pos2 Position) []Position {
  diff := Position{
    Row: pos2.Row - pos1.Row,
    Col: pos2.Col - pos1.Col,
  }
  start := getStart(bounds, pos1, diff)
  return getPoints(bounds, start, diff, make([]Position, 0))
}

func getPoints(bounds Position, pos Position, diff Position, points []Position) []Position {
  if pos.Col >= 0 && pos.Col < bounds.Col && pos.Row >= 0 && pos.Row < bounds.Row {
    return getPoints(bounds, Position{
      Row: pos.Row + diff.Row,
      Col: pos.Col + diff.Col,
    }, diff, append(points, pos))
  }
  return points
}

func getStart(bounds Position, pos Position, diff Position) Position {
  if pos.Col - diff.Col >= 0 && pos.Col - diff.Col < bounds.Col && pos.Row - diff.Row >= 0 && pos.Row - diff.Row < bounds.Row {
    return getStart(bounds, Position{
      Row: pos.Row - diff.Row,
      Col: pos.Col - diff.Col,
    }, diff)
  }
  return pos
}

func getAntiNodes(input Input) []Position {
  antiNodes := make([]Position, 0)

  visited := make(map[Position]bool)
  for _, positions := range input.Antennas {

    for i := 0; i < len(positions); i++ {
      for j := i + 1; j < len(positions); j++ {
        pos1, pos2 := positions[i], positions[j]
        
        rowDiff := pos2.Row - pos1.Row
        colDiff := pos2.Col - pos1.Col

        if rowDiff == 0 && colDiff == 0 {
          continue
        }
        
        linePoints := getLine(input.Bounds, pos1, pos2)

        for _, point := range linePoints {
          if !visited[point] {
            //printAntiNode(input.Bounds, 'T', pos1, pos2, point)
            antiNodes = append(antiNodes, point)
            visited[point] = true
          }
        }
      }
    }
  }

  return antiNodes
}

func printAntiNode(bounds Position, frequency rune, pos1 Position, pos2 Position, candidate Position) {
  for i := 0; i < bounds.Row; i++ {
    for j := 0; j < bounds.Col; j++ {
      if pos1.Row == i && pos1.Col == j {
        fmt.Print(string(frequency))
      } else if pos2.Row == i && pos2.Col == j {
        fmt.Print(string(frequency))
      } else if candidate.Row == i && candidate.Col == j {
        fmt.Print(string('#'))
      } else {
        fmt.Print(string('.'))
      }
    }
    fmt.Println()
  }
  fmt.Println()
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
