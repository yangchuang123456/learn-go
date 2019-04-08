package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Travis(t *testing.T){
	err :=  errors.New("test err")
	assert.Error(t,err)
}
