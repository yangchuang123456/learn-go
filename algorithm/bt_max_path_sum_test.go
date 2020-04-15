package algorithm

import (
	"github.com/stretchr/testify/assert"
	"github.com/yangchuang123456/learn-go/algorithm/common"
	"log"
	"testing"
)

/*
Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting nodes to any nodes in the tree along the parent-child connections. The path must contain at least one nodes and does not need to go through the root.

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



func MaxLeftRightPathMaxSum(root *common.TreeNode) (leftMax ,rightMax int){
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


func maxPathSumFunc(root *common.TreeNode) (int,bool) {
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

//方法一: 分别计算过根节点的最大路径和左右子树的最大路径，然后在三个值中取最大值即可。计算根节点的最大路径和计算左右子树的最大路径都可以递归求得
func maxPathSum1(root *common.TreeNode) int{
	ret,_:=maxPathSumFunc(root)
	return ret
}

//方法二: 递归计算树的顶点到叶子节点的最大路径和，在递归调用过程中一直计算sum=根节点+左子树顶点到叶子的最大路径+右子树顶点到叶子的最大路径.递归完成后最大sum即为最大路径和
func subFunc(root *common.TreeNode,maxSum *int) int{
	if root == nil{
		return 0
	}
	left := subFunc(root.Left,maxSum)
	right := subFunc(root.Right,maxSum)

	tmp := root.Val
	if left > 0{
		tmp += left
	}
	if right >0{
		tmp += right
	}
	if tmp > *maxSum{
		*maxSum = tmp
	}

	if left < right{
		if right>0{
			return root.Val+right
		}

	}else{
		if left >0{
			return  root.Val+left
		}
	}

	return root.Val
}

func maxPathSum2(root *common.TreeNode) int{
	if root == nil{
		return 0
	}
	res := root.Val
	subFunc(root,&res)
	return res
}



func inOrderTraverseTree(tree *common.TreeNode,slice *[]int){
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
	tree := common.GetTestTree(slice)
	tmpSlice := make([]int,0)
	inOrderTraverseTree(tree,&tmpSlice)
	assert.Equal(t,[]int{9,-10,15,20,7},tmpSlice)
	log.Println("the tree data is:",tmpSlice)

	a,b,c = 1,2,3
	slice = []*int{&a,&b,&c}
	tree = common.GetTestTree(slice)
	tmpSlice = make([]int,0)
	inOrderTraverseTree(tree,&tmpSlice)
	assert.Equal(t,[]int{2,1,3},tmpSlice)
	log.Println("the tree data is:",tmpSlice)
}

func Test_maxPathSum(t *testing.T){
	a,b,c := 1,2,3
	slice := []*int{&a,&b,&c}
	tree := common.GetTestTree(slice)
	maxSum:= maxPathSum2(tree)
	assert.Equal(t,a+b+c,maxSum)

	a,b,c,d,e := -10,9,20,15,7
	slice = []*int{
		&a,&b,&c,nil,nil,&d,&e,
	}
	tree = common.GetTestTree(slice)
	maxSum= maxPathSum2(tree)
	assert.Equal(t,42,maxSum)

	a,b = 1,2
	slice = []*int{
		&a,&b,
	}
	tree = common.GetTestTree(slice)
	maxSum = maxPathSum2(tree)
	assert.Equal(t,3,maxSum)

	a,b = 2,-1
	slice = []*int{
		&a,&b,
	}
	tree = common.GetTestTree(slice)
	maxSum = maxPathSum2(tree)
	assert.Equal(t,2,maxSum)

	a,b = -2,1
	slice = []*int{
		&a,&b,
	}
	tree = common.GetTestTree(slice)
	maxSum = maxPathSum2(tree)
	assert.Equal(t,1,maxSum)
}
