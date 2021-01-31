package hashTable

import (
	"algo/datastructures/linkedList"
	"hash/fnv"
)

type hashTable struct {
	array       []*linkedList.LinkedList
	currentSize int
	capacity    int
}

// This will contain the key and the value
type obj [2]interface{}

func (h hashTable) Size() int {
	return h.currentSize
}

func New() *hashTable {
	return &hashTable{
		array:       make([]*linkedList.LinkedList, 8),
		capacity:    8,
		currentSize: 0,
	}
}

func (h hashTable) hasExceededLoadFactor() bool {
	lambda := float32(h.currentSize) / float32(h.capacity)
	if lambda >= 0.75 {
		return true
	}

	return false
}

func (h hashTable) iterateLL(newArray []*linkedList.LinkedList, list *linkedList.LinkedList) {
	var currentNode *linkedList.Node
	for {
		currentNode = list.Next(currentNode)
		if currentNode != nil {
			val := currentNode.Value.(obj)
			key := val[0].(string)
			index, _ := h.getIndex(key)
			if newArray[index] == nil {
				ll := linkedList.New()
				newArray[index] = ll
			}

			newArray[index].Append(currentNode.Value)
			continue
		}

		break
	}
}

func (h *hashTable) rehashing() {
	h.capacity = h.capacity * 2

	newArray := make([]*linkedList.LinkedList, h.capacity)
	for _, list := range h.array {
		if list != nil {
			h.iterateLL(newArray, list)
		}
	}

	h.array = newArray
}

func (h *hashTable) Set(key string, value interface{}) bool {
	//Check if rehashing is needed
	if h.hasExceededLoadFactor() {
		h.rehashing()
	}

	index, err := h.getIndex(key)
	if err != nil {
		return false
	}

	if h.array[index] == nil {
		ll := linkedList.New()
		h.array[index] = ll
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
	h.currentSize--
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
