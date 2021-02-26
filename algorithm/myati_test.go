package algorithm

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)
const Int32MAX = int(^uint32(0) >> 1)
const Int32MIN = ^Int32MAX

func myAtoi(str string) int{
	flag := false
	negFlag := false
	overflowFlag := false
	var tmp int
	for _,v := range str{
		v1 := string(v)
		if !flag && (v1!=" " && v1!= "+" && v1!="-" && (v1<"0" || v1>"9")){
			break
		}else if !flag && (v1== "+" || v1=="-" || (v1>="0"&&v1<="9")){
			if v1 == "-"{
				negFlag = true
			}

			if v1>="0"&&v1<="9" {
				tmp = int(v) - 0x30
			}

			flag = true
			log.Println("the flag and tmp is1: ",flag,tmp)
		}else if flag && (v1<"0" || v1>"9"){
			break
		}else if flag && (v1 >="0" && v1<="9"){
			log.Println("the flag and tmp is2: ",flag,tmp)
			tmp *=10
			if tmp < 0{
				overflowFlag = true
				break
			}
			log.Println("the flag and tmp is3: ",flag,tmp)
			tmp += int(v) -0x30
			log.Println("the flag and tmp is4: ",flag,tmp)
		}

		if flag && tmp < 0{
			overflowFlag = true
		}
	}

	if negFlag{
		tmp = ^tmp + 1
	}

	log.Println("the overflowFlag and negFlag is:",overflowFlag,negFlag)

	if overflowFlag && !negFlag{
		return Int32MAX
	}else if overflowFlag && negFlag{
		return Int32MIN
	}

	if tmp > Int32MAX {
		return Int32MAX
	}else if tmp < Int32MIN {
		return Int32MIN
	}

	return tmp
}

func Test_myATOI(t *testing.T){
	s := "42"
	value := myAtoi(s)
	assert.Equal(t,42,value)

	s = "+42"
	value = myAtoi(s)
	assert.Equal(t,42,value)

	s = "     -42"
	value = myAtoi(s)
	assert.Equal(t,-42,value)

	s = "     -4wre2"
	value = myAtoi(s)
	assert.Equal(t,-4,value)

	s = "-91283472332"
	value = myAtoi(s)
	assert.Equal(t,-2147483648,value)

	s = "-9223372036854775809"
	value = myAtoi(s)
	assert.Equal(t,-2147483648,value)

	s = "18446744073709551617"
	value = myAtoi(s)
	assert.Equal(t,2147483647,value)
}
