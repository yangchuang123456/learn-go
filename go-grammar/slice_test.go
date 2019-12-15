package go_grammar

import (
	"fmt"
	"log"
	"testing"
)

func Test_sliceCap(t *testing.T){

	a := make([]int, 1)
	b := make([]int, 1, 2)

	a[0] = 321

	fmt.Println("len(a)=", len(a), ", cap(a)=", cap(a))
	fmt.Println("len(b)=", len(b), ", cap(b)=", cap(b))

	a = append(a, 123)
	b = append(b, 123)

	fmt.Println("len(a)=", len(a), ", cap(a)=", cap(a))
	fmt.Println("len(b)=", len(b), ", cap(b)=", cap(b))

	c := a[:1]

	fmt.Println("len(a)=", len(a), ", cap(a)=", cap(a), ", a[0]=", a[0])
	fmt.Println("len(c)=", len(c), ", cap(c)=", cap(c), ", c[0]=", c[0])

	// 把原slice值覆盖了
	c = append(c, 456)

	fmt.Println("len(a)=", len(a), ", cap(a)=", cap(a), ", a[1]=", a[1])
	fmt.Println("len(c)=", len(c), ", cap(c)=", cap(c), ", c[1]=", c[1])

	// 分裂了
	c = append(c, 789)
	a = append(a, 666)

	fmt.Println("len(a)=", len(a), ", cap(a)=", cap(a), ", a[2]=", a[2])
	fmt.Println("len(c)=", len(c), ", cap(c)=", cap(c), ", c[2]=", c[2])

	fmt.Println("Hello, playground")


	//test cap
	d := make([]int,64)
	e := d[1:]
	f := d[:1]
	fmt.Println("len(d)=", len(d), ", cap(d)=", cap(d))
	fmt.Println("len(e)=", len(e), ", cap(e)=", cap(e))
	fmt.Println("len(f)=", len(f), ", cap(f)=", cap(f))
}


func Test_slice(t *testing.T){
	a := []byte{1,2,3,4,5}

	log.Println("the a[1:5] is:",a[1:4])
}


func changeSlice1(a []int){
	a[0] =2
}

func changeSlice2(a []int){
	fmt.Printf("the a addr is11:%p\r\n",&a)
	log.Println("the a len and cap is11:",len(a),cap(a))
	a = append(a,2)
	fmt.Printf("the a addr is22:%p\r\n",&a)
	log.Println("the a len and cap is22:",len(a),cap(a))
}

func Test_SliceFuncParameter(t *testing.T){
	a := []int{1,1,1}
	changeSlice1(a)
	log.Println("the a is:",a)
	changeSlice2(a)
	log.Println("the a is:",a)
}