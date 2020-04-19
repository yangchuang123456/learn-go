package algorithm

import (
	"github.com/stretchr/testify/assert"
	"log"
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
func abs(a int) int{
	return (a+(a>>31))^(a>>31)
}

func subPresent(presentVec []int, target int) int{
	log.Println("the presentVec is:",presentVec)
	log.Println("the target is:",target)

	len := len(presentVec)
	if len == 2{
		subValue,sumValue := 0,0
		sumValue = presentVec[0]+presentVec[1]
		subValue = abs(presentVec[0]-presentVec[1])

		if abs(sumValue-target)<=abs(subValue-target){
			return abs(sumValue-target)
		}else{
			return abs(subValue-target)
		}
	}

	subTarget := abs(presentVec[len-1]-target)
	return subPresent(presentVec[:len-1],subTarget)
}

func maxPresent( presentVec []int ) int {
	if len(presentVec) ==0{
		return 0
	}else if len(presentVec) ==1{
		return presentVec[0]
	}

	return subPresent(presentVec,0)
}

//动态规划
func maxPresent_dynamic_planning( presentVec []int ) int {
	return 0
}

func Test_maxPresent(t *testing.T){
	presentVec := []int{1,2,3,4}
	ret := maxPresent(presentVec)
	assert.Equal(t,0,ret)

	presentVec = []int{1,3,5}
	ret = maxPresent(presentVec)
	assert.Equal(t,1,ret)

	presentVec = []int{1,2,3,6}
	ret = maxPresent(presentVec)
	assert.Equal(t,0,ret)

	presentVec = []int{41,467,334,1,169,224,478,358}
	ret = maxPresent(presentVec)
	assert.Equal(t,0,ret)
}

func Test_maxPresentDynamicPlanning(t *testing.T){
	presentVec := []int{1,2,3,4}
	ret := maxPresent_dynamic_planning(presentVec)
	assert.Equal(t,0,ret)

	presentVec = []int{1,3,5}
	ret = maxPresent_dynamic_planning(presentVec)
	assert.Equal(t,1,ret)
}
