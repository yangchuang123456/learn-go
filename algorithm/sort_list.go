package algorithm

/*
Sort a linked list in O(n log n) time using constant space complexity.

Example 1:

Input: 4->2->1->3
Output: 1->2->3->4
Example 2:

Input: -1->5->3->4->0
Output: -1->0->3->4->5

*/


type listNode struct {
  Val int
  Next *listNode
}

//归并排序,合并时由于是链表不需要额外空间
func sortList(head *listNode) *listNode {
	if head ==nil{
		return nil
	}

	if head.Next == nil{
		return head
	}

	slow := head
	fast := head

	for{
		if fast.Next == nil || fast.Next.Next == nil{
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow
	right := mid.Next
	mid.Next = nil
	sortedLeft := sortList(head)
	sortedRight := sortList(right)

	tmpHead := &listNode{}
	res:=tmpHead
	//merge
	for{
		if sortedRight==nil&&sortedLeft==nil{
			break
		} else if sortedRight == nil && sortedLeft!=nil{
			tmpHead.Next=sortedLeft
			sortedLeft=sortedLeft.Next
		} else if sortedRight !=nil && sortedLeft == nil{
			tmpHead.Next=sortedRight
			sortedRight=sortedRight.Next
		}else {
			if sortedLeft.Val <= sortedRight.Val{
				tmpHead.Next = sortedLeft
				sortedLeft = sortedLeft.Next
			}else{
				tmpHead.Next = sortedRight
				sortedRight = sortedRight.Next
			}
		}

		tmpHead=tmpHead.Next
	}
	return res.Next
}