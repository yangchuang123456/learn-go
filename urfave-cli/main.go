package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
	"github.com/yangchuang123456/learn-go/urfave-cli/flags"
	"os"
)

func print_flagInfos(ctx *cli.Context){
	log.Info("the context arg number is:","number",ctx.NArg())
	log.Info("the context Flags args is:","args",ctx.Args())

	log.Info("the context Flags number is:","number",ctx.NumFlags())
	log.Info("the context Flags names is:","name",ctx.FlagNames())
	log.Info("the pprofFlag is:","pprofFlag",ctx.GlobalBool("pprof"))
	log.Info("the http_port is:","http_port",ctx.GlobalInt("http_port"))
	log.Info("the m_port is:","m_port",ctx.GlobalInt("m_port"))
	log.Info("the pprofPortFlag is:","pprofPortFlag",ctx.GlobalInt("pprofport"))
	log.Info("the upload_url flag is:","upload_url",ctx.GlobalString("upload_url"))
}

func initLog(maxLvl log.Lvl){
	log.Root().SetHandler(log.LvlFilterHandler(maxLvl, log.StdoutHandler))
}

func main(){
	initLog(log.LvlInfo)
	app := cli.NewApp()
	app.Flags = append(flags.Flags, flags.PProfFlags...)

	fmt.Printf("the \" is \" ")

	log.Info("the app flags number is:","number",len(app.Flags))

	app.Action = func(c *cli.Context) error {
		print_flagInfos(c)
		return nil
	}

	log.Info("the os args is:","args",os.Args)
	log.Info("the os args len is:","len",len(os.Args))
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
		log.Error("dipperin run failed", "err", err)
		os.Exit(1)
	}
}
