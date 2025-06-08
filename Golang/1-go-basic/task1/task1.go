package main

import "fmt"

// 思路
// 交换律：a ^ b ^ c <=> a ^ c ^ b
// 任何数于0异或为任何数 0 ^ n => n
// 相同的数异或为0: n ^ n => 0
func singleNumber(nums []int) int {
	var single int
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	var nums1 = [...]int{1, 1, 2, 2, 3, 3, 4, 4, 5}
	var nums2 = [...]int{1, 1, 2, 13, 4, 2, 3, 3, 4}
	fmt.Println(singleNumber(nums1[:]))
	fmt.Println(singleNumber(nums2[:]))
}
