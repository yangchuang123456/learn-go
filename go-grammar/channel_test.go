package go_grammar

import (
	"fmt"
	"testing"
)

func Test_channel(t *testing.T){
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}
