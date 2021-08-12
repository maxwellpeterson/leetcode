package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, nn int) *ListNode {
	var target *ListNode // The node before the node we want to remove.

	for ii, node := 1, head; node.Next != nil; ii, node = ii+1, node.Next {
		if ii == nn {
			target = head
		} else if ii > nn {
			target = target.Next
		}
	}

	if target == nil {
		return head.Next
	} else {
		target.Next = target.Next.Next
		return head
	}
}
