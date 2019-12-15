package algorithm

import (
	"log"
	"testing"
)

/*
The gray code is a binary numeral system where two successive values differ in only one bit.

Given a non-negative integer n representing the total number of bits in the code, print the sequence of gray code. A gray code sequence must begin with 0.

Example 1:

Input: 2
Output: [0,1,3,2]
Explanation:
00 - 0
01 - 1
11 - 3
10 - 2

For a given n, a gray code sequence may not be uniquely defined.
For example, [0,2,3,1] is also a valid gray code sequence.

00 - 0
10 - 2
11 - 3
01 - 1
Example 2:

Input: 0
Output: [0]
Explanation: We define the gray code sequence to begin with 0.
             A gray code sequence of n has size = 2^n, which for n = 0 the size is 2^0 = 1.
             Therefore, for n = 0 the gray code sequence is [0].
*/

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	if n <0{
		return nil
	}

	sliceNumber := 1 << uint(n)
	ret := make([]int,sliceNumber)

	for i:=1;i<=n;i++{
		if i==1 {
			ret[0] = 0
			ret[1] = 1
			continue
		}

		lastNumber := 1 << uint(i-1)
		tmpSlice := ret[:lastNumber]
		for j:=0;j<lastNumber;j++{
			//上一个列表逆序并高位补１，然后和之前的数组拼接即为i+1阶的格雷码
			ret[lastNumber+j]=tmpSlice[lastNumber-j-1]+lastNumber
		}
	}

	return ret
}

func Test_grayCode(t *testing.T){
	n :=3
	ret := grayCode(n)
	log.Println("the ret is:",ret)
}
