package main

import (
	"fmt"
	"strings"
)

// 思路
// 两两对应必定存在() or [] or {}
// 替换对应字符串为空，for循环替换之后，最后结果必定是空字符串
func isValid(s string) bool {
	for i := 0; i < len(s); i++ {
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
	}
	return s == ""
}

func main() {
	fmt.Println(isValid("({[]})"))
	fmt.Println(isValid("({[]}"))
}
