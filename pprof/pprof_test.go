package pprof

import (
	"github.com/stretchr/testify/assert"
	"os"
	"runtime/pprof"
	"testing"
)

func Test_PProf(t *testing.T){
	f, err := os.Create("cpuProfile.txt")
	assert.NoError(t,err)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	f, err = os.Create("memProfile")
	assert.NoError(t,err)
	pprof.WriteHeapProfile(f)
	f.Close()
}
