package leetcode

import "fmt"


type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var cur *ListNode = head

	for cur != nil {
		head = head.Next
		cur.Next = prev
		prev = cur
		cur = head
	}

	return prev
}

func Test_reverseList() {
	fmt.Println("test reverse list")
	var a ListNode = ListNode{3, nil}
	var b ListNode = ListNode{2, &a}
	var c ListNode = ListNode{1, &b}

	ret := reverseList(&c)
	//ret := &c

	for ret != nil {
		fmt.Println("value = ", ret.Val)
		ret = ret.Next
	}

}