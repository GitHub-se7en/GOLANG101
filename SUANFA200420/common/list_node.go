package common

//一个简单的链表就这么形成了
type ListNode struct {
	Value int
	Next  *ListNode
}

//返回一个类型的指针，但是next没有值
func NewListNode(v int) *ListNode {
	return &ListNode{
		Value: v,
		Next:  nil,
	}
}

//这个方法的功能是给定一个链表，然后就能从头找到最后一个
func (l *ListNode) AddNext(v int) {
	for n := l; n != nil; n = n.Next {
		if n.Next == nil {
			new := NewListNode(v)
			n.Next = new
			//必须break，要不然的话，就会出现循环新的值的情况
			break
		}
	}
}

func LinkedListToSlice(node *ListNode) []int {
	var out []int

	for n := node; n != nil; n = n.Next {
		out = append(out, n.Value)
	}

	return out
}
