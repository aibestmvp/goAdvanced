package main

import (
	"fmt"
	"math"
)

// Shape 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// 实现Shape接口的Area方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 实现Shape接口的Perimeter方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 结构体
type Circle struct {
	Radius float64
}

// 实现Shape接口的Area方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 实现Shape接口的Perimeter方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 创建Rectangle实例
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("矩形: 面积=%.2f, 周长=%.2f\n", rect.Area(), rect.Perimeter())

	// 创建Circle实例
	circle := Circle{Radius: 4}
	fmt.Printf("圆形: 面积=%.2f, 周长=%.2f\n", circle.Area(), circle.Perimeter())
}
