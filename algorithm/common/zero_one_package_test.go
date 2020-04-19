package common

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
0-1 背包问题
*/

//返回物品價值和選擇方案
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
		//超出背包容量，直接選擇下一個物品
		if quality[i] > c{
			choice[i]=0
			return zeroOnePackage(price[:i],quality[:i],c)
		}

		//未超出背包容量，比較放不放那個更優
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

func Test_package(t *testing.T){
	price := []int{3,4,5,6}
	quality := []int{2,3,4,5}
	c := 8
	ret,choice,err := zeroOnePackage(price,quality,c)
	assert.NoError(t,err)
	assert.Equal(t,10,ret)
	assert.Equal(t,[]int{0,1,0,1},choice)
}