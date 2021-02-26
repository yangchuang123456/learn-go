package go_grammar

import (
	list2 "container/list"
	"log"
	"testing"
)

func Test_list(t *testing.T){
	list := list2.New()

	list.PushBack(1)
	list.PushBack(2)

	log.Println("the list value is:",list.Front().Value,list.Back().Value)

	list.Remove(list.Front())

	log.Println("the list value is:",list.Front().Value,list.Back().Value)
}
