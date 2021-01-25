package main

func main() {

}

type Node struct {
	value int
	next  *Node
}

func hasCycle(head *Node) bool {
	fast := head
	slow := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			return true
		}
	}
	return false
}
