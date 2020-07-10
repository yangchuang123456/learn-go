package go_rice

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_findBox(t *testing.T)  {
	builtinGen, err := rice.FindBox("/home/ipfsmain/yc/project/learn-go/go-rice")
	log.Println("the builtinGen info is:","builtinGen")
	assert.NoError(t,err)
	genBytes, err := builtinGen.Bytes("devnet.car")
	assert.NoError(t,err)

	log.Println("the genBytes is:",genBytes)
}
