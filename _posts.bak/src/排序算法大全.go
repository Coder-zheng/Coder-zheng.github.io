package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand_numbs := generateRandomNumber(0, 30000, 10000)
	items_bubble := rand_numbs
	items_bubble2 := rand_numbs

	items_selection := rand_numbs
	items_selection2 := rand_numbs

	items_insertion := rand_numbs
	items_insertion2 := rand_numbs

	items_quick := rand_numbs
	items_quick2 := rand_numbs

	items_shell := rand_numbs

	items_comb := rand_numbs

	items_merge := rand_numbs

	start := time.Now()
	bubbleSort(items_bubble)
	elapsed := time.Since(start)
	fmt.Printf("Cost time bubbleSort %v\n", elapsed)

	start = time.Now()
	bubbleSortUsingSortPackage(sort.IntSlice(items_bubble2))
	end := time.Now()
	fmt.Printf("Cost time bubbleSortUsingSortPackage %v\n", end.Sub(start))

	start = time.Now()
	selectionSort(items_selection)
	end = time.Now()
	fmt.Printf("Cost time selectionSort %v\n", end.Sub(start))

	start = time.Now()
	selectionSortUsingSortPackage(sort.IntSlice(items_selection2))
	end = time.Now()
	fmt.Printf("Cost time selectionSortUsingSortPackage %v\n", end.Sub(start))

	start = time.Now()
	insertionSort(items_insertion)
	end = time.Now()
	fmt.Printf("Cost time insertionSort %v\n", end.Sub(start))

	start = time.Now()
	insertionSortUsingSortPackage(sort.IntSlice(items_insertion2))
	end = time.Now()
	fmt.Printf("Cost time insertionSortUsingSortPackage %v\n", end.Sub(start))

	start = time.Now()
	QuickSort(items_quick, 0, len(items_quick)-1)
	end = time.Now()
	fmt.Printf("Cost time QuickSort %v\n", end.Sub(start))

	start = time.Now()
	sort.Ints(items_quick2)
	end = time.Now()
	fmt.Printf("Cost time sort.Ints %v\n", end.Sub(start))

	start = time.Now()
	shellshort(sort.IntSlice(items_shell))
	end = time.Now()
	fmt.Printf("Cost time shellshort %v\n", end.Sub(start))

	start = time.Now()
	combsort(sort.IntSlice(items_comb))
	end = time.Now()
	fmt.Printf("Cost time combsort %v\n", end.Sub(start))

	start = time.Now()
	mergeSort(sort.IntSlice(items_merge))
	end = time.Now()
	fmt.Printf("Cost time mergeSort %v\n", end.Sub(start))

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
		n = n - 1
	}
}

func bubbleSortUsingSortPackage(data sort.Interface) {
	r := data.Len() - 1
	for i := 0; i < r; i++ {
		for j := r; j > i; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}

func selectionSort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

func selectionSortUsingSortPackage(data sort.Interface) {
	r := data.Len() - 1
	for i := 0; i < r; i++ {
		min := i
		for j := i + 1; j <= r; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}

func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}

func insertionSortUsingSortPackage(data sort.Interface) {
	r := data.Len() - 1
	for i := 1; i <= r; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func QuickSort(src []int, first, last int) {
	flag := first
	left := first
	right := last

	if first >= last {
		return
	}

	for first < last {
		for first < last {
			if src[last] >= src[flag] {
				last -= 1
				continue
			} else {
				tmp := src[last]
				src[last] = src[flag]
				src[flag] = tmp
				flag = last
				break
			}
		}

		for first < last {
			if src[first] <= src[flag] {
				first += 1
				continue
			} else {
				tmp := src[first]
				src[first] = src[flag]
				src[flag] = tmp
				flag = first
				break
			}
		}
	}

	QuickSort(src, left, flag-1)
	QuickSort(src, flag+1, right)
}

func shellshort(items []int) {
	var (
		n    = len(items)
		gaps = []int{1}
		k    = 1
	)

	for {
		gap := pow(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}

	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if items[j-gap] > items[j] {
					items[j-gap], items[j] = items[j], items[j-gap]
				}
				j = j - gap
			}
		}
	}
}

func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func combsort(items []int) {
	var (
		n       = len(items)
		gap     = len(items)
		shrink  = 1.3
		swapped = true
	)

	for swapped {
		swapped = false
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}
		for i := 0; i+gap < n; i++ {
			if items[i] > items[i+gap] {
				items[i+gap], items[i] = items[i], items[i+gap]
				swapped = true
			}
		}
	}
}

func mergeSort(items []int) []int {
	var n = len(items)

	if n == 1 {
		return items
	}

	middle := int(n / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, n-middle)
	)
	for i := 0; i < n; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// Either left or right may have elements left; consume them.
	// (Only one of the following loops will actually be entered.)
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

func generateRandomNumber(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}

	nums := make([]int, 0)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		num := r.Intn((end - start)) + start

		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	var a int = 1
	fmt.Println("a:", a)
	return nums
}
