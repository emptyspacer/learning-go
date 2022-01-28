package main

import (
  "fmt"
)

func main() {
  var a int
  var b float64
  var c string
  fmt.Print("input> ")
  fmt.Scanf("%d %f %s",&a,&b,&c)
  fmt.Printf("%T, %v\n",a,a)
  fmt.Printf("%T, %v\n",b,b)
  fmt.Printf("%T, %v",c,c)

}