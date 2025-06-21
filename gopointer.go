package main

import "fmt"

// 定义一个接收整数指针的函数
func addTen(num *int) {
        *num += 10 // 通过指针修改原始值
}

func main() {
        value := 5
        fmt.Println("修改前的值:", value)

        addTen(&value) // 传递value的地址

        fmt.Println("修改后的值:", value)
}