package main

import "GOLANG101/SUANFA200420/common"

/**
- Implement a stack with a method getMax() that returns the largest element in
  the stack in O(1) time.
Approach:
- We use two stack implementation themselves: one holds all the items and the
  other holds all the maximum values after each push() and pop().
- That way, we could keep track of your maximum value up to date in constant
  time.
*/
/**
要在O(1)的时间里面完成取出栈里面的最大值，实现的思路很简单，包装一个栈，在压栈和出栈的时候就已经把每一步的最大值都放到另外一个栈里面了
*/
type maxStack struct {
	stack    *common.Stack
	maxStack *common.Stack
}

func newMaxStack() *maxStack {
	s := common.NewStack()
	ms := common.NewStack()
	return &maxStack{
		stack:    s,
		maxStack: ms,
	}
}

func (s *maxStack) push(i int) {
	//压栈，其实就是在list后面添加数据
	s.stack.Push(i)
	//还是从底到顶的写法，这个和求list中的最大值是相似的，
	if s.maxStack.Size() == 0 || i >= s.maxStack.Top().(int) {
		s.maxStack.Push(i)
	}
}

//push和pop的时候同步更新最大栈，
//突然间想到，如果是按照只保留最大的一个数的做法的话，万一这个栈pop了，很有可能自己就求不出最大数据了
//而按照这种做法的话，就能实现按照大小的排列实现，
func (s *maxStack) pop() int {
	i := s.stack.Pop().(int)
	if i == s.maxStack.Top().(int) {
		s.maxStack.Pop()
	}
	return i
}

func (s *maxStack) getMax() int {
	return s.maxStack.Top().(int)
}
