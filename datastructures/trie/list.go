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
type ListNode struct {
	key  string
	next *ListNode
	data *DataBlock
}

/*
Trie implementation as a map
*/
type List struct {
	root *ListNode
}

/*
function to put data into the trie
*/
func (l *List) put(key string, data string) error {
	rootNode := l.root
	newData := new(ListNode)
	if rootNode == nil {
		newData.key = key
		newData.next = nil
		newData.data = new(DataBlock)
		newData.data.dataPointer = &data
		l.root = newData
		fmt.Printf("Record %s written\n", data)
		return nil
	}
	currentNode := rootNode
	nextNode := currentNode.next
	for true {
		if nextNode == nil {
			newData.key = key
			newData.next = nil
			newData.data = new(DataBlock)
			newData.data.dataPointer = &data
			currentNode.next = newData
			fmt.Printf("Record %s written\n", data)
			return nil
		}
		nextNode = nextNode.next
	}
	return nil
}

/*
function to get data from the trie
*/
func (l *List) get(key string) (string, error) {
	rootNode := l.root
	if rootNode == nil {
		return "", errors.New("Record not found")
	}
	currentNode := rootNode
	nextNode := currentNode.next
	for true {
		if currentNode.key == key {
			return *currentNode.data.dataPointer, nil
		}
		nextNode = nextNode.next
		if nextNode == nil {
			break
		}
	}
	return "", errors.New("Record not found")
}
