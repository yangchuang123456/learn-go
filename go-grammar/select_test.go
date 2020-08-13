package go_grammar

import "testing"

func Test_select(t *testing.T){
	testChan1 := make(chan bool,0)
	testChan2 := make(chan bool,1)

	select {
	case <-testChan1:
	case <-testChan2:
	//不加default会一直阻塞
	default:

	}
}


