package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

type Input struct {
  Rules map[int][]int
  Updates [][]int
}

func parseInput(strInput string) Input {
  sections := strings.Split(strings.TrimSpace(strInput), "\n\n")

  strRules := strings.Split(strings.TrimSpace(sections[0]), "\n")
  strUpdates := strings.Split(strings.TrimSpace(sections[1]), "\n")

  input := Input{
    Rules: make(map[int][]int),
    Updates: make([][]int, len(strUpdates)),
  }

  for _, strRule := range strRules {
    parts := strings.Split(strings.TrimSpace(strRule), "|")

    num1, err1 := strconv.Atoi(parts[0])
    num2, err2 := strconv.Atoi(parts[1])

    if err1 != nil || err2 != nil {
      // Should not happen
      fmt.Printf("Error parsing numbers: %v, %v\n", err1, err2)
      continue
    }

    if input.Rules[num1] == nil {
      input.Rules[num1] = make([]int, 0)
    }

    input.Rules[num1] = append(input.Rules[num1], num2)
  }

  for i, strUpdate := range strUpdates {
    parts := strings.Split(strings.TrimSpace(strUpdate), ",")
    input.Updates[i] = make([]int, len(parts))
    for j, part := range parts {
      num, err := strconv.Atoi(part)

      if err != nil {
        panic(err)
      }

      input.Updates[i][j] = num
    }
  }

  return input
}

func calculate(input Input) int {
  total := 0

  for _, update := range getInvalidUpdates(input) {
    if validPerm, ok := findValidPermutation(update, input.Rules); ok {
      total += validPerm[len(validPerm) / 2]
    }
  }

  return total
}

func findValidPermutation(update []int, rules map[int][]int) ([]int, bool) {
  used := make(map[int]bool)
  return backtrack(update, []int{}, used, rules)
}

func backtrack(update []int, current []int, used map[int]bool, rules map[int][]int) ([]int, bool) {
  if len(current) == len(update) {
    if isValid(current, 0, rules) {
      return current, true
    }
    return nil, false
  }

  for i := 0; i < len(update); i++ {
    if used[i] {
      continue
    }
    if len(current) > 0 && !isValidNext(current[len(current)-1], update[i], rules) {
      continue
    }

    used[i] = true
    current = append(current, update[i])

    if result, ok := backtrack(update, current, used, rules); ok {
      return result, true
    }

    used[i] = false
    current = current[:len(current)-1]
  }

  return nil, false
}

func getInvalidUpdates(input Input) [][]int {
  invalidUpdates := make([][]int, 0)

  for _, update := range input.Updates {
    if (!isValid(update, 0, input.Rules)) {
      invalidUpdates = append(invalidUpdates, update)
    }
  }

  return invalidUpdates
}

func isValid(update []int, i int, rules map[int][]int) bool {
  if i == len(update)-1 {
    return true
  }
  if isValidNext(update[i], update[i+1], rules) {
    return isValid(update, i+1, rules)
  }

  return false
}

func isValidNext(current int, next int, rules map[int][]int) bool {
  allowedNext, exists := rules[current]
  if !exists {
    return false
  }

  for _, v := range allowedNext {
    if v == next {
      return true
    }
  }
  return false
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
