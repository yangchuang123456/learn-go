package algorithm

import (
	"github.com/stretchr/testify/assert"
	"log"
	"sort"
	"testing"
)

/*
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/
func threeSumClosest(nums []int, target int) int {
	//将数组排序
	sort.Ints(nums)
	var first, last, smallValue, bigValue int
	flag := false
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		if !flag {
			smallValue = nums[i] + nums[i+1] + nums[i+2]
			bigValue = nums[len(nums)-1] + nums[len(nums)-2] + nums[len(nums)-3]
			flag = true
		}

		first = i + 1
		last = len(nums) - 1

		for ; first < last; {
			tmpDiff := nums[i] + nums[first] + nums[last]
			log.Println("the tmpDiff smallValue bigValue is111:", tmpDiff, smallValue, bigValue)
			if tmpDiff < target {
				if tmpDiff > smallValue {
					smallValue = tmpDiff
				}
				first++
			} else if tmpDiff > target {
				if tmpDiff < bigValue {
					bigValue = tmpDiff
				}
				last--
			} else {
				return target
			}

			log.Println("the tmpDiff smallValue bigValue is222:", tmpDiff, smallValue, bigValue)
		}
	}

	log.Println("the smallValue and bigValue is:", smallValue, bigValue)

	if target-smallValue < bigValue-target {
		return smallValue
	} else {
		return bigValue
	}
}

func Test_threeSumClosest(t *testing.T) {
	a := []int{1, 1, 1, 0}
	target := -100

	result := threeSumClosest(a, target)
	assert.Equal(t, 2, result)

	a = []int{-1, 2, 1, -4}
	target = 1
	result = threeSumClosest(a, target)
	assert.Equal(t, 2, result)

	var s string
	s = "hello"
	log.Println("the s len is:", len(s),s[0])
}
