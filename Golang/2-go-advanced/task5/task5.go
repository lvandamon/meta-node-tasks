package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 打印形状信息的通用函数
func printInfo(s Shape) {
	fmt.Println("面积: ", s.Area())
	fmt.Println("周长: ", s.Perimeter())
}

func main() {
	rectangle := Rectangle{3, 4}
	printInfo(&rectangle)

	circle := Circle{10}
	printInfo(&circle)

}
