package go_grammar

import "testing"

func generateIntPointer() (*int,error){
	a :=1
	return &a,nil
}

//golang中　if 和 for声明的变量作用域属于控制代码块内部．因此下述代码语法错误
func variable() *int{
/*	if b,err := generateIntPointer();err!=nil{
		panic("test 1324324")
	}
	return b*/
	return nil
}

func Test_variable(t *testing.T){

}
