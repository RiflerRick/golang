package main

import (
	"errors"
	"fmt"
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
	t map[int]*Node
}

func (t *Trie) getKey(data string) int {
	i := data[0]
	return int(i)
}

func (t *Trie) createNode() *Node {
	n := new(Node)
	n.left = nil
	n.right = nil
	n.key = -1
	n.data = new(DataBlock)
	return n
}

/*
function to put data into the trie
*/
func (t *Trie) put(key string, data string) error {
	tKey := t.getKey(key)
	if t.t[tKey] == nil {
		t.t[tKey] = t.createNode()
	}
	currentNode := t.t[tKey]
	for c := range key {
		if int(c) < currentNode.key {
			if currentNode.left != nil {
				currentNode = currentNode.left
				continue
			}
			currentNode.left = t.createNode()
		} else {
			if currentNode.right != nil {
				currentNode = currentNode.right
				continue
			}
			currentNode.right = t.createNode()
		}
	}
	currentNode.data.dataPointer = &data
	fmt.Printf("Record %s written\n", data)
	return nil
}

/*
function to get data from the trie
*/
func (t *Trie) get(key string) (string, error) {
	tKey := t.getKey(key)
	if t.t[tKey] == nil {
		return "", errors.New("Record not found")
	}
	currentNode := t.t[tKey]
	for c := range key {
		if int(c) < currentNode.key {
			if currentNode.left == nil {
				return "", errors.New("Record not found")
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				return "", errors.New("Record not found")
			}
			currentNode = currentNode.right
		}
	}
	if currentNode.data.dataPointer != nil {
		return *currentNode.data.dataPointer, nil
	}
	return "", errors.New("Record not found")
}
