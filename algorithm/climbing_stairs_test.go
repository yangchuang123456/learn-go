package algorithm

import (
	"log"
	"math/big"
	"testing"
)

/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

*/


func Cnm(m,n int)int{
	if n == 0 {
		return 0
	}

	if m == 0 {
		return 1
	}

	if n == m{
		return 1
	}

	if n < m {
		return 0
	}


	var tmpA,tmpB int
	if n-m < m{
		tmpA = m
		tmpB = n-m
	}else{
		tmpA = n-m
		tmpB = m
	}

	mulValue1 :=big.NewInt(1)
	mulValue2 :=big.NewInt(1)
	for i:=tmpA+1;i<=n;i++{
		log.Println("the i from tmpA+1 to n is:",i)
		mulValue1.Mul(mulValue1,big.NewInt(int64(i)))
		log.Printf("the mulValue1 is:%x\r\n",mulValue1)
	}

	for i:=1;i<=tmpB;i++{
		log.Println("the i from 1 to tmpB is:",i)
		mulValue2.Mul(mulValue2,big.NewInt(int64(i)))
	}

	log.Println("the n m is:",n,m)
	log.Println("the mulValue1 and mulValue2 is:",mulValue1,mulValue2)

	return int(mulValue1.Div(mulValue1,mulValue2).Int64())
}

func climbStairs(n int) int {
	var ret int

	for i:=0;i<=n/2;i++{
		//log.Println("the i n-1 and Cnm(i,n-i) is:",i,n-i,Cnm(i,n-i))
		ret += Cnm(i,n-i)
	}

	return ret
}

func Test_climbStairs(t *testing.T){
	n := 45
	ret := climbStairs(n)
	log.Println("the ret is:",ret)
}

func Test_Cnm(t *testing.T){
	n :=31
	m :=14
	ret := Cnm(m,n)
	log.Println("the ret is:",ret)
}