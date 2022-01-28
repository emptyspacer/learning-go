package main

import (
  "fmt"
)

// passing pointers and returning addresses -> https://www.youtube.com/watch?v=YS4e4q9oBaU&t=13639s

func addTo(array *[]int) {
  *array = append(*array,4,5,6)
}

func main() {

  array := []int{1,2,3}
  fmt.Println(array)

  addTo(&array)
  fmt.Println(array)

  /*
  This is an example of the usefullness of pointers
  
  By default function props are actually value copies of the variables that are passed into them

  But by using pointers you can change variables from within a function and have them correspond to the version of the variable that exists outside of the function

  When working with large and hard to handle variables, using pointers can also hugely help with performance, as instead of having to copy that variable, you are simply pointing to where it is stored on the computers memory
  */

}