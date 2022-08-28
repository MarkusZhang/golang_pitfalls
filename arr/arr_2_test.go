package arr

import (
	"fmt"
	"testing"
)

func mutate_2(nums []int) {
	nums[0] = 100
}

func TestArr_2(t *testing.T) {
	arr := []int{1,2,3}
	mutate_2(arr)
	fmt.Println(arr)
}

