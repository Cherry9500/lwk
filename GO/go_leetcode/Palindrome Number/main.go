//回文数
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}

	if num < 10 {
		return false
	}

	s := strconv.Itoa(num) //类型转换
	fmt.Println(s, reflect.TypeOf(s))
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(9))
}
