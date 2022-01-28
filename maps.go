package main

import (
	"fmt"
)

func main() {
	// go maps are essentially identical to python dicts, except that the keys and values must be of set type
	// the declaration syntax is seen below, but to break it down it is map[keyType]valueType{*contents*}

	example := map[string]int{
		"biggie": 1,
		"cheese": 2, // at the end of maps and structs the last line or variable requires a comma after it (ONLY IF THE MAP OR STRUCT IS BEING DEFINED OVER MULTIPLE LINES)
	}

	fmt.Println(example["biggie"])

	delete(example, "biggie")

	fmt.Println(example["biggie"])
  // keys that are not in the function will return 0 (or the 0 value of the value type)

  _, exists := example["biggie"]
  // using two variables means the second variable is a bool that determines whether the key exists or not

  fmt.Println(exists)

	example["chungus"] = 69

	fmt.Println(example["chungus"])
}
