package go_grammar

import (
	"log"
	"reflect"
	"testing"
)

var testVariable  = 123

func Test_getVariableName(t *testing.T){
	value := reflect.ValueOf(testVariable)
	log.Printf("the variable name is:%v",value.NumField())
}
