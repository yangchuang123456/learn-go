package algorithm

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

/*
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"
Example 2:

Input: num1 = "123", num2 = "456"
Output: "56088"
Note:

The length of both num1 and num2 is < 110.
Both num1 and num2 contain only digits 0-9.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
You must not use any built-in BigInteger library or convert the inputs to integer directly.

*/

func multiply(num1 string, num2 string) string {
	num1Len := len(num1)
	num2Len := len(num2)
	result := make([]byte,num1Len+num2Len)


	var carry int
	for i:=num1Len-1;i>=0;i--{
		for j:=num2Len-1;j>=0;j--{
			product:=int((num1[i]-0x30)*(num2[j]-0x30))
			carry = product/10
			result[i+j+1] += byte(product%10)
			log.Println("the i j is:",i,j,i+j)
			result[i+j] += byte(carry)
			if result[i+j+1] >= 10{
				result[i+j]+= result[i+j+1]/10
				result[i+j+1] %=10
			}

			log.Println("the result is:",result)
		}
	}

	var index int
	flag := false
	for i,v := range result{
		if v != 0 && !flag{
			index = i
			flag = true
		}
		result[i] +=0x30
	}

	if !flag{
		return "0"
	}

	result = result[index:]

	return string(result)
}

func Test_multiply(t *testing.T){
	num1 := "2"
	num2 := "3"
	result := multiply(num1,num2)
	assert.Equal(t,"6",result)

	num1 = "498828660196"
	num2 = "840477629533"
	result = multiply(num1,num2)
	assert.Equal(t,"419254329864656431168468",result)

	num1 = "123"
	num2 = "456"
	result = multiply(num1,num2)
	assert.Equal(t,"56088",result)
}
