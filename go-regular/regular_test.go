package go_regular

import (
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
)

func Test_Regular(t *testing.T){
	str := "vm213vmrivmmewarvm$532vm1232vm"

	pattern1 := "^vm"
	regular1,err := regexp.Compile(pattern1)
	assert.NoError(t,err)
	log.Printf("the pattern1 match result is:%v",regular1.FindAllStringIndex(str,-1))

	pattern2 := "vm$"
	regular2,err := regexp.Compile(pattern2)
	assert.NoError(t,err)
	log.Printf("the pattern1 match result is:%v",regular2.FindAllStringIndex(str,-1))

	pattern3 := "vm"
	regular3,err := regexp.Compile(pattern3)
	assert.NoError(t,err)
	log.Printf("the pattern1 match result is:%v",regular3.FindAllStringIndex(str,-1))

}
