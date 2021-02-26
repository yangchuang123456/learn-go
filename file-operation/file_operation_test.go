package file_operation

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var testFilePath = "testFile"

func Test_FileOperation(t *testing.T) {
	file, err := os.Create(testFilePath)
	assert.NoError(t, err)

	file.Write([]byte("1234568787"))
	file.Close()

	file, err = os.OpenFile(testFilePath, os.O_RDWR|os.O_CREATE, 0644)
	assert.NoError(t, err)

	stat, err := file.Stat()
	assert.NoError(t, err)
	readByte := make([]byte, stat.Size())
	_, err = file.Read(readByte)
	assert.NoError(t,err)

	_,err=file.Seek(0,0)
	assert.NoError(t,err)
	err=file.Truncate(0)
	assert.NoError(t,err)

	file.Write([]byte("ycTest"))
	file.Close()
}
