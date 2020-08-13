package go_grammar

import (
	"log"
	"testing"
)

func Test_append(t *testing.T) {
	a := make([]*int, 0, 5)
	a = append(a, nil...)
	log.Println("the a is;", a)
	b := 1
	a = append(a, &b)
	log.Println("the a is;", a)
	log.Println("the a[0:0] is:", a[0:1])
}
