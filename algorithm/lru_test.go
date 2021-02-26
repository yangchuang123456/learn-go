package algorithm

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"testing"
)

type Node struct {
	key   int
	value int
}

type LRUCache struct {
	nodes    *list.List
	values   map[int]*list.Element
	capacity int
	lock     sync.RWMutex
}

func Constructor(capacity int) LRUCache {
	ret := LRUCache{
		nodes:    list.New(),
		values:   make(map[int]*list.Element, capacity),
		capacity: capacity,
		lock:     sync.RWMutex{},
	}
	return ret
}

func (this *LRUCache) Get(key int) int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	log.Println("get key ~~~~~~~~~~~~~~~~~~~~", key)
	log.Println("the map is~~~~~~~~~~~~~~~~~~~~:", this.values)
	if value, ok := this.values[key]; ok {
		log.Println("the get value is:", value.Value)
		log.Println("the remove value is:", value)
		this.nodes.Remove(value)
		this.nodes.PushBack(value.Value)
		this.values[key] = this.nodes.Back()
		//this.nodes.MoveToBack(&value)
		return value.Value.(*Node).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.lock.Lock()
	defer this.lock.Unlock()

	log.Println("put key ~~~~~~~~~~~~~~~~~~", key, value)
	log.Println("the map is~~~~~~~~~~~~~~~~~~~~ before putting:", this.values)

	if val, ok := this.values[key]; ok {
		this.nodes.Remove(val)
		val.Value.(*Node).value = value
		this.nodes.PushBack(val.Value)
		this.values[key] = this.nodes.Back()
		return
	}

	if this.nodes.Len() >= this.capacity {
		deleteNode := this.nodes.Front()
		this.nodes.Remove(deleteNode)
		log.Println("the deleteNode is:", deleteNode.Value)
		delete(this.values, deleteNode.Value.(*Node).key)
		log.Println("the map is after putting22:", this.values)
	}

	node := Node{key: key, value: value}
	this.nodes.PushBack(&node)
	this.values[key] = this.nodes.Back()
	return
}

func Test_lru(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	ret := lru.Get(1) // returns 1
	assert.Equal(t, 1, ret)

	lru.Put(3, 3)    // evicts key 2
	ret = lru.Get(2) // returns -1 (not found)
	assert.Equal(t, -1, ret)

	lru.Put(4, 4)    // evicts key 1
	ret = lru.Get(1) // returns -1 (not found)
	assert.Equal(t, -1, ret)

	ret = lru.Get(3) // returns 3
	assert.Equal(t, 3, ret)

	ret = lru.Get(4) // returns 4
	assert.Equal(t, 4, ret)
}

func generalLruTest(operation []string, parameter [][]int, result []interface{}, t *testing.T) {
	if operation[0] != "LRUCache" {
		panic("the first operation isn't LRUCache")
	}

	if len(parameter[0]) != 1 {
		panic("the first input parameter isn't LRUCache capacity")
	}

	var lru LRUCache
	for index, op := range operation {
		switch op {
		case "LRUCache":
			lru = Constructor(parameter[0][0])
		case "put":
			if len(parameter[index]) != 2 {
				panic("the put parameter isn't 2")
			}
			lru.Put(parameter[index][0], parameter[index][1])
		case "get":
			if len(parameter[index]) != 1 {
				panic("the get parameter isn't 1")
			}
			value := lru.Get(parameter[index][0])

			assert.Equal(t, result[index].(int), value)
		}
	}
}

func Test_lru_general1(t *testing.T) {
	operation := []string{"LRUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"}
	parameter := [][]int{
		{10}, {10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26},
	}
	result := []interface{}{
		nil, nil, nil, nil, nil, nil, -1, nil, 19, 17, nil, -1, nil, nil, nil, -1, nil, -1, 5, -1, 12, nil, nil, 3, 5, 5, nil, nil, 1, nil, -1, nil, 30, 5, 30, nil, nil, nil, -1, nil, -1, 24, nil, nil, 18, nil, nil, nil, nil, -1, nil, nil, 18, nil, nil, -1, nil, nil, nil, nil, nil, 18, nil, nil, -1, nil, 4, 29, 30, nil, 12, -1, nil, nil, nil, nil, 29, nil, nil, nil, nil, 17, 22, 18, nil, nil, nil, -1, nil, nil, nil, 20, nil, nil, nil, -1, 18, 18, nil, nil, nil, nil, 20, nil, nil, nil, nil, nil, nil, nil,
	}

	generalLruTest(operation, parameter, result, t)
}

func Test_lru_general2(t *testing.T) {
	operation := []string{
		"LRUCache", "put", "put", "get", "put", "put", "get"}

	parameter := [][]int{
		{2}, {2, 1}, {2, 2}, {2}, {1, 1}, {4, 1}, {2},
	}

	result := []interface{}{
		nil,nil,nil,2,nil,nil,-1,
	}

	generalLruTest(operation, parameter, result, t)
}
