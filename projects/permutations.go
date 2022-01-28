package main

import (
	"fmt"
)

func getIndex(array []int, item int) int {
	for i, s := range array {
		if s == item {
			return i
		}
	}
	return len(array)
}

func remove(array *[]int, item int) {
	index := getIndex(*array, item)
	*array = append((*array)[:index], (*array)[index+1:]...)
}

func getPermutations(permutationsArray *[][]int, array []int, currentPermutation []int) {
	if len(currentPermutation) == len(array) {
		*permutationsArray = append(*permutationsArray, currentPermutation)
		return
	}
	temp := make([]int, len(array))
	copy(temp, array)
	for _, s := range currentPermutation {
		remove(&temp, s)
	}

	for _, s := range temp {
		getPermutations(permutationsArray, array, append(currentPermutation, s))
	}
}

func permutations(array ...int) [][]int {
	permutationsArray := [][]int{}
	getPermutations(&permutationsArray, array, []int{})
	return permutationsArray
}

func main() {
	fmt.Println(permutations(1, 2, 3, 4,5,6))
}
