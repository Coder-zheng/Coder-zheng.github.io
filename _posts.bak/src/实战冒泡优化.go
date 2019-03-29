package main
import "fmt"

func main() {
	arr := []int{1, 2, 4, 5, 3, 6, 9, 7, 8, 0}
	sortBubble(arr)
	fmt.Println("arr:",arr)
}
func sortBubble(arr []int) {
	n := len(arr)-1
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			if arr[j-1]>arr[j] {
				arr[j],arr[j-1] = arr[j-1],arr[j]
			}
		}
	}
}
