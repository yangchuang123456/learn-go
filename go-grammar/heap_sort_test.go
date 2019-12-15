package go_grammar

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

type sortHeap []int

func (s sortHeap) Len() int {
	return len(s)
}

func (s sortHeap) Less(i, j int) bool{
	return s[i] < s[j]
}

func (s sortHeap) Swap(i, j int){
	s[i],s[j] = s[j],s[i]
}

func (s *sortHeap) Push(x interface{}){
	*s=append(*s,x.(int))
}

func (s *sortHeap) Pop() interface{}{
	res := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]

	return res
}

func Test_Heap(t *testing.T) {
	testHeap := sortHeap{2,4,1,0,3,6,5}
	heap.Init(&testHeap)
	heap.Push(&testHeap,7)

	len := testHeap.Len()
	for i:=0;i<len;i++{
		value := heap.Pop(&testHeap)
		assert.Equal(t,i,value)
	}
}


