// 合并两个有序数组
package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, nums2 []int) []int {
	sorted := []int{}
	sorted = append(nums1, nums2...)
	sort.Ints(sorted)
	return sorted
}
func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}
	fmt.Println(merge(nums1, nums2))
}
