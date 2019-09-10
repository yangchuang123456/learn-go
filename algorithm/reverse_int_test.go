package algorithm

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func reverse(x int) int {
	var	b,res int
	for ;x/10!=0 || x%10 !=0;{
		b = x % 10
		res*=10
		res += b
		x/=10
	}
	if res >int(0x7fffffff) || res <int(-0x7fffffff-1){
		return 0
	}
	return res
}

func Test_reverseInt(t *testing.T){
	a := 123
	assert.Equal(t,321,reverse(a))

	log.Println("the 2^32 -1 is:",-0x7fffffff-1)
}
