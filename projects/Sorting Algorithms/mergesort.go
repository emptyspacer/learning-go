package main

import (
	"fmt"
)

func combineSortedLists(l1 []int, l2 []int ) []int {
	l1index := 0
	l2index := 0

	combined := []int{}

	for l1index < len(l1) && l2index < len(l2) {
		if l1[l1index] < l2[l2index] {
			combined = append(combined,l1[l1index])
			l1index++
		} else {
			combined = append(combined,l2[l2index])
			l2index++
		}
	}

	for ; l1index < len(l1); l1index++ {
		combined = append(combined,l1[l1index])
	}

	for ; l2index < len(l2); l2index++ {
		combined = append(combined,l2[l2index])
	}

	return combined

}

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	middle := len(array) / 2

	var right []int = array[middle:]
	var left []int = array[:middle]

	sortedRight := mergeSort(right)
	sortedLeft := mergeSort(left)

	return combineSortedLists(sortedLeft,sortedRight)

}

func main() {
	unsorted := []int{1,6,4,3,7,2,5}
	fmt.Println(mergeSort(unsorted))
}