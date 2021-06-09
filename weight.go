package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "errors"
  "strconv"
)

var (
  reader = bufio.NewReader(os.Stdin)
  errCannotReadInput = errors.New("cannot read input")
  errInputMustBePositive = errors.New("input must be positive")
)

func main() {
  fmt.Println("Average weight calculator")
  rounds := getRounds()
  if rounds == 0 {
    fmt.Println("No rounds entered. Goodbye.")
    return
  }
  var totalWeight int
  for i := 0; i < rounds; i++ {
    totalWeight += getWeight()
  }
  average := totalWeight / rounds
  fmt.Println("The average weight is", average)
}

func getRounds() int {
  fmt.Print("Enter the number of weights to average: ")
  rounds, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(errCannotReadInput)
  }
  rStr := trimNewline(rounds)
  r := toInt(rStr)
  if r < 0 {
    log.Fatal(errInputMustBePositive)
  }
  return r
}

func getWeight() int {
  fmt.Print("Enter a weight: ")
  weight, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(errCannotReadInput)
  }
  wStr := trimNewline(weight)
  w := toInt(wStr)
  if w < 0 {
    log.Fatal(errInputMustBePositive)
  }
  fmt.Println("A weight of", w, "was registered.")
  return w
}

func trimNewline(s string) string {
  l := len(s)
  t := s[:l-1]
  return t
}

func toInt(s string) int {
  i, err := strconv.Atoi(s)
  if err != nil {
    log.Fatal(errors.New("cannot convert string to int"))
  }
  return i
}
