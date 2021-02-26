package pprof

import (
	"github.com/stretchr/testify/assert"
	"os"
	"runtime/pprof"
	"testing"
	"time"
)

func executeFunc(){
	//i :=200000
	i:=10
	for{
		i--

		func(){
			time.Sleep(time.Second)
		}()
//		sha256.Sum256([]byte{byte(i)})
		if i==0{
			break
		}
	}
}

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
	executeFunc()

	pprof.StopCPUProfile()

	//os.Remove("Profile.cpu")
	//os.Remove("Profile.mem")
}

func Test_PProf_scope(t *testing.T){
	cpuFile1, err := os.Create("Profile1.cpu")
	assert.NoError(t,err)
	defer cpuFile1.Close()
	go func() {
		err:=pprof.StartCPUProfile(cpuFile1)
		assert.NoError(t,err)
	}()
	for{
		;
	}
}
