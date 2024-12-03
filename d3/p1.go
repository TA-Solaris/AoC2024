package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "regexp"
)

type Input struct {
  Commands []Command
}

type Command struct {
  Num1 int
  Num2 int
}

func parseInput(strInput string) Input {
  pattern := `mul\((\d+),(\d+)\)`
  re := regexp.MustCompile(pattern)
  matches := re.FindAllStringSubmatch(strInput, -1)

  input := Input{
    Commands: make([]Command, len(matches)),
  }

  for i, match := range matches {
    num1, err1 := strconv.Atoi(match[1])
    num2, err2 := strconv.Atoi(match[2])

    if err1 != nil || err2 != nil {
      // Should not happen
      fmt.Printf("Error parsing numbers: %v, %v\n", err1, err2)
      continue
    }

    input.Commands[i] = Command{
      Num1: num1,
      Num2: num2,
    }
  }

  return input
}

func calculate(input Input) int {
  total := 0

  for _, command := range input.Commands {
    total += command.Num1 * command.Num2
  }

  return total
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
