package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//question 1
//Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?

//answer1
//遍历n个根节点,递归计算其左右bst树个数,然后根据左右子树个数计算总bst树个数
func numTrees1(n int) int {
	total :=0
	if n == 1{
		return 1
	} else if n==0{
		return 0
	}
	for i:=0;i<n;i++{
		leftNumber := numTrees1(i)
		rightNumber := numTrees1(n-i-1)

		if leftNumber*rightNumber == 0{
			total += leftNumber+rightNumber
		} else {
			total += leftNumber*rightNumber
		}
	}

	return total
}

//answer2
//算法原理和上述递归相同,当实现上采用动态规划的思想,不用每次都递归计算子数个数,即:子树个数只计算一次即可
//假设给定n是bst个数为G(n). 设置初始值G(0)=1,G(1)=1, G(n)=∑G(i-1)*G(n-i) i等于1到n
//算法2相较于算法1,其每次计算n时,不用重复计算 numTree(1)到numTree(n-1),算法从小到大依次将计算的numberTree(i)缓存起来,并提供给后续计算使用,
//因此相对于算法1直接递归计算,效率有很大提升. 这也是动态规划相较于分治策略的不同之处
func numTrees2(n int) int{
	G := make([]int,n+1)
	G[0] =1
	G[1] =1

	for i:=2;i<=n;i++{
		for j:=1;j<=i;j++{
			G[i] += G[j-1]*G[i-j]
		}
	}

	return G[n]
}

func Test_uniqueBinarySearch(t *testing.T){
	assert.Equal(t,1,numTrees1(1))
	assert.Equal(t,2,numTrees1(2))
	assert.Equal(t,5,numTrees1(3))
	assert.Equal(t,14,numTrees1(4))

	assert.Equal(t,1,numTrees2(1))
	assert.Equal(t,2,numTrees2(2))
	assert.Equal(t,5,numTrees2(3))
	assert.Equal(t,14,numTrees2(4))
}
