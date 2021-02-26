package algorithm

import (
	"log"
	"sort"
	"testing"
)

/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/

//先将数组按升序排列，然后以左侧数据为基点，利用双指针依次遍历计算．对于重复数据直接跳过
func threeSum1(nums []int) [][]int{
	//sort nums
	sort.Ints(nums)
	log.Println("the nums is:",nums)
	res := make([][]int,0)
	var first,last int
	for i:=0;i<len(nums)-2;i++{
		//处理重复的情况
		if i!=0 && nums[i] == nums[i-1]{
			continue
		}

		//最小的数大于0直接退出
		if nums[i] > 0{
			break
		}

		first = i+1
		last = len(nums)-1

		for ;first < last;{
			//如果三个数符号相同，直接退出
			if nums[i]*nums[last] >0{
				break
			}
			//处理重复情况
			if first!= i+1 && nums[first] == nums[first-1]{
				first ++
				continue
			}
			if nums[i]+nums[first]+nums[last] > 0{
				last --
			}else if nums[i]+nums[first]+nums[last] < 0{
				first ++
			}else{
				res = append(res,[]int{nums[i],nums[first],nums[last]})
				last --
				first ++
			}
		}
	}

	return res
}

func Test_threeSum(t *testing.T){
	nums := []int{-2,0,0,2,2}
	result:=threeSum1(nums)
	log.Println("the result is:","result",result)

}
