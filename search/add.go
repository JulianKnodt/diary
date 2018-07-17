package search

type Trie struct {
	Value byte
	Keys  []string
	Left  Trie
	Right Trie
	Mid   Trie
}

func (t *Trie) Add(data []byte, key string) {

}

func Add(data []bytes) {
	// load Trie
}
