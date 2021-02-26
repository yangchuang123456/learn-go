package go_grammar

import (
	"log"
	"runtime/debug"
	"testing"
	"time"
)

func f1(){
	log.Println("call f1")
}

func f2(){
	log.Println("call f2")
}

func Test_debugStack(t *testing.T){
	exit := make(chan bool)
	go func() {
		tick := time.NewTicker(time.Second)
		for {
			select {
			case <-tick.C:
				debug.PrintStack()
			case <-exit:
				return
			}
		}
	}()

	timer := time.NewTimer(5*time.Second)
	for {
		f1()
		f2()
		select {
		case <-timer.C:
			close(exit)
			return
		default:
			time.Sleep(time.Second)
			break
		}
	}
}
