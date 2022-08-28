package arr

import (
	"fmt"
	"testing"
)

func mutate(nums [3]int) {
	nums[0] = 100
}

func TestArr(t *testing.T) {
	arr := [3]int{1,2,3}
	mutate(arr)
	fmt.Println(arr)
}
