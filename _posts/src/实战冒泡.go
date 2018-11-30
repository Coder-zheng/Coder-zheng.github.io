package main

import "fmt"

func main() {
	arr := []int{1, 2, 4, 5, 3, 6, 9, 7, 8, 0}
	bubbleSort(arr)
	fmt.Println("arr:", arr)
}
func bubbleSort(items []int) {
	var (
		n       = len(items)
		swapped = true
	)
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		// n = n - 1
	}
}
