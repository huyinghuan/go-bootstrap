package server

import "github.com/kataras/iris/v12"

func GetApp() *iris.Application {
	app := iris.New()
	return app
}
