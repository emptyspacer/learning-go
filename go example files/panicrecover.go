package main

import (
  "fmt"
)

func main() {
  // panic is similar to python except (to raise an error message). there are inbuilt panic messages, such as when you try and divide 1 by 0, but if you are building an application you might want to create your own error messages

  fmt.Println("start")
  panic("something happened")
  fmt.Println("end") // this will not print as the application will panic

  //panics happen after deferred functions

  /*
    defer fmt.Println("eee")
    panic("ooo")
    
    eee will be printed

    RECOVERY:
    start from https://youtu.be/YS4e4q9oBaU?t=14204

  */


}