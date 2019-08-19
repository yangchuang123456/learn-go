package crypto

import (
	"crypto/des"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func getTestData(length int)[]byte{
	slice := make([]byte,0)
	for i:=0;i<length;i++{
		slice = append(slice,0x01)
	}
	return slice
}

func Test_Des(t *testing.T){
	key :=getTestData(8)
	cryptBlock,err:=des.NewCipher(key)
	assert.NoError(t,err)

	plain := getTestData(15)
	cipher := make([]byte,16)
	cryptBlock.Encrypt(cipher,plain)

	log.Println("the cipher is:",cipher)

	plainText := make([]byte,16)
	cryptBlock.Decrypt(plainText,cipher)
	log.Println("the decrypted plain text is:",plainText)
}
