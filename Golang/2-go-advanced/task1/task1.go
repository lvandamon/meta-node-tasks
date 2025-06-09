package main

import "fmt"

func addTen(num *int) {
	*num += 10
}

func main() {
	num := 5
	addTen(&num)
	fmt.Println("调用函数前的值:", num) // 15
}
