package P763

import (
	"testing"
)

func TestParitionLabels(t *testing.T) {
	input := "ababcbacadefegdehijhklij"
	output := partitionLabels(input)
	t.Log(output)
}
