package app_run

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"learn-go/urfave-cli/flags"
	"log"
	"os"
	"testing"
)

func print_flagInfos(ctx *cli.Context){
	return
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

func Test_appRun(t *testing.T){
	app := cli.NewApp()
	app.Flags = append(flags.Flags, flags.PProfFlags...)

	log.Println("the app flags number is:","number",len(app.Flags))

	app.Action = func(c *cli.Context) error {
		print_flagInfos(c)
		return nil
	}

	errArgs := []string{"./main","--node_name","default_b0","--data_dir","/home/qydev/tmp/dipperin_apps/default_b0","--node_type","3","--p2p_listener",":10003","--http_port","10004","--ws_port","10005","--ipc_path","/home/qydev/tmp/dipperin_apps/default_b0/dipperin.ipc","--soft_wallet_pwd","123","--is_start_mine","1","--pprof=true","--debug_mode","2","--log_level","info","--is_upload_node_data","1","--upload_url","http://127.0.0.1:8887/api/dipperin_nodes","--m_port","9202","--pprofport","20002"}
	err := app.Run(errArgs)
	assert.NoError(t,err)

	log.Println("*****************************************")
	log.Println("*****************************************")
	log.Println("*****************************************")


	correctArgs := []string{"./main","--node_name","default_b0","--data_dir","/home/qydev/tmp/dipperin_apps/default_b0","--node_type","3","--p2p_listener",":10003","--m_port","9202","--pprofport","20002","--http_port","10004","--ws_port","10005","--ipc_path","/home/qydev/tmp/dipperin_apps/default_b0/dipperin.ipc","--soft_wallet_pwd","123","--is_start_mine","1","--pprof=true","--debug_mode","2","--log_level","info","--is_upload_node_data","1","--upload_url","http://127.0.0.1:8887/api/dipperin_nodes"}
	err = app.Run(correctArgs)
	assert.NoError(t,err)
}

func Test_ParsePar(t *testing.T){
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		log.Println("the parse parameter is:",c.Args().Get(0))
		return nil
	}

	testInput := "proving"
	app.Run([]string{"./testParsePar",testInput})

	log.Println("the os args is:",os.Args)
}
