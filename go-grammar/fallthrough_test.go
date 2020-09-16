package go_grammar

import (
	"fmt"
	"testing"
)

func fallThrough(a int) {
	switch {
	case a == 0:
		//fmt.Println("case 0")
		fallthrough
	case a == 1:
		//fmt.Println("case 1")
		fallthrough
	case a == 2:
		fmt.Println("case 2")
		return
	default:
		fmt.Println("default")
		return
	}
	fmt.Println("fall through end")
}

func Test_fallThrough(t *testing.T) {
	a := 0
	fallThrough(a)
	a = 1
	fallThrough(a)
	a = 2
	fallThrough(a)
	a = 3
	fallThrough(a)
}
