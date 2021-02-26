package go_grammar

import (
	"log"
	"testing"
)

func Test_negative(t *testing.T){
	a := -123
	b := a/10
	c := a%10
	log.Println("the b c is:",b,c)
}
