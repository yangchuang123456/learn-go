package go_math

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_Log(t *testing.T){
	a := 2<<2
	assert.Equal(t,8,a)

	c := math.Log2(float64(a))
	assert.Equal(t,3.0,c)

	a = 9
	c = math.Ceil(math.Log2(float64(a)))
	assert.Equal(t,4.0,c)
}
