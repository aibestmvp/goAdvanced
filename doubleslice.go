package main

import "fmt"

func doubleSlice(slicePtr *[]int) {
	// 解引用指针获取切片
	slice := *slicePtr
	// 遍历切片并修改每个元素
	for i := range slice {
		slice[i] *= 2
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Before:", nums)
	doubleSlice(&nums)
	fmt.Println("After:", nums)
}
