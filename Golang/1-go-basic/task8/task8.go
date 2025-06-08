package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 如果区间数量小于等于1，直接返回
	if len(intervals) <= 1 {
		return intervals
	}

	// 先按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果集，先把第一个区间加入
	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]

		// 如果当前区间的起始位置 <= 上一个区间的结束位置，说明有重叠
		if current[0] <= last[1] {
			// 合并区间，结束位置取两者中较大的
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 没有重叠，直接加入结果集
			merged = append(merged, current)
		}
	}

	return merged
}

func main() {
	// 测试用例1
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals1)) // 输出: [[1 6] [8 10] [15 18]]

	// 测试用例2
	intervals2 := [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals2)) // 输出: [[1 5]]

	// 测试用例3 - 完全包含的情况
	intervals3 := [][]int{{1, 4}, {2, 3}}
	fmt.Println(merge(intervals3)) // 输出: [[1 4]]

	// 测试用例4 - 空输入
	intervals4 := [][]int{}
	fmt.Println(merge(intervals4)) // 输出: []

	// 测试用例5 - 单个区间
	intervals5 := [][]int{{1, 3}}
	fmt.Println(merge(intervals5)) // 输出: [[1 3]]
}
