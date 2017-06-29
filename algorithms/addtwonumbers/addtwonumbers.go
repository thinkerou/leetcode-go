// Source : https://leetcode.com/problems/add-two-numbers/#/description
// Author : thinkerou
// Date   : 2017-06-30

/*
 *
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order and each of their nodes contain a single digit.
 * Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 *
 */

package addtwonumbers

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var pListNode *ListNode = new(ListNode)
	cur := pListNode
	i := 0
	for l1 != nil || l2 != nil {
		a := 0
		b := 0
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}
		sum := a + b + i
		i = sum / 10
		j := sum % 10
		cur.Val = j
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}
		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
		if l1 != nil || l2 != nil {
			p := new(ListNode)
			cur.Next = p
			cur = p
		}
	}
	if i > 0 {
		p := new(ListNode)
		p.Val = i
		p.Next = nil
		cur.Next = p

	}
	return pListNode
}
