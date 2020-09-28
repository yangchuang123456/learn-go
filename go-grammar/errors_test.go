package go_grammar

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
	"testing"
)

func Test_error_wrap(t *testing.T){
	error1 := errors.New("error1")
	error2:= xerrors.Errorf("failed to look up actor state nonce: %w", error1)

	result:=xerrors.Is(error2,error1)
	assert.Equal(t,true,result)

	error3 := xerrors.Errorf("wrap error:%s",error1)
	result=xerrors.Is(error3,error1)
	assert.Equal(t,false,result)
}
