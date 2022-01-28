package main

import (
  "fmt"
)

type Human struct {
  name   string
  height float64 `meters`
  weight float64 `kilograms`
}

func (h Human) bmi() string {
  heightSquared := h.height*h.height
  return fmt.Sprintf("%s's bmi is %v",h.name,h.weight/heightSquared)
}

func main() {
  matt := Human{
    name: "Matt",
    height: 1.8,
    weight: 75,
  }
  fmt.Println(matt.bmi())
}