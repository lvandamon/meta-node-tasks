package main

import (
	"fmt"
)

// 思路
// 字符串数组长度  =0, 返回空
// 字符串数组长度  =1, 返回本身
// 字符串数组长度 >=2, 以第一个字符串为基准，遍历之后的字符串
func longestCommonPrefix(strs []string) string {
	result := ""
	if len(strs) == 0 {
		return result
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// 以第一个字符串为基准，逐个字符比较
	for i := 0; i < len(strs[0]); i++ {
		j := 1 // 从第二个字符串开始比较
		// 检查所有字符串在i位置的字符是否相同
		for j < len(strs) && i < len(strs[j]) && strs[0][i] == strs[j][i] {
			j++
		}
		// 如果所有字符串在i位置都有相同字符
		if j == len(strs) {
			result += string(strs[0][i])
		} else {
			break
		}
	}
	return result
}

func main() {
	s1 := []string{"flower", "flow", "flight"}
	s2 := []string{"dog1", "dog2", "dog3", "dog4"}
	s3 := []string{"da1", "db2", "dc3", "dd4"}
	fmt.Println("s1===>", longestCommonPrefix(s1))
	fmt.Println("s2===>", longestCommonPrefix(s2))
	fmt.Println("s3===>", longestCommonPrefix(s3))
}
