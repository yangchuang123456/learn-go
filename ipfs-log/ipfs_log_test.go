package ipfs_log

import (
	logging "github.com/ipfs/go-log"
	"os"
	"testing"
)
var log = logging.Logger("sectors")

func Test_ipfs_golog(t *testing.T){
	os.Setenv("GOLOG_TRACING_FILE", "ipfs_golog.file")
	logging.SetupLogging()
	logging.SetLogLevel("sectors","info")

	log.Info("test ipfs go log")
}
