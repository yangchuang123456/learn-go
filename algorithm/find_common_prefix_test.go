package algorithm
/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.

*/

//递归
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	if len(strs) == 1{
		return strs[0]
	}
	if len(strs) == 2{
		var tmpLen int
		s:=make([]byte,0)
		len1 := len([]byte(strs[0]))
		len2 := len([]byte(strs[1]))
		if len1 < len2{
			tmpLen = len1
		}else{
			tmpLen = len2
		}

		for i:=0;i<tmpLen;i++{
			if strs[0][i] == strs[1][i]{
				s = append(s,strs[0][i])
			}else{
				break
			}
		}
		return string(s)
	}

	tmpStrs := make([]string,0)
	tmpStrs = append(tmpStrs,strs[0],longestCommonPrefix(strs[1:]))

	return longestCommonPrefix(tmpStrs)
}

