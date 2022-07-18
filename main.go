package main

import (
	"github.com/kataras/iris/v12"
	config "go-hyper-blog/configs"
	"os"
	"path"
)

func newApp() *iris.Application {
	iris := iris.New()
	//注册路由文件
	//router.InitApp(iris)

	//定位yml文件
	currentDir, _ := os.Getwd()
	configFileDir := path.Join(currentDir, "config.yml")
	config.Init(configFileDir)
	println(config.Setting.App.BindAddress)
	println(config.Setting.App.Port)
	println(config.Setting.Log.Level)
	return iris
}

func main() {
	newApp()
}
