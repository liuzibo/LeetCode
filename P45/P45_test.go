package P45

import (
	"testing"
)

func TestSolution(t *testing.T) {
	input := []int{2, 3, 1, 1, 4}
	output := 2
	if jump(input) != output {
		t.Errorf("Solution(%v) = %v; want %v", input, jump(input), output)
	}
}
