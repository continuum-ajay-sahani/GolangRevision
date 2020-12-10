package main

import "fmt"

func selectionSort(elements []int) []int {

	for i, v := range elements {
		tmpi := i
		for j := i; j < len(elements); j++ {
			if elements[j] > v {
				tmpi = j
			}
		}
		if i != tmpi {
			tmp := elements[tmpi]
			elements[tmpi] = elements[i]
			elements[i] = tmp
		}

	}
	return elements
}

func main() {
	e := []int{23, 67, 5, 4, 9, 1}
	se := selectionSort(e)
	fmt.Println(se)
}
