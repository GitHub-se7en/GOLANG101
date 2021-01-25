package main
/*
Problem:
- Given a string, find the length of the longest substring in it with no more
  than k distinct characters.
Example:
- Input: string="araaci", k=1
  Output: 2
  Explanation: Longest substring with no more than 1 distinct characters is "aa".
- Input: string="araaci", k=2
  Output: 4
  Explanation: Longest substring with no more than 2 distinct characters is "araa".
- Input: string="araaci", k=3
  Output: 5
  Explanation: Longest substring with no more than 3 distinct characters is "araac".
Approach:
- Use a hashmap to remember the frequency of each character we have seen.
- Insert characters until we have k distinct characters in the map to be consider a
  window.
- Shrink the window until there is no more k distinct characters in the map and keep
  updating the maximum window length at each step.
Cost:
- O(n) time, O(k) space where k is the number of characters in the map.
*/
/**
最长子串K个不同值，灵魂是的，不同值使用的是map实现然后最长是使用最开始的位置
麻烦的地方不是在这，而是在模糊，不清楚，看到的是两个for循环，
看懂算法，以数据结构分类，先是串，然后是树，然后是图，这里面有动串，有检查串，
就像是反转字符，即便是知道，但是关键点是第一步的整体反转啊
*/
func ()  {

}