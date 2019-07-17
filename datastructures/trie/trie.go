package main

import (
	"fmt"
	"sync"
)

/*
In memory data source
O(1) time complexity
*/

/*
Node that points to the hashmap
*/
type Node struct {
	left  *Node
	right *Node
	key   int
	data  *DataBlock
}

/*
Trie implementation as a map
*/
type Trie struct {
	t     map[int]*Node
	tLock map[int]*sync.RWMutex // trie level lock
}

func (t *Trie) getKey(data string) int {
	print(data)
	i := data[0]
	return int(i)
}

func (t *Trie) createNode(key int) *Node {
	n := new(Node)
	n.left = nil
	n.right = nil
	n.key = key
	n.data = new(DataBlock)
	return n
}

func releaseLock(l *sync.RWMutex) {
	fmt.Println("releasing lock")
	l.Unlock()
}

func releaseReadLock(l *sync.RWMutex) {
	fmt.Println("releasing read lock")
	l.RUnlock()
}

func acquireLock(l *sync.RWMutex) {
	fmt.Println("acquiring lock")
	l.Lock()
}

func acquireReadLock(l *sync.RWMutex) {
	fmt.Println("acquiring read lock")
	l.RLock()
}

/*
function to put data into the trie
*/
func (t *Trie) put(key string, data string, resp chan interface{}) {
	tKey := t.getKey(key)
	if t.tLock[tKey] == nil {
		fmt.Println("root node found nil")
		t.tLock[tKey] = new(sync.RWMutex)
		acquireLock(t.tLock[tKey])
		t.t[tKey] = t.createNode(tKey)
	} else {
		acquireLock(t.tLock[tKey])
	}
	currentNode := t.t[tKey]
	for i := 0; i < len(key); i++ {
		c := int(key[i])
		fmt.Println(c)
		if c < currentNode.key {
			if currentNode.left != nil {
				// fmt.Println("left node not nil")
				currentNode = currentNode.left
				continue
			}
			fmt.Println("creating left node")
			currentNode.left = t.createNode(c)
			currentNode = currentNode.left
		} else {
			if currentNode.right != nil {
				// fmt.Println("right node not nil")
				currentNode = currentNode.right
				continue
			}
			fmt.Println("creating right node")
			currentNode.right = t.createNode(c)
			currentNode = currentNode.right
		}
	}
	currentNode.data.dataPointer = &data
	releaseLock(t.tLock[tKey])
	fmt.Printf("Record %s written\n", data)
	resp <- nil
}

/*
function to get data from the trie
*/
func (t *Trie) get(key string, resp chan interface{}) {
	tKey := t.getKey(key)

	if t.tLock[tKey] == nil {
		fmt.Println("root node found nil")
		resp <- nil
	}
	acquireReadLock(t.tLock[tKey])
	currentNode := t.t[tKey]
	for i := 0; i < len(key); i++ {
		c := int(key[i])
		fmt.Println(c)
		if int(c) < currentNode.key {

			if currentNode.left == nil {
				fmt.Println("left node nil")
				releaseReadLock(t.tLock[tKey])
				resp <- nil
			}
			fmt.Println("left node not nil")
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				fmt.Println("right node nil")
				releaseReadLock(t.tLock[tKey])
				resp <- nil
			}
			fmt.Println("right node not nil")
			currentNode = currentNode.right
		}
	}
	if currentNode.data.dataPointer != nil {
		d := *currentNode.data.dataPointer
		releaseReadLock(t.tLock[tKey])
		resp <- d
	} else {
		releaseReadLock(t.tLock[tKey])
	}
	resp <- nil
}
