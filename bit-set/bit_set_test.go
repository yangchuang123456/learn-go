package bit_set

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"

	"github.com/willf/bitset"
	"log"
	"testing"
)


func Test_bitSet(t *testing.T) {
	b := bitset.New(64)
	b.Set(1)
	b.Set(63)

	log.Printf("the b is:%s", b.String())

	buff := make([]byte,8)
	binary.BigEndian.PutUint64(buff,b.Bytes()[0])

	setIndex,result := b.NextSet(0)
	log.Printf("the setIndex is:%v",setIndex)
	log.Printf("the result is:%v",result)
	log.Printf("the b is:%s", b.String())

	setBuffer := make([]uint,3)
	b.Set(2)
	b.Set(0)
	number,sets:= b.NextSetMany(0,setBuffer)

	log.Printf("the number is:%v",number)
	log.Printf("the sets is:%v",sets)

	setCount := b.Count()
	assert.Equal(t,uint(4),setCount)
}