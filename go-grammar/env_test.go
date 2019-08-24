package go_grammar

import (
	"log"
	"os/exec"
	"regexp"
	"testing"
)


func getGoPath()string{
	cmd := exec.Command("go","env")
	str,_ := cmd.Output()

	pattern := `GOPATH="(.*?)"{1}?`
	regular1,_ := regexp.Compile(pattern)
	out := regular1.FindStringSubmatch(string(str))
	return string(out[1])
}

func Test_GetGoPath(t *testing.T){
	//goPath := os.Getenv("GOPATH")
	//log.Println("the goPath is:",goPath)

	goPath := getGoPath()
	log.Println("the goPath is:",goPath)
}


