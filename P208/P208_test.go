package p208

import "testing"

func TestTrie(t *testing.T) {
	t.Run("插入与搜索单词", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("apple")
		if got := trie.Search("apple"); !got {
			t.Errorf("Search(\"apple\") = %v, want true", got)
		}
		if got := trie.Search("app"); got {
			t.Errorf("Search(\"app\") = %v, want false", got)
		}
		if got := trie.StartsWith("app"); !got {
			t.Errorf("StartsWith(\"app\") = %v, want true", got)
		}
		trie.Insert("app")
		if got := trie.Search("app"); !got {
			t.Errorf("Search(\"app\") after insert = %v, want true", got)
		}
	})

	t.Run("空字符串", func(t *testing.T) {
		trie := Constructor()
		if got := trie.Search(""); got {
			t.Errorf("Search(\"\") = %v, want false", got)
		}
		if got := trie.StartsWith(""); !got {
			t.Errorf("StartsWith(\"\") = %v, want true", got)
		}
		trie.Insert("")
		if got := trie.Search(""); !got {
			t.Errorf("Search(\"\") after insert = %v, want true", got)
		}
	})

	t.Run("搜索不存在的单词", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("banana")
		if got := trie.Search("apple"); got {
			t.Errorf("Search(\"apple\") = %v, want false", got)
		}
		if got := trie.Search("ban"); got {
			t.Errorf("Search(\"ban\") = %v, want false", got)
		}
		if got := trie.Search("bananaa"); got {
			t.Errorf("Search(\"bananaa\") = %v, want false", got)
		}
	})

	t.Run("前缀匹配", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("app")
		trie.Insert("apple")
		trie.Insert("application")
		if got := trie.StartsWith("app"); !got {
			t.Errorf("StartsWith(\"app\") = %v, want true", got)
		}
		if got := trie.StartsWith("appl"); !got {
			t.Errorf("StartsWith(\"appl\") = %v, want true", got)
		}
		if got := trie.StartsWith("application"); !got {
			t.Errorf("StartsWith(\"application\") = %v, want true", got)
		}
		if got := trie.StartsWith("b"); got {
			t.Errorf("StartsWith(\"b\") = %v, want false", got)
		}
		if got := trie.StartsWith("applications"); got {
			t.Errorf("StartsWith(\"applications\") = %v, want false", got)
		}
	})

	t.Run("重复插入", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("word")
		trie.Insert("word")
		trie.Insert("word")
		if got := trie.Search("word"); !got {
			t.Errorf("Search(\"word\") after duplicate inserts = %v, want true", got)
		}
	})

	t.Run("共享前缀的单词", func(t *testing.T) {
		trie := Constructor()
		words := []string{"car", "card", "care", "careful", "cars"}
		for _, w := range words {
			trie.Insert(w)
		}
		for _, w := range words {
			if got := trie.Search(w); !got {
				t.Errorf("Search(%q) = %v, want true", w, got)
			}
		}
		notExist := []string{"ca", "caref", "carefull", "cardd", "cart"}
		for _, w := range notExist {
			if got := trie.Search(w); got {
				t.Errorf("Search(%q) = %v, want false", w, got)
			}
		}
		if got := trie.StartsWith("ca"); !got {
			t.Errorf("StartsWith(\"ca\") = %v, want true", got)
		}
		if got := trie.StartsWith("care"); !got {
			t.Errorf("StartsWith(\"care\") = %v, want true", got)
		}
		if got := trie.StartsWith("careful"); !got {
			t.Errorf("StartsWith(\"careful\") = %v, want true", got)
		}
		if got := trie.StartsWith("d"); got {
			t.Errorf("StartsWith(\"d\") = %v, want false", got)
		}
	})

	t.Run("单字符单词", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("a")
		trie.Insert("b")
		if got := trie.Search("a"); !got {
			t.Errorf("Search(\"a\") = %v, want true", got)
		}
		if got := trie.Search("b"); !got {
			t.Errorf("Search(\"b\") = %v, want true", got)
		}
		if got := trie.Search("c"); got {
			t.Errorf("Search(\"c\") = %v, want false", got)
		}
		if got := trie.StartsWith("a"); !got {
			t.Errorf("StartsWith(\"a\") = %v, want true", got)
		}
	})

	t.Run("长单词", func(t *testing.T) {
		trie := Constructor()
		long := "abcdefghijklmnopqrstuvwxyz"
		trie.Insert(long)
		if got := trie.Search(long); !got {
			t.Errorf("Search(long word) = %v, want true", got)
		}
		if got := trie.StartsWith("abcdefghijklmnopqrstuvwxy"); !got {
			t.Errorf("StartsWith(prefix of long word) = %v, want true", got)
		}
		if got := trie.Search("abcdefghijklmnopqrstuvwxy"); got {
			t.Errorf("Search(prefix without end) = %v, want false", got)
		}
	})

	t.Run("LeetCode官方示例", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("apple")
		if got := trie.Search("apple"); !got {
			t.Errorf("Search(\"apple\") = %v, want true", got)
		}
		if got := trie.Search("app"); got {
			t.Errorf("Search(\"app\") = %v, want false", got)
		}
		if got := trie.StartsWith("app"); !got {
			t.Errorf("StartsWith(\"app\") = %v, want true", got)
		}
		trie.Insert("app")
		if got := trie.Search("app"); !got {
			t.Errorf("Search(\"app\") = %v, want true", got)
		}
	})
}
