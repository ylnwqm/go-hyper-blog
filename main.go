package main

import (
	"github.com/kataras/iris/v12"
	config "go-hyper-blog/configs"
	"go-hyper-blog/middleware"
	"os"
	"path"
)

func newApp() *iris.Application {
	app := iris.New()
	//注册路由文件
	//router.InitApp(iris)

	//定位yml文件
	currentDir, _ := os.Getwd()
	configFileDir := path.Join(currentDir, "config.yml")
	config.Init(configFileDir)

	//注册iris配置
	app.Configure(iris.WithConfiguration(config.Setting.Iris))
	//注册全局中间件  目前没有用到，取自其他项目
	app.UseGlobal(middleware.IrisRequestHandler)
	return app
}

func main() {
	newApp()
}
