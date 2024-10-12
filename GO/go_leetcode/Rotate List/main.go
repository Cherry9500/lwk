// 旋转链表
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		n++
	}
	cur.Next = head
	k = n - k%n
	for i := 0; i < k; i++ {
		cur = cur.Next
	}
	head = cur.Next
	cur.Next = nil
	return head
}

func build(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	p := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{Val: nums[i]}
		p.Next = node
		p = p.Next
	}
	return head
}

func (head *ListNode) travel() {
	for p := head; p != nil; p = p.Next {
		fmt.Print(p.Val)
		if p.Next != nil {
			fmt.Print("->")
		}
	}
	fmt.Println("<nil>")
}

func main() {

	nums := []int{1, 2, 3, 4, 5}
	head := build(nums)
	head.travel()
	rotateRight(head, 2).travel()

}
