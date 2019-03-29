package main
import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	arr:=generateRandomNumber(1,100,10)
	fmt.Println("arr:",arr)
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

    return nums
}