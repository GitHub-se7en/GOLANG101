package main

/*
Problem:
- Determine if a singly-linked list has a cycle.
Approach:
- Keep two pointers starting at the first node such that: every time one moves
  one node ahead, the other moves 2 nodes ahead.
- If the linked list has a cycle, the faster one will catch up with the slow
  one. Otherwise, the faster one will each the end.
Solution:
- Keep two pointers, a slow one and a fast one, at the beginning.
- Traverse the list and move the fast one 2 nodes once the slow one move a node.
- If the fast one catches the slow one, there exists a cycle.
Cost:
- O(n) time and O(1) space.
*/
/**
这个意思是用两个指针，第一个指针每向前移动一步，第二个指针就向前移动两步，如果有环的话，就会出现第二个追上第一个的情况
*/
func main() {

}
