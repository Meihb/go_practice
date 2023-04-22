package main

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

 

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
 

提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案


*/
func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)
	for k, v := range nums {
		if diffIndex, ok := hm[target-v]; ok {
			return []int{diffIndex, k}
		} else {
			hm[v] = k
		}
	}

	return []int{}

}

/*
作者：joelang
链接：https://leetcode-cn.com/problems/two-sum/solution/tong-guo-hash_mapji-suan-jie-guo-by-joel-5ilc/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func twoSum2(nums []int, target int) []int {
	hash_map := make(map[int]int)
	for k, v := range nums {
		if vv, ok := hash_map[target-v]; ok {
			return []int{vv, k}
		} else {
			hash_map[v] = k
		}
	}
	return []int{}
}

/**
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

 

示例 1：
	2->4->3
	5->6->4
=   7->0->8

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]

示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
 

提示：

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {

}

func main() {

	//res:=twoSum([]int{2,7,11,15},9)
	//res := twoSum2([]int{3, 3}, 6)
	//fmt.Println(res)
	//res = twoSum([]int{3, 3}, 6)
	//fmt.Println(res)
}
