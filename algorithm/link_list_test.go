package algorithm

import (
	"container/heap"
	"log"
	"sort"
	"testing"
)

/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/

//对多个有序链表进行合并操作，并且分析其算法复杂度
type ListNode struct {
	Val  int
	Next *ListNode
}

//暴力合并，先遍历链表所有节点将其放到数组中，将数组进行排序，然后将排序后的数组从新构造成有序链表即可
//算法的时间复杂度为　N + N*logN + N  = N*logN (使用排序算法的时间复杂度为N*logN)
//算法空间复杂度为O(N) +O(N) = O(N) (排序空间复杂度＋新建链表的空间复杂度)
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	array := make([]int, 0)

	for _, list := range lists {
		for ; list != nil; list = list.Next {
			array = append(array, list.Val)
		}
	}
	sort.Ints(array)

	retList := &ListNode{}
	tmpList := retList
	for i, val := range array {
		tmpList.Val = val
		if i != len(array)-1 {
			tmpList.Next = &ListNode{}
			tmpList = tmpList.Next
		}
	}

	return retList
}

//第一个值为lists中的链表索引,第二个为节点值
type sortValue [2]int

type sortHeap []sortValue

func (s sortHeap) Len() int {
	return len(s)
}

func (s sortHeap) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func (s sortHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *sortHeap) Push(x interface{}) {
	*s = append(*s, x.(sortValue))
}

func (s *sortHeap) Pop() interface{} {
	if s.Len() == 0{
		return nil
	}
	res := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]

	return res
}

//分别对k个链表的头数据选取最小的追加到新创建的链表中，在对于k个链表头数据选取最小元素使用优先队列的方式来实现，可以降低排序的时间复杂度．
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0{
		return nil
	}
	s := make(sortHeap, 0)
	heap.Init(&s)
	res := &ListNode{}
	tmpRes := res
	tmpLists := lists

	//push 所有list第一个节点值
	for k, v := range tmpLists {
		if v != nil{
			heap.Push(&s, sortValue{k, v.Val})
			tmpLists[k] = tmpLists[k].Next
		}
	}

	if s.Len() ==0{
		return nil
	}

	for {
		popValue := heap.Pop(&s)
		res.Val = popValue.(sortValue)[1]
		index := popValue.(sortValue)[0]
		if tmpLists[index] !=nil{
			heap.Push(&s,sortValue{index,tmpLists[index].Val})
			tmpLists[index] = tmpLists[index].Next
		}

		if s.Len() ==0{
			break
		}
		res.Next = &ListNode{}
		res = res.Next
	}

	return tmpRes
}

func getTestList() []*ListNode {
	list1Len := 5
	list2Len := 6
	list3Len := 7

	list1 := &ListNode{}
	list2 := &ListNode{}
	list3 := &ListNode{}
	lists := []*ListNode{list1, list2, list3}

	tmpList1 := list1
	for i := 0; i < list1Len; i++ {
		tmpList1.Val = i * 2
		log.Println("the list1 node data is:", tmpList1.Val)
		if i != list1Len-1 {
			tmpList1.Next = &ListNode{}
			tmpList1 = tmpList1.Next
		}
	}

	tmpList2 := list2
	for i := 0; i < list2Len; i++ {
		tmpList2.Val = i*2 + 1
		log.Println("the list2 node data is:", tmpList2.Val)
		if i != list2Len-1 {
			tmpList2.Next = &ListNode{}
			tmpList2 = tmpList2.Next
		}
	}

	tmpList3 := list3
	for i := 0; i < list3Len; i++ {
		tmpList3.Val = i * 3
		log.Println("the list3 node data is:", tmpList3.Val)
		if i != list3Len-1 {
			tmpList3.Next = &ListNode{}
			tmpList3 = tmpList3.Next
		}
	}

	return lists
}

func Test_mergeKLists(t *testing.T) {
	testListNode := getTestList()
	res := mergeKLists1(testListNode)

	for ; res != nil; res = res.Next {
		log.Println("the list node data is:", res.Val)
	}
}
