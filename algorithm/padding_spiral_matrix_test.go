package algorithm

import (
	"log"
	"testing"
)

/*
Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

Example:

Input: 3
Output:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]

*/

func generateMatrix(n int) [][]int {
	res := make([][]int,n)
	for i := range res{
		res[i] = make([]int,n)
	}
	t,l :=0,0
	r,b := n-1,n-1
	for i:=1;i<=n*n;{
		for j:=l;j<=r;j++{
			res[t][j] = i
			i++
		}
		t++

		for j:=t;j<=b;j++{
			res[j][r] = i
			i++
		}
		r--

		for j:=r;j>=l;j--{
			res[b][j] = i
			i++
		}
		b--

		for j:=b;j>=t;j--{
			res[j][l] = i
			i++
		}
		l++
	}
	return res
}


func Test_generateMatrix(t *testing.T){
	n := 4
	ret := generateMatrix(n)
	log.Println("the res is:",ret)
}