package sol

type Node struct {
	Children  map[byte]*Node
	EndOfWord bool
}
type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{
			Children:  make(map[byte]*Node),
			EndOfWord: false,
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	wLen := len(word)
	for idx := 0; idx < wLen; idx++ {
		r := word[idx]
		if _, exists := cur.Children[r]; !exists {
			cur.Children[r] = &Node{
				Children:  make(map[byte]*Node),
				EndOfWord: false,
			}
		}
		cur = cur.Children[r]
	}
	cur.EndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return DFS(0, this.root, &word)
}

func DFS(i int, root *Node, word *string) bool {
	cur := root
	wLen := len(*word)
	if wLen == i {
		return cur.EndOfWord
	}
	c := (*word)[i]
	if c == '.' {
		for _, child := range cur.Children {
			if DFS(i+1, child, word) { // check next char
				return true
			}
		}
		return false
	} else {
		if _, exists := cur.Children[c]; !exists {
			return false
		}
		return DFS(i+1, cur.Children[c], word)
	}
}
