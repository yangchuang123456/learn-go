package go_grammar

import (
	"log"
	"math/big"
	"testing"
)

func Test_big_int(t *testing.T){
	a := big.NewInt(10)
	a.Div(a,big.NewInt(2))
	log.Println("the a is:",a)

	a.Div(big.NewInt(20),big.NewInt(5))
	log.Println("the a is:",a)


	a = big.NewInt(100)
	b := big.NewInt(0)
	b.Cmp(a.Div(a,big.NewInt(2)))
	b.Div(a,big.NewInt(2))

	log.Println("the a and b is:",a,b)

	a = big.NewInt(100)
	b = big.NewInt(0).Div(a,big.NewInt(2))
	log.Println("the a and b is:",a,b)
}
