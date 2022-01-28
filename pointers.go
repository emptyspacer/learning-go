package main

import (
  "fmt"
)

// see https://go.dev/tour/moretypes/1 for more information

func main() {
  var a int = 42;
  var b *int = &a // put asterix before type to specify what type the pointer is pointing to
  fmt.Println(a,b)

  /*
  There are two operators when talking about pointers: the referencing operator (&) and the dereferencing operator (*). 

  By default a pointer such as b will reference the memory location of that variable (as seen in the first print statment)
  */

  //the referencing operator
  fmt.Println(&a) // will be the same value as b

  //the dereferencing operator
  fmt.Println(*b) // will be the same value as a

  //*b will always give the latest value of a
  a = 77
  fmt.Println(a,*b)

}