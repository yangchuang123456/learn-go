package algorithm

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

//leetcode
/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/

//由于两个数组都是有序的，因此可以根据数组长度判断合并后的有序数组的中位数的index,
//遍历，每次分别记录两个有序数组的下标，若数据小则其被选入到合并的数组中，用mid1记录其值，然后数组下标加１
//当一个数组已被挑选完后，则每次从剩下的数组挑选即可
//当遍历到中位数的下标时，根据总数组长度的奇偶性计算中位数
//算法的时间复杂度为log(n+m)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	var mid0, numIndex1, numIndex2, mid1 int
	totalNum := len1 + len2

	for i := 0; i <= (totalNum-1)/2+1; i++ {
		if numIndex1 < len1 && numIndex2 < len2 {
			if nums1[numIndex1] <= nums2[numIndex2] {
				mid1 = nums1[numIndex1]
				numIndex1++
			} else {
				mid1 = nums2[numIndex2]
				numIndex2++
			}
		} else {
			if numIndex1 < len1 {
				mid1 = nums1[numIndex1]
				numIndex1++
			} else {
				mid1 = nums2[numIndex2]
				numIndex2++
			}
		}

		if totalNum%2 == 1 {
			if i == (totalNum-1)/2 {
				return float64(mid1)
			}
		} else {
			if i == (totalNum-1)/2 {
				mid0 = mid1
			} else if i == (totalNum-1)/2+1 {
				return (float64(mid0) + float64(mid1)) / 2
			}
		}

	}

	return float64(0)
}

func Test_GetMedianFromSortedSlice(t *testing.T) {
	slice1 := []int{1, 5, 7}
	slice2 := []int{3, 9}

	median := findMedianSortedArrays(slice1, slice2)
	assert.Equal(t, 5.0, median)

	log.Println("the median is:", median)

	slice1 = []int{1, 5, 7}
	slice2 = []int{3, 6, 9}
	median = findMedianSortedArrays(slice1, slice2)
	log.Println("the median is:", median)
	assert.Equal(t, 5.5, median)

}
