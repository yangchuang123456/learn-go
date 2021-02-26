package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true

*/

//递归
func isValid(s string) bool {
	if s == ""{
		return true
	}
	if len(s) <2{
		return false
	}
	if len(s) == 2{
		if (s[0] == '(' && s[1] == ')') || (s[0] == '{' && s[1] == '}')||(s[0] == '[' && s[1] == ']'){
			return true
		}else{
			return false
		}
	}

	var number int
	for i:=1;i<len(s);i++{

		if s[i] == s[0]{
			number ++
		}
		if (s[0] == '(' && s[i] == ')') || (s[0] == '{' && s[i] == '}')||(s[0] == '[' && s[i] == ']'){
			if number == 0{
				if isValid(string(s[1:i]))&&isValid(string(s[i+1:])) {
					return true
				}else{
					return false
				}
			}else{
				number --
			}

		}
	}

	return false
}

//使用栈来处理，对于左括号，压栈，遇到栈顶元素对称的右括号则出栈，若处理结束，栈为空则返回true
//此问题使用栈来处理，时间和空间复杂度都为O(n),n为字符串长度
func isValid1(s string) bool{
	stack := make([]byte,0)

	for i:=0;i<len(s);i++{
		if s[i] == '(' || s[i]=='{' || s[i] == '['{
			stack = append(stack,s[i])
		} else {
			if len(stack) ==0{
				return false
			}
			if (s[i] == ')'&&stack[len(stack)-1]=='(')||(s[i] == '}'&&stack[len(stack)-1]=='{')||(s[i] == ']'&&stack[len(stack)-1]=='['){
				stack = stack[:len(stack)-1]
			}else{
				return false
			}
		}
	}

	if len(stack) == 0{
		return true
	}else{
		return false
	}
}

func Test_isValue(t *testing.T){
	s := "()"
	res:=isValid1(s)
	assert.Equal(t,true,res)

	s = "()[]{}"
	res = isValid1(s)
	assert.Equal(t,true,res)

	s = "{}{}()[]"
	res = isValid1(s)
	assert.Equal(t,true,res)

	s = "{{}}()[]"
	res = isValid1(s)
	assert.Equal(t,true,res)
}
