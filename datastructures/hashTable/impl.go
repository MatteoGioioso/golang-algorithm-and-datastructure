package hashTable

import (
	"algo/datastructures/linkedList"
	"hash/fnv"
)

type hashTable struct {
	array       [8]*linkedList.LinkedList
	currentSize int
	capacity int
}

// This will contain the key and the value
type obj [2]interface{}

func (h hashTable) Size() int {
	return h.currentSize
}

func New() *hashTable {
	const initialSize = 8
	return &hashTable{
		array: [initialSize]*linkedList.LinkedList{},
		capacity: 0,
		currentSize: initialSize,
	}
}

func (h *hashTable) Set(key string, value interface{}) bool {
	index, err := h.getIndex(key)
	if err != nil {
		return false
	}

	// Check if index is out of bound
	if index > h.capacity {
		// TODO implement array resize
		return false
	}

	data := obj{key, value}
	h.array[index].Append(data)
	h.currentSize++
	return true
}

func (h hashTable) Get(key string) interface{} {
	index, _ := h.getIndex(key)

	node := h.searchInLinkedList(key, index)
	if node == nil {
		return nil
	}

	o := node.Value.(obj)
	return o[1]
}

func (h *hashTable) Delete(key string) bool {
	index, err := h.getIndex(key)
	if err != nil {
		return false
	}

	node := h.searchInLinkedList(key, index)

	if node == nil {
		return true
	}

	h.array[index].Delete(node)
	h.capacity--
	return true
}

func (h hashTable) searchInLinkedList(key string, index int) *linkedList.Node {
	var currentNode *linkedList.Node
	for {
		currentNode = h.array[index].Next(currentNode)

		if currentNode == nil {
			return nil
		}

		o := currentNode.Value.(obj)

		if o[0] == key {
			return currentNode
		}
	}
}

func (h hashTable) getIndex(s string) (int, error) {
	hash, err := h.hash(s)
	if err != nil {
		return 0, err
	}

	return hash % h.capacity, err
}

func (h hashTable) hash(s string) (int, error) {
	hash := fnv.New32a()
	if _, err := hash.Write([]byte(s)); err != nil {
		return 0, nil
	}
	return int(hash.Sum32()), nil
}
