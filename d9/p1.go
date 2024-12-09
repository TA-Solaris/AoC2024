package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

type File struct {
  Index int
  Length int
  Free bool
}

type Input struct {
  Files []File
}

func parseInput(strInput string) Input {
  parts := strings.Split(strings.TrimSpace(strInput), "")

  input := Input{
    Files: make([]File, len(parts)),
  }

  for i, part := range parts {
    num, err := strconv.Atoi(part)
    if err != nil {
      panic(err)
    }
    input.Files[i] = File{
      Index: i/2,
      Length: num,
      Free: i%2 != 0,
    }
  }

  return input
}

func calculate(input Input) int {
  reduced := reduceFiles(input.Files, 0)
  //fmt.Println(reduced)
  return sumFiles(reduced)
}

func reduceFiles(files []File, i int) []File {
  printFiles(files)
  if i >= len(files) {
    return files
  }
  if files[i].Free {
    li := len(files)-1
    if files[li].Free {
      files = removeFile(files, li)
      return reduceFiles(files, i)
    } else if files[i].Length > files[li].Length {
      files[i].Length -= files[li].Length
      temp := files[li]
      files = removeFile(files, li)
      files = addFile(files, i, temp)
      return reduceFiles(files, i+1)
    } else if files[i].Length == files[li].Length {
      files[i].Index = files[li].Index
      files[i].Free = false
      files = removeFile(files, li)
      return reduceFiles(files, i+1)
    } else {
      files[li].Length -= files[i].Length
      files[i].Index = files[li].Index
      files[i].Free = false
      return reduceFiles(files, i+1)
    }
  }
  return reduceFiles(files, i+1)
}

func removeFile(files []File, index int) []File {
  if index < 0 || index >= len(files) {
    panic(nil)
  }
  return append(files[:index], files[index+1:]...)
}

func addFile(files []File, index int, value File) []File {
  if index < 0 || index > len(files) {
    panic(nil)
  }
  return append(files[:index], append([]File{value}, files[index:]...)...)
}

func sumFiles(files []File) int {
  total := 0

  for i := 0; 0 < len(files); i++ {
    total += i * files[0].Index
    files[0].Length--
    if files[0].Length <= 0 {
      files = removeFile(files, 0)
    }
  }

  return total
}

func printFiles(files []File) {
  for _, file := range files {
    if file.Length == 0 {
      fmt.Print("#") // Ghosts 👻
    }
    for i := 0; i < file.Length; i++ {
      if file.Free {
        fmt.Print(".")
      } else {
        fmt.Print(file.Index)
      }
    }
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
