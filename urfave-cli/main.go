package main

import (
	"fmt"
	"github.com/urfave/cli"
	"learn-go/urfave-cli/flags"
	"log"
	"os"
)

func print_flagInfos(ctx *cli.Context){
	log.Println("the context arg number is:","number",ctx.NArg())
	log.Println("the context Flags args is:","args",ctx.Args())

	log.Println("the context Flags number is:","number",ctx.NumFlags())
	log.Println("the context Flags names is:","name",ctx.FlagNames())
	log.Println("the pprofFlag is:","pprofFlag",ctx.GlobalBool("pprof"))
	log.Println("the http_port is:","http_port",ctx.GlobalInt("http_port"))
	log.Println("the m_port is:","m_port",ctx.GlobalInt("m_port"))
	log.Println("the pprofPortFlag is:","pprofPortFlag",ctx.GlobalInt("pprofport"))
	log.Println("the upload_url flag is:","upload_url",ctx.GlobalString("upload_url"))
}

func main(){
	app := cli.NewApp()
	app.Flags = append(flags.Flags, flags.PProfFlags...)

	fmt.Printf("the \" is \" ")

	log.Println("the app flags number is:","number",len(app.Flags))

	app.Action = func(c *cli.Context) error {
		print_flagInfos(c)
		return nil
	}

	log.Println("the os args is:","args",os.Args)
	log.Println("the os args len is:","len",len(os.Args))
	printString := "{"
	var tmpArg string
	for i,arg := range os.Args{
		tmpArg ="\"" +arg+"\""
		if i==0{
			printString = printString+tmpArg
		}else {
			printString = printString + "," +tmpArg
		}
	}
	printString += "}"
	fmt.Printf("the printString is:%v\r\n",printString)

	if err := app.Run(os.Args); err != nil {
		log.Println("dipperin run failed", "err", err)
		os.Exit(1)
	}
}
