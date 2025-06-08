package main

import (
	"fmt"
	"strconv"
)

// 思路
// 字符串反转，值和原字符串比较
// 相等 返回 true 不相等返回 false
func isPalindrome(x int) bool {
	strNumber := strconv.Itoa(x)
	reverseStrNumber := ""
	for length := len(strNumber); length > 0; length-- {
		reverseStrNumber += string(strNumber[length-1])
	}
	reverseNum, error := strconv.Atoi(reverseStrNumber)
	if error != nil {
		fmt.Println("Failure to cast String to int")
	}
	return reverseNum == x
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
}
