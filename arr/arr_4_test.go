package arr

import (
	"fmt"
	"testing"
)

func mutate_4(nums *[]int) {
	*nums = append(*nums, 4)
}

func TestArr_4(t *testing.T) {
	arr := []int{1,2,3}
	mutate_4(&arr)
	fmt.Println(arr)
}
