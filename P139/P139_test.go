package P139

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	result := wordBreak(s, wordDict)
	if result != true {
		t.Errorf("wordBreak(%s, %v) = %v, want %v", s, wordDict, result, true)
	}
}
