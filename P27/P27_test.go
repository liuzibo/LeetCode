package P27

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	res := removeElement(nums, 2)
	fmt.Println(res)
}
