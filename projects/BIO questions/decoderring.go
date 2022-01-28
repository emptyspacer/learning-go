package main

import (
  "fmt"
  "strings"
  "time"
)

func translate(pos int, ring []string) int {
  newPos := pos
  ringLength := len(ring)
  for newPos >= ringLength {
    newPos -= ringLength
  }
  return newPos
}

func getIndex(arr []string, v string) int {
  for x, s := range arr {
    if s == v {
      return x
    }
  }
  return 0
}

func main() {
  firstRing := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
  tempFirstRing := make([]string, len(firstRing))
  copy(tempFirstRing, firstRing)

  secondRing := []string{}

  fmt.Print("decoder-ring> ")

  var wordString string

  var n int

  fmt.Scanf("%d %s", &n, &wordString)

  wordString = strings.ToUpper(wordString)

  start := time.Now()

  word := strings.Split(wordString, "")

  currentPos := 0

  for len(tempFirstRing) > 0 {
    currentPos += (n - 1)
    currentPos = translate(currentPos, tempFirstRing)
    secondRing = append(secondRing, tempFirstRing[currentPos])
    tempFirstRing = append(tempFirstRing[:currentPos], tempFirstRing[currentPos+1:]...)

  }

  fmt.Println(strings.Join(secondRing[:6], ""))

  encrypted := []string{}

  for i := range word {
    encrypted = append(encrypted, secondRing[getIndex(firstRing, word[i])])
    secondRing = append(secondRing[1:], secondRing[0])

  }

  fmt.Println(strings.Join(encrypted, ""))

  elapsed := time.Since(start)

  fmt.Printf("program finished in %s", elapsed)

}
