package algorithm

import (
	"log"
	"testing"
)

func threeSum(nums []int) [][]int {
	if len(nums)<3{
		return [][]int{}
	}

	result := make([][]int,0)
	tmpLen := len(nums)
	tmpNum := make([]int,0)
	for i:=1;i<tmpLen;i++{
		for j:= i+1;j<tmpLen;j++{
			if nums[0]+nums[i]+nums[j] == 0{
				log.Println("the i j is:",i,j)
				result = append(result,[]int{nums[0],nums[i],nums[j]})
				if nums[0] == nums[i]{
					break
				}
			}
		}

		if nums[0] != nums[i]{
			tmpNum = append(tmpNum,nums[i])
		}
	}


	result = append(result,threeSum(tmpNum)...)

	return result
}

func Test_threeSum(t *testing.T){
	nums := []int{0,0,0,0}
	result:=threeSum(nums)
	log.Println("the result is:","result",result)
}
