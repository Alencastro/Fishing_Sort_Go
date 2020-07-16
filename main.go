package main

import (
	"fmt"
)

var (
	arr = [10]int{3, 1, 5, 2, 8, 6, 7, 10, 9, 4}
)

func sort() {
	for c := 0; c < len(arr)-1; c++ {
		if arr[c] > arr[(c+1)] {
			arr[c], arr[(c+1)] = arr[(c+1)], arr[c]
			c = -1
		}
	}
}
func main() {
	fmt.Println(arr)
	sort()
	fmt.Println(arr)
}