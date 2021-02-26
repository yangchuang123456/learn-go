package common

import "sync"

type Queue struct {
	data []interface{}
	mux sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{
		data:make([]interface{},0),
	}
}

func (q *Queue) Enqueue(data interface{}) {
	q.mux.Lock()
	defer q.mux.Unlock()
	q.data = append(q.data,data)
}

func (q *Queue) Dequeue() interface{}{
	q.mux.Lock()
	defer q.mux.Unlock()
	ret := q.data[0]
	q.data = q.data[1:]
	return ret
}

func (q *Queue) Front() interface{}{
	q.mux.RLock()
	defer q.mux.RUnlock()
	return q.data[0]
}

func (q *Queue) Size() int{
	return len(q.data)
}

func (q *Queue) IsEmpty() bool{
	if len(q.data) == 0{
		return true
	}else{
		return false
	}
}
