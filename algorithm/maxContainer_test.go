package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

 
Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49

*/

//暴力破解，时间复杂度　O(n^2)
func maxArea(height []int) int {
	tmpLen := len(height)
	if tmpLen < 2 {
		return 0
	}

	maxContainer := 0
	for i := range height {
		for j := range height[i+1:] {
			var tmpValue int
			if height[i] < height[j+i+1] {
				tmpValue = height[i]
			} else {
				tmpValue = height[j+i+1]
			}

			//log.Println("the i j value is: ",i,j+i+1,height[i],height[j+i+1])
			if maxContainer < (j+1)*tmpValue {
				maxContainer = (j + 1) * tmpValue
			}
		}
	}

	return maxContainer
}

/*
双指针法，时间复杂度　O(n)

这种方法背后的思路在于，两线段之间形成的区域总是会受到其中较短那条长度的限制。此外，两线段距离越远，得到的面积就越大。

我们在由线段长度构成的数组中使用两个指针，一个放在开始，一个置于末尾。 此外，我们会使用变量 maxareamaxarea 来持续存储到目前为止所获得的最大面积。 在每一步中，我们会找出指针所指向的两条线段形成的区域，更新 maxareamaxarea，并将指向较短线段的指针向较长线段那端移动一步。

*/

func maxArea2(height []int) int {
	if len(height) < 2{
		return 0
	}
	var max,length,width int
	l := 0
	r := len(height) - 1
	for {
		if l==r{
			break
		}

		width = r -l
		if height[l] < height[r]{
			length = height[l]
			l++
		}else{
			length = height[r]
			r--
		}

		if max < length*width{
			max = length*width
		}
	}

	return max
}

func Test_maxArea(t *testing.T) {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea2(a)
	assert.Equal(t, 49, res)
}
