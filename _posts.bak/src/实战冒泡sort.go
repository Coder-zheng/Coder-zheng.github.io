package main
import (
	"sort"
	"fmt"
)
type Arr []int
func (arr Arr)Len()int{return len(arr)}
func (arr Arr)Swap(i,j int){arr[i],arr[j] = arr[j],arr[i]}
func (arr Arr)Less(i,j int)bool{return arr[i]<arr[j]}
func main() {
	arr := []int{1, 2, 4, 5, 3, 6, 9, 7, 8, 0}
	bubbleSortUsingSortPackage(Arr(arr))
	fmt.Println("arr:",arr)
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