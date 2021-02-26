package niuke

import "fmt"

func main(){
	var number int
	fmt.Scanln(&number)
	height := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scanln(&height[i])
	}

	res := make([]int, number)
	for i := 0; i < number; i++ {
		left := height[:i]
		right := height[i+1:]

		leftValue, rightValue := 0, 0
		leftFlag := 0
		if len(left) != 0 {
			leftFlag = left[len(left)-1]
			leftValue = 1
			for j := len(left) - 1; j >= 0; j-- {
				if left[j] > leftFlag {
					leftValue++
					leftFlag = left[j]
				}
			}
		}

		rightFlag := 0
		if len(right) != 0 {
			rightFlag = right[0]
			rightValue = 1
			for j := 0; j < len(right); j++ {
				if right[j] > rightFlag {
					rightValue++
					rightFlag = right[j]
				}
			}
		}

		res[i] = leftValue + rightValue + 1
	}

	fmt.Println(res)
}
