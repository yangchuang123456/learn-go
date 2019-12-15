package algorithm

import (
	"log"
	"testing"
)

/*
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

//全排列问题可以使用回溯法的排列树框架解决
func backTrack(t int,nums []int,res *[][]int) {
	if t == len(nums){
		tmp := make([]int,len(nums))
		copy(tmp,nums)
		*res = append(*res,tmp)
	}
	for i:=t;i<len(nums);i++{
		nums[i],nums[t] = nums[t],nums[i]
		backTrack(t+1,nums,res)
		nums[t],nums[i] = nums[i],nums[t]
	}
}

func permute(nums []int) [][]int {
	if len(nums) == 0{
		return [][]int{}
	}
	res := make([][]int,0)
	backTrack(0,nums,&res)
	return res
}

func Test_permute(t *testing.T){
	nums := []int{1,2,3}
	res:=permute(nums)
	log.Println("the res is:",res)
}