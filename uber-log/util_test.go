package uber_log

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"testing"
)

//func TestRandSpecialCharacters(t *testing.T) {
//	x, err := RandSpecialCharacters(-1)
//	assert.Error(t, err)
//	fmt.Println(x)
//
//	x, err = RandSpecialCharacters(0)
//	assert.Error(t, err)
//
//	x, err = RandSpecialCharacters(1)
//	assert.NoError(t, err)
//	assert.Len(t, []rune(x), 1)
//	fmt.Println(x)
//
//	x, err = RandSpecialCharacters(2)
//	assert.NoError(t, err)
//	assert.Len(t, []rune(x), 2)
//	fmt.Println(x)
//
//	x, err = RandSpecialCharacters(3)
//	assert.NoError(t, err)
//	assert.Len(t, []rune(x), 3)
//	fmt.Println(x)
//
//	x, err = RandSpecialCharacters(4)
//	assert.NoError(t, err)
//	assert.Len(t, []rune(x), 4)
//	fmt.Println(x)
//
//	x, err = RandSpecialCharacters(5)
//	assert.NoError(t, err)
//	assert.Len(t, []rune(x), 5)
//	fmt.Println(x)
//}

func TestCurExecPath(t *testing.T) {
	fmt.Println(CurExecPath())
}

func TestRandANum(t *testing.T) {
	x := RandANum(8)
	fmt.Println(x)
	assert.True(t, x < 8)
	assert.True(t, x >= 0)
}

func TestCurExecPath1(t *testing.T) {
	xPath := "/home/ipfsmain/softwares/lotus"
	fmt.Println(filepath.Dir(xPath))
}

func Test_log(t *testing.T){
	err:=os.Setenv("XLotusLogOn","true")
	assert.NoError(t,err)
	l:=GetXDebugLog("lotus")
	l.Info("test debug",zap.String("key","value"))

	l2:=GetXDebugLog("actor")
	l2.Info("test debug",zap.String("key","value"))

	os.RemoveAll("./logs")
}
