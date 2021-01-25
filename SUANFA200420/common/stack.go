package common

import (
	"container/list"
	"log"
)

//list是go内置的，利用list组成stack
type Stack struct {
	elements *list.List
}

func NewStack() *Stack {
	l := list.New()
	return &Stack{l}
}

func (q *Stack) Empty() bool {
	return q.Size() == 0
}

func (q *Stack) Size() int {
	return q.elements.Len()
}

//这个是一个list，1,2,3  那么1是开头，3是最后
//如果换成栈的话，是不是就倒过来了呢？压栈的和出栈的
//top函数应该是返回栈顶元素
func (q *Stack) Top() interface{} {
	return q.elements.Back().Value
}

//Push函数把元素添加到栈顶，但是如果是按照list来看的话，那就是在后面添加元素
func (q *Stack) Push(v interface{}) {
	q.elements.PushBack(v)
}

//pop函数移除栈顶元素
func (q *Stack) Pop() interface{} {
	return q.elements.Remove(q.elements.Back())
}
func (q *Stack) Print() {
	for e := q.elements.Front(); e != nil; e = e.Next() {
		log.Println(e.Value)
	}
}
