package main

import "fmt"

func Quicksort(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}

	first := arr[0]         //[15]
	low := make([]int, 0)   //[9,10,2,4,8,12] 9 [2,4,8]+[10,12]
	mid := make([]int, 0)   //[]
	hight := make([]int, 0) //[30,45,63,234] 30 []+[45,63,234]
	mid = append(mid, first)
	for _, v := range arr {
		if v < first {
			low = append(low, v)
		} else if v > first {
			hight = append(hight, v)
		} else {
			mid = append(mid, v)
		}
	}
	low, hight = Quicksort(low), Quicksort(hight)
	myarry := append(append(low, mid...), hight...)
	return myarry
}

func main() {
	arr := []int{15, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(Quicksort(arr))

}
