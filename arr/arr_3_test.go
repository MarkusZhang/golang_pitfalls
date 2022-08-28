package arr

import (
	"fmt"
	"testing"
)

func mutate_3(nums []int) {
	nums = append(nums, 4)
}

func TestArr_3(t *testing.T) {
	arr := []int{1,2,3}
	mutate_3(arr)
	fmt.Println(arr)
}

