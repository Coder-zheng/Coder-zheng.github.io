package main
import "fmt"
func main() {
	items_quick:= []int{11,23,4,2,8,6,10,5}
	QuickSort(items_quick, 0, len(items_quick)-1)
	fmt.Println(items_quick)
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
