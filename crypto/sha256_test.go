package crypto

import (
	"crypto/sha256"
	"log"
	"testing"
)

func Test_hash256(t *testing.T){
	src := make([]byte,7878667676)
	hash :=sha256.Sum256(src)
	log.Println("the calculated hash is:",hash)
}
