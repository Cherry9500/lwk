// 合并两个有序链表
package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

//建立链表元素
func (head *ListNode) build(list []int) {
	// if len(list) == 0 {
	// 	return nil
	// }

	for i := len(list) - 1; i >= 0; i-- {
		p := &ListNode{data: list[i]}
		// fmt.Println(p.data)
		p.next = head.next
		head.next = p
	}
}

//链接链表元素
func (head *ListNode) travel() {
	for p := head.next; p != nil; p = p.next {
		fmt.Print(p.data)
		if p.next != nil {
			fmt.Print("->")
		}
	}
	fmt.Println("<nil>")
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.data < list2.data {
		list1.next = mergeTwoLists(list1.next, list2)
		return list1
	}
	list2.next = mergeTwoLists(list2.next, list1)
	return list2
}

func main() {
	nums1 := []int{1, 2, 4, 6}
	nums2 := []int{1, 3, 4, 8}
	l1 := &ListNode{}
	l2 := &ListNode{}
	l1.build(nums1)
	l2.build(nums2)
	l1.travel()
	l2.travel()

	result := mergeTwoLists(l1, l2).next
	result.travel()
}
