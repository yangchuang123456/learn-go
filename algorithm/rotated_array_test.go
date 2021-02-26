package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

*/


func search(nums []int, target int) int {
	if len(nums) == 0{
		return -1
	}

	leftIndex := 0
	rightIndex := len(nums)-1
	midIndex := (leftIndex + rightIndex)/2

	for {
		if leftIndex>rightIndex{
			return -1
		}

		if leftIndex == rightIndex && nums[leftIndex] != target{
			return -1
		}
		leftValue := nums[leftIndex]
		midValue := nums[midIndex]
		rightValue := nums[rightIndex]
		if target == leftValue{
			return leftIndex
		}

		if target == midValue{
			return midIndex
		}

		if target == rightValue{
			return rightIndex
		}

		if midValue > leftValue{
			if target < midValue && target >leftValue{
				leftIndex = leftIndex +1
				rightIndex = midIndex -1
				midIndex = (leftIndex + rightIndex)/2
			}else if target < leftValue || target > midValue{
				leftIndex = midIndex +1
				midIndex = (leftIndex + rightIndex)/2
			}
		}else{
			if target > midValue && target <rightValue{
				leftIndex =midIndex +1
				rightIndex = rightIndex -1
				midIndex = (leftIndex + rightIndex)/2
			} else if target < midValue || target > rightValue{
				rightIndex = midIndex -1
				midIndex = (leftIndex + rightIndex)/2
			}
		}
	}
}

func Test_search(t *testing.T){
	arr := []int{1,3}
	target := 0
	index := search(arr,target)
	assert.Equal(t,-1,index)
}
