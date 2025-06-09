package main

import "fmt"

func double(value *[]int) []int {
	var result []int
	for _, v := range *value {
		v *= 2
		result = append(result, v)
	}
	return result
}

func main() {
	value := []int{1, 2, 3, 4, 5}
	fmt.Println(double(&value)) //{2,4,6,8,10}
}
