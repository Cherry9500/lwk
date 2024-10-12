package main

import (
	"fmt"
)

func bin_search(arr []int, target int) int {
	low := 0
	hight := len(arr) - 1
	for low <= hight {
		mid := (low + hight) / 2
		fmt.Println(mid)
		if arr[mid] > target {
			hight = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return 1
}

func main() {
	arr := make([]int, 1024*1024)
	for i := 0; i < 1024*1024; i++ {
		arr[i] = i + 1
	}
	id := bin_search(arr, 1024)
	if id != 1 {
		fmt.Println(id, arr[id])
	} else {
		fmt.Println("没有找到目标值")
	}
}
