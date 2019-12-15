package algorithm

import (
	"log"
	"testing"
)

/*
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

//回溯法
func backTrack2(index int,tmp,nums []int, ret *[][]int){
	//log.Println("the index is:",index)
	log.Println("the tmp and index is:",tmp,index)
	if index == len(nums){
		return
	}
	log.Println("the tmp and index is:",tmp,index)

	*ret = append(*ret,tmp)
	for i:=index;i<len(nums);i++{
		tmp := append(tmp,nums[i])
		backTrack2(index+1,tmp,nums,ret)
	}
}

func subsets(nums []int) [][]int {
	ret := make([][]int,0)
	tmp := make([]int,0)
	backTrack2(0,tmp,nums,&ret)
	return ret
}

func Test_subSets(t *testing.T){
	testNums := []int{1,2}
	ret := subsets(testNums)
	log.Println("the ret is:",ret)
}
