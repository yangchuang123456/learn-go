package algorithm

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
0-1 背包问题
*/

//返回物品价值和选择方案,递归实现
func zeroOnePackage(price,quality []int,c int) (int,[]int,error){
	if len(price) != len(quality){
		return 0,nil,errors.New("the input price and quality len error")
	}
	number := len(price)
	if number ==0 || c==0{
		return 0,[]int{},nil
	}
	choice := make([]int,number)

	for i:=number-1;i>=0;i--{
		//超出背包容量，直接选择下一個物品
		if quality[i] > c{
			choice[i]=0
			return zeroOnePackage(price[:i],quality[:i],c)
		}

		//未超出背包容量，比较放不放那个更优
		value1,choice1,_ := zeroOnePackage(price[:i],quality[:i],c-quality[i])
		value1+=price[i]
		value2,choice2,_ := zeroOnePackage(price[:i],quality[:i],c)

		if value1 > value2{
			copy(choice,choice1)
			choice[i]=1
			return value1,choice,nil
		}else{
			copy(choice,choice2)
			choice[i]=0
			return value2,choice,nil
		}
	}

	return 0, nil, nil
}

//二维数组动态规划求解背包问题
func zeroOnePackage_dp1(price,quality []int,c int) (int,[]int,error){
	if len(price) != len(quality){
		return 0,nil,errors.New("the input price and quality len error")
	}

	//init package table
	number := len(price)
	packageTable := make([][]int, number+1)
	for i,_ := range  packageTable{
		packageTable[i] = make([]int,c+1)
	}

	for i:=1;i<=number;i++{
		for j:=1;j<=c;j++{
			if quality[i-1] > j{
				packageTable[i][j] = packageTable[i-1][j]
				continue
			}
			if packageTable[i-1][j] < packageTable[i-1][j-quality[i-1]]+price[i-1]{
				packageTable[i][j] = packageTable[i-1][j-quality[i-1]]+price[i-1]
			}else {
				packageTable[i][j] = packageTable[i-1][j]
			}
		}
	}
	return packageTable[number][c],nil,nil
}

//对于packageTable[i][j] 其值取决于 packageTable[i-1][x](0=<x<j) 因此可以使用一维数组进一步压缩空间使用情况.
//由于使用一维数组, 因此内层capacity遍历需要从后往前,因为若从前往后的话,上一层的数据被使用时有可能已被下一层更新了
func zeroOnePackage_dp2(price,quality []int,c int) (int,[]int,error){
	if len(price) != len(quality){
		return 0,nil,errors.New("the input price and quality len error")
	}

	number := len(price)
	dp := make([]int,c+1)
	for i:=1;i<=number;i++{
		for j:=c;j>=1;j--{
			if quality[i-1] > j{
				continue
			}

			if dp[j] < dp[j-quality[i-1]]+price[i-1]{
				dp[j] = dp[j-quality[i-1]]+price[i-1]
			}
		}
	}
	return dp[c],nil,nil
}


func Test_package(t *testing.T){
	price := []int{3,4,5,6}
	quality := []int{2,3,4,5}
	c := 8
	ret,choice,err := zeroOnePackage(price,quality,c)
	assert.NoError(t,err)
	assert.Equal(t,10,ret)
	assert.Equal(t,[]int{0,1,0,1},choice)

	ret,_,err = zeroOnePackage_dp1(price,quality,c)
	assert.NoError(t,err)
	assert.Equal(t,10,ret)

	ret,_,err = zeroOnePackage_dp2(price,quality,c)
	assert.NoError(t,err)
	assert.Equal(t,10,ret)
}