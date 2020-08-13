package go_grammar

import (
	"log"
	"runtime"
	"testing"
)

func Test_map(t *testing.T) {
	a := make(map[int]int)
	a[1] =1
	a[2] =2

	b := a
	a = map[int]int{}
	runtime.GC()

	log.Println("the a is: ",a)
	log.Println("the b is: ",b)
}

func Test_map_copy(t *testing.T) {
	a := make(map[int]int)
	a[1] =1
	a[2] =2

	b := a
	b[3] = 3
	b[4] = 4
	b[5] = 5

	log.Println("the a is: ",a)
	log.Println("the b is: ",b)
}