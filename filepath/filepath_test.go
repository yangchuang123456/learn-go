package filepath

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("the dir is:%v",dir)
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}


func Test_filePathWalk(t *testing.T){
	currentDir := GetCurrentDirectory()
	log.Printf("the current dir is:%v",currentDir)

	filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		log.Printf("the path is:%v",path)
		//log.Printf("the walk info name is:%v",info.Name())
		return nil
	})

}
