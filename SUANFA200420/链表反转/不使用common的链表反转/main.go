package main

//看清楚了,这可是连接的自己啊
type Node struct {
	value int
	//也就是说,下一个依然自己的指针类型
	next *Node
}

func main() {

}

func reverse(head *Node) *Node {
	var pre *Node
	//这他娘的还有死循环在这
	for head != nil {
		temp := head.next
		head.next = pre
		head = pre
		head = temp
	}
	return pre
}
