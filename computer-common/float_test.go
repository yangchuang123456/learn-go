package computer_common

import (
	"log"
	"testing"
)

func Test_float(t *testing.T){
	a := float32(34.6)
	b := float32(34.0)
	log.Println("the a-b is:",a-b)
	a = float32(34.5)
	log.Println("the a-b is:",a-b)

	a = float32(34.7)
	log.Println("the a-b is:",a-b)
}
