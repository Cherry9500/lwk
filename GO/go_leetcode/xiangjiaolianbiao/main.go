//相交链表
package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func main() {
	sort := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(sort, target))
}
