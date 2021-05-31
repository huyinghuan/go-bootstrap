package main

import (
	"flag"
	"log"
	"strings"

	"talk.miniapp.mgtv.com/server"
)

var Version = "Dev"
var BuildTime = ""

func main() {
	var port string
	flag.StringVar(&port, "port", "", "端口号")
	flag.Parse()
	var msg = []string{
		"Programe",
		"Version  : %s",
		"BuildTime: %s",
		"Author:  : %s",
	}
	log.Println(strings.Join(msg, "  -"), Version, BuildTime, "yinghuan@mgtv.com")
	if port == "" {
		log.Fatal("端口不能为空,启动时需添加参数,如: --port 8888 ")
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	app := server.GetApp()
	app.Listen(":" + port)
}
