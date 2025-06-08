package main

import "fmt"

// 使用双重循环暴力解法,时间复杂度O(n²)
func twoSum(nums []int, target int) []int {
	// 外层循环遍历数组中的每个元素
	for i := 0; i < len(nums); i++ {
		// 内层循环遍历当前元素之后的所有元素
		for j := i + 1; j < len(nums); j++ {
			// 检查两数之和是否等于目标值
			if nums[i]+nums[j] == target {
				// 找到匹配，返回两个索引
				return []int{i, j}
			}
		}
	}

	// 如果没有找到符合条件的两个数，返回空切片
	return []int{}
}

// 使用哈希表优化查找两数之和,时间复杂度O(n),空间复杂度O(n)
func twoSum1(nums []int, target int) []int {
	// 创建哈希表存储元素值和索引
	numMap := make(map[int]int)

	// 遍历数组
	for i, num := range nums {
		// 计算当前元素所需的补数
		complement := target - num
		// 检查补数是否在哈希表中
		if index, ok := numMap[complement]; ok {
			// 找到匹配，返回两个索引
			return []int{index, i}
		}
		// 将当前元素存入哈希表
		numMap[num] = i
	}

	// 无解返回空切片
	return []int{}
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println(twoSum(nums1, target1))
	fmt.Println(twoSum1(nums1, target1))

	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println(twoSum(nums2, target2))
	fmt.Println(twoSum1(nums2, target2))

	nums3 := []int{3, 3}
	target3 := 6
	fmt.Println(twoSum(nums3, target3))
	fmt.Println(twoSum1(nums3, target3))

	nums4 := []int{2, 7, 11, 15}
	target4 := 17
	fmt.Println(twoSum(nums4, target4))
	fmt.Println(twoSum1(nums4, target4))
}
