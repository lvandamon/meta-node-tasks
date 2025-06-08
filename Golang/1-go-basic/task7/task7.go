package main

import "fmt"

// 思路
// 空数组，长度返回为 0
// 非空数组，使用快慢指针，slow指针记录不重复的元素位置
func removeDuplicates(nums []int) int {
	// 处理空数组的情况
	if len(nums) == 0 {
		return 0
	}

	// slow指针记录最后一个不重复元素的位置
	slow := 0

	// fast指针遍历整个数组
	for fast := 1; fast < len(nums); fast++ {
		// 当发现新的不重复元素时
		if nums[slow] != nums[fast] {
			// 移动slow指针
			slow++
			// 将不重复元素复制到下一个位置
			nums[slow] = nums[fast]
		}
		// 如果元素相同，fast继续后移直到找到下一个不重复元素
	}

	// 返回不重复元素的个数(长度)
	return slow + 1
}
func main() {
	fmt.Println(removeDuplicates([]int{}))
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{1, 1}))
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{1, 1, 1, 1, 1, 1, 1, 2, 3, 3, 4, 4, 4}))
}
