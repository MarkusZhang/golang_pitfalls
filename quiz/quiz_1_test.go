package quiz

import (
	"fmt"
	"testing"
)

func TestQuiz_1(t *testing.T) {
	nums := make([][]int, 4)
	nums[0] = []int{1}
	nums[1] = append(nums[0],2) // 1, 2
	nums[2] = append(nums[1],3) // 1, 2, 3
	nums[3] = append(nums[2],4) // 1, 2, 3, 4

	nums[1][0] = 100
	nums[2][0] = 200
	nums[3][0] = 300

	for i := range nums {
		fmt.Println(i,nums[i])
	}
}
