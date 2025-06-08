package main

import "fmt"

func merge(intervals [][]int) [][]int {
	//TODO
	return [][]int{}
}

func main() {
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals1))

	intervals2 := [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals2))
}
