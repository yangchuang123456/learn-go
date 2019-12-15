package algorithm

import (
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"testing"
)

/*
Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.

Example 1:

Input: [1,2,3]

       1
      / \
     2   3

Output: 6
Example 2:

Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxLeftRightPathMaxSum(root *TreeNode) (leftMax ,rightMax int){
	if root == nil{
		return 0,0
	}

	if root.Right == nil && root.Left == nil{
		return root.Val,root.Val
	}

	leftMax,rightMax = root.Val,root.Val
	tmp1,tmp2 := MaxLeftRightPathMaxSum(root.Left)
	if tmp1 < tmp2{
		if tmp2 > 0{
			leftMax = tmp2 + root.Val
		}
	}else{
		if tmp1 > 0{
			leftMax = tmp1 + root.Val
		}
	}

	tmp1,tmp2 =MaxLeftRightPathMaxSum(root.Right)
	if tmp1 < tmp2{
		if tmp2 > 0{
			rightMax = tmp2 + root.Val
		}

	}else{
		if tmp1 > 0 {
			rightMax = tmp1 + root.Val
		}
	}

	return leftMax,rightMax
}


func maxPathSumFunc(root *TreeNode) (int,bool) {
	if root == nil{
		return  0,true
	}

	tmp1,tmp2 := MaxLeftRightPathMaxSum(root)
	log.Println("the tmp1 and tmp2 is:",tmp1,tmp2)
	var max1 int
	if root.Left !=nil && root.Right!=nil{
		max1 = tmp1+tmp2-root.Val
	} else if root.Left !=nil && root.Right==nil{
		max1 = tmp1
	}else if root.Left ==nil && root.Right!=nil{
		max1 = tmp2
	} else {
		return root.Val,false
	}

	log.Println("the max1 is:",max1)
	max2,leftNil := maxPathSumFunc(root.Left)
	log.Println("the max2 is:",max2,leftNil)
	tmp := max1
	if !leftNil{
		if max1 < max2{
			tmp = max2
		}
	}


	max3,rightNil := maxPathSumFunc(root.Right)
	log.Println("the max3 is:",max3,rightNil)

	if !rightNil{
		if tmp < max3{
			return max3,false
		}
	}

	return tmp,false
}

func maxPathSum(root *TreeNode) int{
	ret,_:=maxPathSumFunc(root)
	return ret
}

func creatCompleteTree(nodeNumber,index int) *TreeNode{
	var tree TreeNode
	if index < nodeNumber{
		tree.Left = creatCompleteTree(nodeNumber,2*index+1)
		tree.Right = creatCompleteTree(nodeNumber,2*index+2)
	}else{
		return nil
	}
	return &tree
}


type queue struct {
	data []interface{}
	mux sync.RWMutex
}

func NewQueue() *queue{
	return &queue{
		data:make([]interface{},0),
	}
}

func (q *queue) Enqueue(data interface{}) {
	q.mux.Lock()
	defer q.mux.Unlock()
	q.data = append(q.data,data)
}

func (q *queue) Dequeue() interface{}{
	q.mux.Lock()
	defer q.mux.Unlock()
	ret := q.data[0]
	q.data = q.data[1:]
	return ret
}

func (q *queue) Front() interface{}{
	q.mux.RLock()
	defer q.mux.RUnlock()
	return q.data[0]
}

func (q *queue) Size() int{
	return len(q.data)
}

func (q *queue) IsEmpty() bool{
	if len(q.data) == 0{
		return true
	}else{
		return false
	}
}

func getTestTree(slice []*int) *TreeNode{
	//先构造完全二叉树
	tree := creatCompleteTree(len(slice),0)

	q := NewQueue()
	q.Enqueue(tree)
	//层次遍历填入数据,并删除空节点
	i := 0
	for !q.IsEmpty(){
		tmpNode := q.Dequeue()
		if tmpNode != nil{
			node := tmpNode.(*TreeNode)
			node.Val = *slice[i]

			//处理左子树
			if 2*i+1 < len(slice){
				if slice[2*i+1] == nil{
					node.Left = nil
					q.Enqueue(nil)
				}else{
					q.Enqueue(node.Left)
				}
			}

			//处理右子树
			if 2*i+2 < len(slice){
				if slice[2*i+2] == nil{
					node.Right = nil
					q.Enqueue(nil)
				}else{
					q.Enqueue(node.Right)
				}
			}
		}
		i++
	}

	return tree
}

func inOrderTraverseTree(tree *TreeNode,slice *[]int){
	if tree !=nil{
		inOrderTraverseTree(tree.Left,slice)
		*slice = append(*slice,tree.Val)
		inOrderTraverseTree(tree.Right,slice)
	}
}

func Test_generateTree(t *testing.T){
	a,b,c,d,e := -10,9,20,15,7
	slice := []*int{
		&a,&b,&c,nil,nil,&d,&e,
	}
	tree := getTestTree(slice)
	tmpSlice := make([]int,0)
	inOrderTraverseTree(tree,&tmpSlice)
	assert.Equal(t,[]int{9,-10,15,20,7},tmpSlice)
	log.Println("the tree data is:",tmpSlice)

	a,b,c = 1,2,3
	slice = []*int{&a,&b,&c}
	tree = getTestTree(slice)
	tmpSlice = make([]int,0)
	inOrderTraverseTree(tree,&tmpSlice)
	assert.Equal(t,[]int{2,1,3},tmpSlice)
	log.Println("the tree data is:",tmpSlice)
}

func Test_maxPathSum(t *testing.T){
	a,b,c := 1,2,3
	slice := []*int{&a,&b,&c}
	tree := getTestTree(slice)
	maxSum:= maxPathSum(tree)
	assert.Equal(t,a+b+c,maxSum)

	a,b,c,d,e := -10,9,20,15,7
	slice = []*int{
		&a,&b,&c,nil,nil,&d,&e,
	}
	tree = getTestTree(slice)
	maxSum= maxPathSum(tree)
	assert.Equal(t,42,maxSum)

	a,b = 1,2
	slice = []*int{
		&a,&b,
	}
	tree = getTestTree(slice)
	maxSum = maxPathSum(tree)
	assert.Equal(t,3,maxSum)

	a,b = 2,-1
	slice = []*int{
		&a,&b,
	}
	tree = getTestTree(slice)
	maxSum = maxPathSum(tree)
	assert.Equal(t,2,maxSum)

	a,b = -2,1
	slice = []*int{
		&a,&b,
	}
	tree = getTestTree(slice)
	maxSum = maxPathSum(tree)
	assert.Equal(t,1,maxSum)
}
