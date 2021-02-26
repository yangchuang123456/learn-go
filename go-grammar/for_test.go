package go_grammar

import (
	"log"
	"testing"
)

func Test_for(t *testing.T){
	i:=0
	for ;i<3;i++{
		break
	}
	log.Println("the i is:",i)
}
