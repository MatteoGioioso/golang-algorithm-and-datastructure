package hashTable

type HashTable interface {
	Set(key string, value interface{}) bool
	Get(key string) interface{}
	Delete(key string) bool
	Size() int
}
