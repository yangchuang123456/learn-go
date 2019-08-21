package go_grammar

import (
	"log"
	"testing"
)

func Test_goRange(t *testing.T){
	a := make([]int,10)
	for i:= range a{
		a[i] = i
	}

	log.Printf("the a is:%v",a)

	for i,v := range a{
		if i == 3{
			a = append(a[:3],a[4:]...)
			log.Printf("the a is:%v",a)
		}
		log.Printf("the i is:%v",i)
		log.Printf("the v is:%v",v)
		log.Printf("~~~~~~~~~~~~~~~~~~~~~~")
	}
}
