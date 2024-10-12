// 环形链表
package main

import (
	"fmt"
)

type ListNode struct {
	val  int
	Next *ListNode
}

func createRingNodeList(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{val: nums[0]}
	tail := head
	for i := 0; i < len(nums); i++ {
		tail.Next = &ListNode{val: nums[i]}
		tail = tail.Next
	}

	if pos >= 0 {
		p := head
		for pos > 0 {
			p = p.Next
			pos--
		}
		tail.Next = p
	}
	return head
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != fast {
		if slow == nil || fast == nil {
			return false
		}
		slow, fast = slow.Next, fast.Next.Next
	}

	return true
}

func main() {
	nums, pos := []int{3, 2, 0, -4}, 1
	head := createRingNodeList(nums, pos)
	fmt.Println(hasCycle(head))
}
