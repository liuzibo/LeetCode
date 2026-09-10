package P416

import (
	"testing"
)

func TestCanPartition(t *testing.T) {

	input := []int{1, 5, 11, 5}
	expected := true
	result := canPartition(input)
	if result != expected {
		t.Errorf("canPartition(%v) = %v; want %v", input, result, expected)
	}
}
