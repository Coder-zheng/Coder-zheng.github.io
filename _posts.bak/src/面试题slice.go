package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"a", "b"}
	modify1(slice1)
	fmt.Println(slice1)
	slice2 := []string{"a", "b"}
	modify2(slice2)
	fmt.Println(slice2)
}
func modify1(data []string) {
	data = nil
}
func modify2(data []string) {
	data[1] = "c"
}
