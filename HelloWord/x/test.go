package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) (ret *ListNode) {
	//先判空，
	if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 == nil {
		return
	}
	//指针初始化
	//pro :=&l1 指针的初始化为啥不是&l1
	pro := l1

	if l1.Val < l2.Val {
		pro = l1
		l1 = l1.Next
	} else {
		pro = l2
		pro = l2.Next
	}
	ret = pro
	for {
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				pro.Next = l1
				l1 = l1.Next
				pro = pro.Next
			} else {
				pro.Next = l2
				l2 = l2.Next
				pro = pro.Next
			}

		} else if l1 == nil && l2 != nil {
			pro.Next = l2
			break
		} else if l1 != nil && l2 == nil {
			pro.Next = l1
			break
		} else {
			return ret
		}
	}

	return ret
}

func Append(list *ListNode, node *ListNode) {

	for list != nil {
		temp := list.Next
		if temp == nil {
			list.Next = node
			break
		}
		list = list.Next
	}
}

func main() {

	headNode := &ListNode{
		Val:  0,
		Next: nil,
	}
	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	Append(headNode, node1)
	Append(headNode, node2)

	headNode1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	node3 := &ListNode{
		Val: 2,
	}
	node4 := &ListNode{
		Val: 3,
	}
	Append(headNode1, node3)
	Append(headNode1, node4)
	ret := mergeTwoLists(headNode, headNode1)
	fmt.Printf("%d", ret)
}
