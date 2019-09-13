package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxArea(height []int) int {
	tmpLen := len(height)
	if tmpLen <2 {
		return 0
	}

	maxContainer := 0
	for i:= range height{
		for j:= range height[i+1:]{
			var tmpValue int
			if height[i] < height[j+i+1]{
				tmpValue = height[i]
			}else{
				tmpValue = height[j+i+1]
			}

			//log.Println("the i j value is: ",i,j+i+1,height[i],height[j+i+1])
			if maxContainer < (j+1)*tmpValue{
				maxContainer = (j+1)*tmpValue
			}
		}
	}

	return maxContainer
}

func Test_maxArea(t *testing.T)  {
	a := []int{1,8,6,2,5,4,8,3,7}
	res := maxArea(a)
	assert.Equal(t,49,res)
}