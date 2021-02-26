package main

import (
	"github.com/GeertJohan/go.rice"
	"log"
)

func main() {
	// 这里写相对于的执行文件的地址
	box, err := rice.FindBox("theme/default")
	if err != nil {
		println(err.Error())
		return
	}
	// 从目录 Box 读取文件
	str, err := box.String("post.txt")
	if err != nil {
		println(err.Error())
		return
	}

	log.Println("the str is:",str)
}