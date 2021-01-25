package main

import (
	"GOLANG101/SUANFA200420/common"
	"log"
)

/**
Problem:
- Reverse a linked list in-place.
*/
/**
链表反转最体现出了变量的使用，正式的项目里面没有这么干的，都是一个变量承载了一次数据，没有让这个变量承载这么多次变量的
*/
/**
如何遍历一个链表？这些节点都是一个连着一个，只是改动指针指向
*/
func main() {
	// define tests input.
	//t0 := common.NewListNode(0)

	//t1 := common.NewListNode(1)

	t2 := common.NewListNode(1)
	t2.AddNext(2)

	t3 := common.NewListNode(1)
	t3.AddNext(2)
	t3.AddNext(3)

	// define tests output.
	tests := []struct {
		in       *common.ListNode
		expected []int
	}{
		//{t0, []int{0}},
		//{t1, []int{1}},
		//{t2, []int{2, 1}},
		{t3, []int{3, 2, 1}},
	}

	for i, tt := range tests {
		log.Println(i)
		node := reverseLinkedList(tt.in)
		common.Equal(tt.expected, common.LinkedListToSlice(node))
	}
}
func reverseLinkedList(node *common.ListNode) *common.ListNode {
	//有一个思考的方法可以彻底捋清楚这个逻辑，那就是先固定好三个指针，然后再扯别的
	current := node
	var previous, next *common.ListNode //用var声明意味着是nil值
	//目前这三个指针current和previous在正确的位置上
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	//这样到最后的结果就是previous指针指向了第一个节点，也是反转前的最后一个节点
	//这里只有三个指针,其余的节点是没办法获取的
	return previous
}
