package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
众所周知，牛能和牛可乐经常收到小粉丝们送来的礼物，每个礼物有特定的价值，他俩想要尽可能按照自己所得价值来平均分配所有礼物。

那么问题来了，在最优的情况下，他俩手中得到的礼物价值和的最小差值是多少呢？
p.s 礼物都很珍贵，所以不可以拆开算哦

示例1
输入
[1,2,3,4]
输出
0
说明
他俩一个人拿1,4 。另一个人拿2,3
*/

func abs(a int) int {
	return (a + (a >> 31)) ^ (a >> 31)
}

//动态规划
//此题为变形版动态规划,以总礼物质量的一半为背包容量,计算背包问题的结果分配给一个人,剩余结果分配给另一个人
func maxPresent_dynamic_planning(presentVec []int) int {

	totalPrice := 0
	for _, v := range presentVec {
		totalPrice += v
	}

	c := totalPrice / 2
	number := len(presentVec)
	quality := make([]int,number)
	copy(quality,presentVec)
	dp := make([]int,c+1)

	for i:=1;i<=number;i++{
		for j:=c;j>=1;j--{
			if j< quality[i-1]{
				continue
			}

			if dp[j] < dp[j-quality[i-1]]+presentVec[i-1]{
				dp[j] = dp[j-quality[i-1]]+presentVec[i-1]
			}
		}
	}

	res := dp[c]

	return abs(totalPrice-2*res)
}

func Test_maxPresentDynamicPlanning(t *testing.T) {
	presentVec := []int{1, 2, 3, 4}
	ret := maxPresent_dynamic_planning(presentVec)
	assert.Equal(t, 0, ret)

	presentVec = []int{1, 3, 5}
	ret = maxPresent_dynamic_planning(presentVec)
	assert.Equal(t, 1, ret)

	presentVec = []int{1, 2, 3, 6}
	ret = maxPresent_dynamic_planning(presentVec)
	assert.Equal(t, 0, ret)

	presentVec = []int{41, 467, 334, 1, 169, 224, 478, 358}
	ret = maxPresent_dynamic_planning(presentVec)
	assert.Equal(t, 0, ret)
}
