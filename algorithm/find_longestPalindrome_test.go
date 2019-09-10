package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
question:
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/

/*
使用动态规划的方式计算最长回文子串
即:如果用f[i][j] 保存子串从i 到j是否是回文子串，那么在求f[i][j] 的时候如果j-i>=2时，如果 f[i][j] 为回文，那么f[i+1][j-1],也一定为回文。否则f[i][j]不为回文
此算法时间复杂度比较高为O(n^2)
*/

func findLongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	stringLen := len(s)
	flag := make(map[[2]int]bool)
	var maxLenIndex [2]int
	for j := 0; j < stringLen; j++ {
		flag[[2]int{j, j}] = true

		for i := 0; i < j; i++ {
			if j-i == 1 {
				flag[[2]int{i, j}] = s[i] == s[j]
			} else {
				flag[[2]int{i, j}] = s[i] == s[j] && flag[[2]int{i + 1, j - 1}]
			}

			if flag[[2]int{i, j}] {
				if maxLenIndex[1]-maxLenIndex[0] < j-i {
					maxLenIndex[0] = i
					maxLenIndex[1] = j
				}
			}
		}
	}

	return s[maxLenIndex[0] : maxLenIndex[1]+1]
}

/*
使用马拉车算法将将复杂度降低到线性
*/

func findLongestPalindrome2(s string) string {
	//填充#扩展s长度为2n+1,这样不论s为奇数还是偶数,不用分开讨论
	paddingS := "#"
	for _, char := range s {
		paddingS = paddingS + string(char) + "#"
	}

	//log.Println("the paddingS is:", "paddingS", paddingS)
	lenS := len(paddingS)
	//log.Println("the paddingS len is:",lenS)

	var id, mx, maxLen, maxLenIndex int
	p := make([]int, lenS)

	for i := 0; i < lenS; i++ {
		//根据i在mx左边还是右边判断初值是否可以利用对称点计算
		if i < mx {
			p[i] = p[2*id-i]
		} else {
			p[i] = 1
		}

	//	log.Println("the i and p[i] is:",i,p[i])
		for ; i-p[i] >= 0 && i+p[i] < lenS && paddingS[i-p[i]] == paddingS[i+p[i]]; {
			p[i] ++
		}

		if mx < i+p[i] {
			id = i
			mx = p[i]
		}

		if maxLen < p[i] {
			maxLenIndex = i
			maxLen = p[i]
		}
	}

	findS := paddingS[maxLenIndex-maxLen+1 : maxLenIndex+maxLen]
	var rS string
	for _, v := range findS {
		if string(v) != "#"{
			rS+=string(v)
		}
	}

	return rS
}

func Test_findLongestPalindrome(t *testing.T) {
	s := "abcdefg"
	assert.Equal(t, "a", findLongestPalindrome2(s))

	s = "abacdefg"
	assert.Equal(t, "aba", findLongestPalindrome2(s))

	s = "abccdefg"
	assert.Equal(t, "cc", findLongestPalindrome2(s))

	s = "abcddcdefg"
	assert.Equal(t, "cddc", findLongestPalindrome2(s))

	s = "babad"
	assert.Equal(t, "bab", findLongestPalindrome2(s))

	s = ""
	assert.Equal(t, "", findLongestPalindrome2(s))

	s = "1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	assert.Equal(t, s, findLongestPalindrome2(s))
}
