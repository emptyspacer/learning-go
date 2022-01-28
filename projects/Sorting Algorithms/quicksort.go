package main

import "fmt"

func swap(array []int, i1 int, i2 int) {
	array[i1],array[i2] = array[i2],array[i1]
}

func partition(array []int, start int, end int) int {
	pivot := array[start]
	pivIndex := start
	for start < end {
		for start < len(array) && array[start] <= pivot {
			start++
		}
		for array[end] > pivot {
			end--
		}
		if start < end {
			swap(array,start,end)
		}
	}
 	swap(array,pivIndex,end)
 	return end
}

func quickSort(array []int, start int, end int) {
	if start < end {
		pivot := partition(array,start,end)
		quickSort(array,start,pivot-1)
		quickSort(array,pivot+1,end)
	}
}

func main() {
	unsorted := []int{1,6,4,3,7,2,5}
	quickSort(unsorted,0,len(unsorted)-1)
	fmt.Println(unsorted)
}
