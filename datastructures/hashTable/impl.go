package hashTable

import (
	"hash/fnv"
)

type hashTable struct {
	array       [8][3]interface{}
	currentSize int
	capacity int
}

func (h hashTable) Size() int {
	return h.currentSize
}

func New() *hashTable {
	return &hashTable{
		[8][3]interface{}{},
		0,
		8,
	}
}

func (h hashTable) Array() [8][3]interface{} {
	return h.array
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

	// Check if there is a collision
	if h.array[index][0] != nil {
		// TODO implement with linked list
		return false
	}

	obj := [3]interface{}{
		key,
		value,
		nil, // pointer to the next node
	}

	h.array[index] = obj
	h.currentSize++
	return true
}

func (h hashTable) Get(key string) interface{} {
	index, _ := h.getIndex(key)
	obj := h.array[index]

	return obj[1]
}

func (h *hashTable) Delete(key string) bool {
	index, err := h.getIndex(key)
	if err != nil {
		return false
	}

	h.array[index] = [3]interface{}{}
	h.capacity--
	return true
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
