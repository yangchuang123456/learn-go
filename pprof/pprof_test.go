package pprof

import (
	"github.com/stretchr/testify/assert"
	"os"
	"runtime/pprof"
	"testing"
	"time"
)

func Test_PProf(t *testing.T){
	cpuFile, err := os.Create("Profile.cpu")
	assert.NoError(t,err)
	defer cpuFile.Close()

	memFile, err := os.Create("Profile.mem")
	assert.NoError(t,err)
	defer memFile.Close()

	pprof.StartCPUProfile(cpuFile)
	pprof.WriteHeapProfile(memFile)

	//check profile code
	i :=100000
	for{
		i--

		func(){
			time.Sleep(100*time.Second)
		}()

		if i==0{
			break
		}
	}

	pprof.StopCPUProfile()

	os.Remove("Profile.cpu")
	os.Remove("Profile.mem")
}
