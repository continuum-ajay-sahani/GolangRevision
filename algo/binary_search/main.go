package main

import "fmt"

func binarySearch(low, high int, elements []int, find int) bool {

	for low <= high {
		mid := low + (high-low)/2
		if elements[mid] == find {
			return true
		} else if elements[mid] > find {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
func main() {
	a := []int{10, 20, 30, 40, 50, 60}
	found := binarySearch(0, len(a)-1, a, 40)
	fmt.Println(found)
	found = binarySearch(0, len(a)-1, a, 44)
	fmt.Println(found)
}
