package algorithm

import (
	"log"
	"testing"
)

/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]

*/

func merge(nums1 []int, m int, nums2 []int, n int)  {
	insert := func(num int,nums []int)[]int{
		log.Println("insert num is:",num)
		log.Println("the nums is:",nums)
		left :=0
		right := len(nums)-1
		mid := (left+right)/2
		ret := make([]int,0)
		if num <= nums[left]{
			ret = append(ret,num)
			ret = append(ret,nums[:]...)

		}else if num >= nums[right]{
			ret = append(ret,nums[:]...)
			ret = append(ret,num)
		}else{
			for {
				if left == right-1{
					break
				}
				if nums[mid] < num{
					left = mid
				}else if nums[mid] > num{
					right = mid
				}else{
					break
				}
				mid = (left+right)/2
			}

			log.Println("the mid is:",mid)
			ret = append(ret,nums[:mid+1]...)
			ret = append(ret,num)
			ret = append(ret,nums[mid+1:]...)
		}


		return ret
	}

	if m == 0{
		for i,v:=range nums2{
			nums1[i] =v
		}
		return
	}

	tmpNums := nums1[:m]
	for i:=0;i<n;i++{
		tmpNums=insert(nums2[i],tmpNums)
	}
	log.Println("the tmpNums is:",tmpNums)

	for i,v:=range tmpNums{
		nums1[i] = v
	}
}

func Test_merge(t *testing.T){
	num1:= []int{0}
	num2:= []int{1}
	merge(num1,0,num2,1)
	log.Println("the num1 is:",num1)
}