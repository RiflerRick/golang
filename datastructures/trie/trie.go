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

func (t *Trie) createNode(key int) *Node {
	n := new(Node)
	n.left = nil
	n.right = nil
	n.key = key
	n.data = new(DataBlock)
	return n
}

/*
function to put data into the trie
*/
func (t *Trie) put(key string, data string) error {
	tKey := t.getKey(key)
	if t.t[tKey] == nil {
		fmt.Println("root node found nil")
		t.t[tKey] = t.createNode(tKey)
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
	fmt.Printf("Record %s written\n", data)
	return nil
}

/*
function to get data from the trie
*/
func (t *Trie) get(key string) (string, error) {
	tKey := t.getKey(key)

	if t.t[tKey] == nil {
		fmt.Println("root node found nil")
		return "", errors.New("Record not found")
	}
	currentNode := t.t[tKey]
	for i := 0; i < len(key); i++ {
		c := int(key[i])
		fmt.Println(c)
		if int(c) < currentNode.key {

			if currentNode.left == nil {
				fmt.Println("left node nil")
				return "", errors.New("Record not found")
			}
			fmt.Println("left node not nil")
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				fmt.Println("right node nil")
				return "", errors.New("Record not found")
			}
			fmt.Println("right node not nil")
			currentNode = currentNode.right
		}
	}
	if currentNode.data.dataPointer != nil {
		return *currentNode.data.dataPointer, nil
	}
	return "", errors.New("Record not found")
}
