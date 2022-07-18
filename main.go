package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-hyper-blog/configs"
	"go-hyper-blog/routes"
	"os"
	"path"
)

type controller struct{}

type errorResponse struct {
	Code    int
	Message string
}

func newApp() *iris.Application {
	app := iris.New()
	//注册模板引擎
	mvcInit(app)
	//注册路由文件
	routes.InitApp(app)
	//定位yml文件
	currentDir, _ := os.Getwd()
	configFileDir := path.Join(currentDir, "config.yml")
	config.Init(configFileDir)
	//注册iris配置
	app.Configure(iris.WithConfiguration(config.Setting.Iris))

	//注册全局中间件  目前没有用到，取自其他项目
	//app.UseGlobal(middleware.IrisRequestHandler)
	return app
}

//初始化引擎
func mvcInit(app *iris.Application) {
	//注册Django模板引擎
	app.RegisterView(iris.Django("/template", ".html"))
	m := mvc.New(app)
	m.Handle(new(controller))

	//渲染500和403错误

}

// implements mvc.Result.
func (e errorResponse) Dispatch(ctx iris.Context) {
	// If u want to use mvc.Result on any method without an output return value
	// go for it:
	//
	view := mvc.View{Code: e.Code, Data: e} // use Code and Message as the template data.
	switch e.Code {
	case iris.StatusNotFound:
		view.Name = "404"
	default:
		view.Name = "500"
	}
	view.Dispatch(ctx)

	// Otherwise use ctx methods:
	//
	// ctx.StatusCode(e.Code)
	// switch e.Code {
	// case iris.StatusNotFound:
	// 	// use Code and Message as the template data.
	// 	ctx.View("404.html", e)
	// default:
	// 	ctx.View("500.html", e)
	// }
}

func main() {
	app := newApp()
	err := app.Run(iris.Addr(fmt.Sprintf("%s:%d", config.Setting.App.BindAddress, config.Setting.App.Port)))
	if err != nil {
		app.Logger().Print("Server failed to start")
	}
}
