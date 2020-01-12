package kit

import (
	"fmt"
	"testing"
)

func Test_Trie(t *testing.T) {
	src := []string{
		"apple",
		"app",
		"apue",
		"duck",
		"application",
	}
	root := buildTrie(src)
	s := "app"
	if succ, res := searchTrie(root, s); succ {
		fmt.Println(res)
	} else {
		fmt.Println("No results found")
	}
}
