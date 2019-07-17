package main

import "sync"

type CacheBackend interface {
	getValue(key string, resp chan interface{})
	putValue(key string, data string, resp chan interface{})
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
	t.trie.tLock = make(map[int]*sync.RWMutex)
}

func (t *TrieBackend) getValue(key string, resp chan interface{}) {
	go t.trie.get(key, resp)
}

func (t *TrieBackend) putValue(key string, data string, resp chan interface{}) {
	go t.trie.put(key, data, resp)
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
