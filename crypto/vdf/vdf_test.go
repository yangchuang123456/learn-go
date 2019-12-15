package vdf

import (
	"encoding/hex"
	"fmt"
	"github.com/harmony-one/vdf/src/vdf_go"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func Test_vdf(t *testing.T){
	input :=  [32]byte{0xde, 0xad, 0xbe, 0xef,0xde, 0xad, 0xbe, 0xef,0xde, 0xad, 0xbe, 0xef,0xde, 0xad, 0xbe,
		0xef,0xde, 0xad, 0xbe, 0xef,0xde, 0xad, 0xbe, 0xef,0xde, 0xad, 0xbe, 0xef,0xde, 0xad, 0xbe, 0xef,}
	vdf := vdf_go.New(10, input)

	outputChannel := vdf.GetOutputChannel()
	start := time.Now()

	vdf.Execute()

	duration := time.Now().Sub(start)


	output := <-outputChannel

	log.Println(fmt.Sprintf("VDF computation finished, result is  %s", hex.EncodeToString(output[:])))
	log.Println(fmt.Sprintf("VDF computation finished, time spent %s", duration.String()))
	assert.Equal(t, true, vdf.Verify(output), "failed verifying proof")
}
