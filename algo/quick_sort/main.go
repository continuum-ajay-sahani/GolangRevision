package main

import (
	"fmt"
)

func quickSort(elements []int) []int {
	if len(elements) < 2 {
		return elements
	}
	pivot := elements[0]
	var lesser = []int{}
	var greater = []int{}

	for _, v := range elements[1:] {
		if pivot > v {
			lesser = append(lesser, v)
		} else {
			greater = append(greater, v)
		}
	}
	lesser = append(lesser, pivot)
	lesser = quickSort(lesser)
	greater = quickSort(greater)
	return append(lesser, greater...)
}

func main() {
	e := []int{23, 67, 5, 4, 9, 1}
	se := quickSort(e)
	fmt.Println(se)
}
