package main

type CacheBackend interface {
	getValue(key string) (string, error)
	putValue(key string, data string) error
	init()
}

type TrieBackend struct {
	trie *Trie
}

type ListBackend struct {
	list *List
}

func (t *TrieBackend) init() {
	t.trie = new(Trie)
	t.trie.t = make(map[int]*Node)
}

func (t *TrieBackend) getValue(key string) (string, error) {
	return t.trie.get(key)
}

func (t *TrieBackend) putValue(key string, data string) error {
	return t.trie.put(key, data)
}

func (l *ListBackend) init() {
	l.list = new(List)
}

func (l *ListBackend) getValue(key string) (string, error) {
	return l.list.get(key)
}

func (l *ListBackend) putValue(key string, data string) error {
	return l.list.put(key, data)
}
