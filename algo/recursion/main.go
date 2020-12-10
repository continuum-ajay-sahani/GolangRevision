package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func countDown(n int) {
	if n < 1 {
		return
	}
	fmt.Println(n)
	countDown(n - 1)
}

func sumMe(sum int, index int, elements []int) int {
	if index > len(elements)-1 {
		return sum
	}
	sum += elements[index]
	index++
	return sumMe(sum, index, elements)
}

func main() {
	//fmt.Println(factorial(5))
	//countDown(5)
	e := []int{10, 6, 20, 4}
	fmt.Println(sumMe(0, 0, e))
}
