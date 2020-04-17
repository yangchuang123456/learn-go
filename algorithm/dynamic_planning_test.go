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

//递归写法
func maxPresent( presentVec []int ) int {
	// write code here
	number := len(presentVec)
	if number >=100{
		return 0
	}else if number == 1{
		return presentVec[0]
	}

	left := maxPresent(presentVec[:number/2])
	right := maxPresent(presentVec[number/2:])

	if left < right{
		return right -left
	}else{
		return left - right
	}
}

//动态规划
func maxPresent_dynamic_planning( presentVec []int ) int {
	number := len(presentVec)
	midValue := make([]int, number)
	copy(midValue,presentVec)
	for {
		midValueLen := len(midValue)
		if midValueLen ==1{
			break
		}

		i:=0
		for i=0;i<midValueLen/2;i++{
			tmpValue:=0
			if midValue[i*2]<midValue[i*2+1]{
				tmpValue = midValue[i*2+1] - midValue[i*2]
			} else{
				tmpValue = midValue[i*2] - midValue[i*2+1]
			}
			midValue[i] = tmpValue
		}

		if midValueLen%2 ==1{
			midValue[i]=midValue[midValueLen-1]
		}
		midValue = midValue[:i+1]
	}
	return midValue[0]
}

func Test_maxPresent(t *testing.T){
	presentVec := []int{1,2,3,4}
	ret := maxPresent_dynamic_planning(presentVec)
	assert.Equal(t,0,ret)

	presentVec = []int{1,3,5}
	ret = maxPresent_dynamic_planning(presentVec)
	assert.Equal(t,1,ret)
}
