package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始处理
	for i := n - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++

		// 如果加1后小于10，直接返回
		if digits[i] < 10 {
			return digits
		}

		// 否则当前位设为0，进位到前一位
		digits[i] = 0
	}

	// 如果所有位都进位了（例如999变成1000的情况）
	// 需要在最前面添加一个1
	return append([]int{1}, digits...)
}

func main() {
	// 测试用例1
	digits1 := []int{1, 2, 3}
	fmt.Println(plusOne(digits1)) // 输出: [1 2 4]

	// 测试用例2
	digits2 := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits2)) // 输出: [4 3 2 2]

	// 测试用例3
	digits3 := []int{9}
	fmt.Println(plusOne(digits3)) // 输出: [1 0]

	// 测试用例4 - 多个9的情况
	digits4 := []int{9, 9, 9}
	fmt.Println(plusOne(digits4)) // 输出: [1 0 0 0]

	// 测试用例5 - 0的情况
	digits5 := []int{0}
	fmt.Println(plusOne(digits5)) // 输出: [1]
}
