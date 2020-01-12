package kit

type trie struct {
	c byte
	son []*trie
	str string
}

func buildTrie(src []string) *trie {
	var root *trie
	root = new(trie)
	root.son = make([]*trie, 26)
	for _, s := range src {
		cur := root
		for _, c := range []byte(s) {
			if cur.son[c-'a'] == nil {
				cur.son[c-'a'] = &trie{c, nil, ""}
				cur.son[c-'a'].son = make([]*trie, 26)
			}
			cur = cur.son[c-'a']
		}
		cur.str = s
	}
	return root
}

func searchTrie(root *trie, target string) (bool, []string) {
	cur := root
	for _, c := range []byte(target) {
		if cur.son[c-'a'] == nil {
			return false, nil
		}
		cur = cur.son[c-'a']
	}
	return true, dfsTrie(cur)
}
func dfsTrie(root *trie) []string {
	var res []string
	if root.str != "" {
		res = append(res, root.str)
	}
	for _, s := range root.son {
		if s != nil {
			res = append(res, dfsTrie(s)...)
		}
	}
	return res
}
