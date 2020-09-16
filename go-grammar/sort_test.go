package go_grammar

import (
	"log"
	"math/rand"
	"sort"
	"testing"
)

type sortStruct struct {
	a int
}

func Test_Sort(t *testing.T) {
	slice := make([]sortStruct,0)
	for i:=0;i<5;i++{
		slice = append(slice,sortStruct{a:rand.Int()})
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].a > slice[j].a
	})

	log.Println("the slice is:")
	log.Println(slice)
}
