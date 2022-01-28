package main

import (
  "fmt"
)

func example1() {
  fmt.Println("start")
  defer fmt.Println("middle") // deferring a function call sends it to the end of the function it is contained within (it is executed when the function ends)
  fmt.Println("end")
}

func example2() {
  defer fmt.Println("start")
  defer fmt.Println("middle")
  defer fmt.Println("end")

  // deferred functions are executed in reverse order (the last deferred function will be executed first)
}

func example3() {
  a := "start"
  defer fmt.Println(a)
  a = "end"

  // this deferred function will print start even though the value of a has changed, as any function props are locked in when the function is deferred (would assume this has something to do with the fact that any arguments passed to a function are actually copies of those arguments)
}

func main() {
  example1()
  fmt.Println()
  example2()
  fmt.Println()
  example3()
}