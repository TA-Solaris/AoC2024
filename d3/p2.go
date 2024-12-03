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
  Type string
  Num1 int
  Num2 int
}

func parseInput(strInput string) Input {
  pattern := `(mul)\((\d+),(\d+)\)|(do)\(\)|(don't)\(\)`
  re := regexp.MustCompile(pattern)
  matches := re.FindAllStringSubmatch(strInput, -1)

  input := Input{
    Commands: make([]Command, len(matches)),
  }

  for i, match := range matches {
    if match[1] == "mul" {
      num1, err1 := strconv.Atoi(match[2])
      num2, err2 := strconv.Atoi(match[3])

      if err1 != nil || err2 != nil {
        // Should not happen
        fmt.Printf("Error parsing numbers: %v, %v\n", err1, err2)
        continue
      }

      input.Commands[i] = Command{
        Type: "mul",
        Num1: num1,
        Num2: num2,
      }
    }
    if match[4] == "do" {
      input.Commands[i] = Command{
        Type: "do",
      }
    }
    if match[5] == "don't" {
      input.Commands[i] = Command{
        Type: "don't",
      }
    }
  }

  return input
}

func calculate(input Input) int {
  total := 0
  enabled := true

  for _, command := range input.Commands {
    if command.Type == "mul" && enabled {
      total += command.Num1 * command.Num2
    }
    if command.Type == "do" {
      enabled = true
    }
    if command.Type == "don't" {
      enabled = false
    }
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
