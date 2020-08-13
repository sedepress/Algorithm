package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	cur := dummy

	for i := 1; i <= n + 1; i++ {
		cur = cur.Next
	}

	for cur != nil {
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}

func main() {
	
}
