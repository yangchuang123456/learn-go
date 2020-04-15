package algorithm

import (
	"github.com/stretchr/testify/assert"
	"github.com/yangchuang123456/learn-go/algorithm/common"
	"testing"
)

/*
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Note:
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

Example 1:

Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
Output: 1
Example 2:

Input: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
Output: 3
Follow up:
What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?

*/

//bst 中序遍历有序，因此中序遍历返回第k个节点值即可。中序遍历可以使用递归遍历或者非递归遍历的方式。
//二叉树的非递归遍历需要借用栈来进行
//此实现时间和空间复杂度均较高，猜测是由于使用通用stack导致
func kthSmallest(root *common.TreeNode, k int) int {
	if root == nil {
		return 0
	}
	stack := common.NewStack()
	stack.Push(root)

	var top interface{}
	var topNode *common.TreeNode
	var err error
	var flag int
	for {
		top, err = stack.Top()
		if err != nil {
			break
		}

		topNode = top.(*common.TreeNode)

//		log.Println("the topNode is:",topNode)
		for topNode.Left != nil {
			stack.Push(topNode.Left)
			topNode = topNode.Left
		}

//		log.Println("the stack len is:",stack.Len())

		for !stack.IsEmpty() {
			top, _ = stack.Pop()
			topNode = top.(*common.TreeNode)
			flag++
			if flag == k {
				return topNode.Val
			}
			if topNode.Right != nil {
				stack.Push(topNode.Right)
				break
			}

//			log.Println("the pop node is:",topNode)
		}
	}

	return 0
}

func Test_kthSmallest(t *testing.T) {
	a, b, c, d := 3, 1, 4, 2
	slice := []*int{
		&a, &b, &c, nil, &d,
	}
	tree := common.GetTestTree(slice)

	k := 1
	res := kthSmallest(tree, k)
	assert.Equal(t, 1, res)

	a, b, c, d, e, f := 5, 3, 6, 2, 4, 1
	slice = []*int{
		&a, &b, &c, &d, &e, nil, nil, &f,
	}
	tree = common.GetTestTree(slice)
	k = 3
	res = kthSmallest(tree, k)
	assert.Equal(t, 3, res)
}
