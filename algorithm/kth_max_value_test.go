package algorithm

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"learn-go/algorithm/common"
	"sort"
	"testing"
)

/*
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/

//排序后直接返回后k个的值
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

type MaxHeap struct {
	common.SortHeap
}

func (h *MaxHeap) Less(i,j int) bool{
	return h.SortHeap[i] > h.SortHeap[j]
}

//使用堆
func findKthLargest2(nums []int, k int) int {
	maxHeap := MaxHeap{SortHeap:nums}
	heap.Init(&maxHeap)
	var ret int
	for i:=0;i<k;i++{
		ret = heap.Pop(&maxHeap).(int)
	}
	return  ret
}

func Test_findKthLargest(t *testing.T){
	testSlice := []int{3,2,1,5,6,4}
	ret := findKthLargest1(testSlice,2)
	assert.Equal(t,5,ret)

	ret2 := findKthLargest2(testSlice,2)
	assert.Equal(t,5,ret2)
}