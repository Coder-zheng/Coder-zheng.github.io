package main
import (
	"container/list"
	"fmt"
)
func main() {
	link := list.New()

	// 循环插入到头部
	for i := 0; i <= 10; i++ {
		link.PushBack(i)
	}

	// 遍历链表
	for p := link.Front(); p != link.Back(); p = p.Next() {
		fmt.Println("Number", p.Value)
	}

}
