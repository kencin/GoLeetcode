package main

import "fmt"

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type List struct {
	Head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(val int) *List {
	return &List{
		Head: NewListNode(val),
	}
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func (list *List) insert(val int) {
	Head := list.Head
	for Head.Next != nil {
		Head = Head.Next
	}
	Head.Next = NewListNode(val)
}

func (list *List) show() {
	Head := list.Head
	fmt.Print(Head.Val)
	fmt.Print("=>")
	for Head.Next != nil {
		Head = Head.Next
		fmt.Print(Head.Val)
		fmt.Print("=>")
	}
	fmt.Println("NULL")
}

func rotateRight(head *ListNode, k int) *ListNode {
	//1. 用余数确认每一个数据的位置，然后重新生成链表，但这样肯定效率不好。而且对于leetcode来说不知道insert函数无法解答
	//2. 计算出谁是头，以头为中心，进行划分拼接
	if k == 0 || head == nil {
		return head
	}
	//记录head
	firstHead, theNewEnd := head, head
	count := 1
	for head.Next != nil {
		head = head.Next
		count++
	}
	if k%count == 0 {
		return firstHead
	}
	//链表改为循环链表
	head.Next = firstHead
	theNewHead := count - k%count
	for i := 0; i < count; i++ {
		if i+1 == theNewHead {
			theNewEnd = firstHead
		}
		if i == theNewHead {
			head = firstHead
			fmt.Println(head.Val)
		}
		firstHead = firstHead.Next
	}
	theNewEnd.Next = nil
	return head

}

func main() {
	list := NewList(1)
	listNum := []int{2}
	for _, i := range listNum {
		list.insert(i)
	}
	list.show()
	list.Head = rotateRight(list.Head, 2)
	list.show()
}
