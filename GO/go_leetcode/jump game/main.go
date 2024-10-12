// 跳跃游戏
package main

import "fmt"

//贪心算法
// func jumpgame(nums []int) bool { //3,2,1,0,4
// 	n := len(nums)
// 	maxPos := 0
// 	for i := 0; i < n; i++ {
// 		if i > maxPos {
// 			return false
// 		}
// 		maxPos = max(maxPos, i+nums[i])
// 	}
// 	fmt.Println(maxPos)
// 	return true
// }

//反向贪心算法
func jumpgame(num []int) bool {
	n := len(num)
	target := n - 1
	for i := n - 2; i >= 0; i-- {
		if i + num[i] >= target {
			target = i
		}
	}
	return target == 0
}

func main() {
	// nums1 := []int{3, 0, 9, 1, 0, 4}
	// fmt.Printf("jumpgame(nums1): %v\n", jumpgame(nums1))
	nums2 := []int{3, 2, 1, 0, 4}
	fmt.Printf("jumpgame(nums2): %v\n", jumpgame(nums2))
}
