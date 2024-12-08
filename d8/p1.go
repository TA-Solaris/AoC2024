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
  return len(getAntiNodesList(input))
}

func getAntiNodes(input Input) map[rune][]Position {
  antiNodes := make(map[rune][]Position)

  for frequency, positions := range input.Antennas {
    visited := make(map[Position]bool)

    for i := 0; i < len(positions); i++ {
      for j := i + 1; j < len(positions); j++ {
        pos1, pos2 := positions[i], positions[j]
        
        rowDiff := pos2.Row - pos1.Row
        colDiff := pos2.Col - pos1.Col

        if rowDiff == 0 && colDiff == 0 {
          continue
        }
        
        candidates := []Position{
          {Row: pos1.Row - rowDiff, Col: pos1.Col - colDiff},
          {Row: pos2.Row + rowDiff, Col: pos2.Col + colDiff},
        }

        for _, candidate := range candidates {
          if candidate.Row >= 0 && candidate.Row < input.Bounds.Row &&
            candidate.Col >= 0 && candidate.Col < input.Bounds.Col {
            if !visited[candidate] {
              //printAntiNode(input.Bounds, frequency, pos1, pos2, candidate)
              antiNodes[frequency] = append(antiNodes[frequency], candidate)
              visited[candidate] = true
            }
          }
        }
      }
    }
  }

  return antiNodes
}

func getAntiNodesList(input Input) []Position {
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
        
        candidates := []Position{
          {Row: pos1.Row - rowDiff, Col: pos1.Col - colDiff},
          {Row: pos2.Row + rowDiff, Col: pos2.Col + colDiff},
        }

        for _, candidate := range candidates {
          if candidate.Row >= 0 && candidate.Row < input.Bounds.Row &&
            candidate.Col >= 0 && candidate.Col < input.Bounds.Col {
            if !visited[candidate] {
              //printAntiNode(input.Bounds, frequency, pos1, pos2, candidate)
              antiNodes = append(antiNodes, candidate)
              visited[candidate] = true
            }
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
