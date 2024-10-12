//两数之和
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i, x := range nums {
		if p, ok := hash[target-x]; ok {
			return []int{p, i}
		}
		hash[x] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 25}
	fmt.Println(twoSum(nums, 9))
}
