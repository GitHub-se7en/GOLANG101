package main

func main() {
	//要什么？要一个装着数据的切片，有输入和输出的结果
	//这种写法叫做切片字面量的写法
	input1 := []int{1, 2, 3, 4}
	input2 := []int{4, 5, 6, 7}
	output := []int{1, 2, 3, 4, 5, 6, 7}
	//目的是什么？是两个数组取出交集来

	//思路是什么？

}
func userMapAndSliceToGetResult(input1, input2 []int) []int {
	thisIsMap := make(map[int]string)
	thisIsSlice := make([]int, 0)
	thisIsResult := []int{nil} //这个写法真的是第一次见到，作用和var是相同的
	//遍历切片，放到map里面，实现去重
	for _, value := range input1 {
		thisIsMap[value] = "true"
	}
	for _, value := range input2 {
		thisIsSlice = append(thisIsSlice, value)
	}
	//最后一步是遍历map数据，把slice里面的数据全部删除掉，剩下的就是
	//for key, value := range thisIsMap {
	//	thisIsSlice//我查了一圈，竟然没找到根据value获取某个值的方法，如果是在遍历一次的话，那不就相当于走了多余的路了吗？
	//
	//}
	//for key, value := range thisIsSlice {
	//倒过来也不行，因为没有现成的delete方法
	//}
	for _, value := range input2 {
		if thisIsMap[value] == "true" {
			thisIsResult = append(thisIsResult, value)
		}
	} //确实是这样，官方的解法是最简洁的，只存不取是不可能的，因为我的slice里面也是需要取数据的
	return thisIsResult
}
