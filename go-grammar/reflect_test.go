package go_grammar

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

var testVariable  = 123

func Test_getVariableName(t *testing.T){
	value := reflect.ValueOf(testVariable)
	log.Printf("the variable name is:%v",value.NumField())
}


// 通用算法函数
func add(args []reflect.Value) (results []reflect.Value) {
	if len(args) == 0 {
		return nil
	}

	var ret reflect.Value

	switch args[0].Kind() {
	case reflect.Int:
		n := 0
		for _, a := range args {
			n += int(a.Int())
		}

		ret = reflect.ValueOf(n)
	case reflect.String:
		ss := make([]string, 0, len(args))
		for _, s := range args {
			ss = append(ss, s.String())
		}

		ret = reflect.ValueOf(strings.Join(ss, ""))
	}

	results = append(results, ret)
	return
}

// 将函数指针参数指向通用算法函数
func makeAdd(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), add) // 这是关键
	fn.Set(v)                             // 指向通用算法函数
}

func Test_makeFunc(t *testing.T){
	var intAdd func(x, y int) int
	var strAdd func(a, b string) string

	makeAdd(&intAdd)
	makeAdd(&strAdd)

	println(intAdd(100, 200))
	println(strAdd("hello, ", "world!"))
}

func Test_getStructField(t *testing.T){
	type type1 struct {
		A int
	}
	type type2 struct {
		B string
	}

	type type3 struct {
		C int
		D type1
		E type2
	}

	type type4 struct {
		C int
		type1
		type2
	}

	testValue := type3{
		C:0,
		D: type1{
			A:1,
		},
		E:type2{
			B:"test",
		},
	}
	value := reflect.ValueOf(testValue)
	number:=value.NumField()
	for i:=0;i<number;i++{
		log.Println("the field is:",value.Field(i))
	}

	log.Println(value.FieldByName("D"))
	log.Println(value.FieldByName("A"))

	testValue2 := type4{
		C:0,
		type1:type1{
			A: 1,
		},
		type2:type2{
			B:"test",
		},
	}
	value = reflect.ValueOf(testValue2)
	number =value.NumField()
	for i:=0;i<number;i++{
		log.Println("the field is:",value.Field(i))
	}
	log.Println(value.FieldByName("A"))
	log.Println(value.FieldByName("B"))
}
