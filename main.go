package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go-hyper-blog/configs"
	"go-hyper-blog/database"
	"go-hyper-blog/routes"
	"os"
	"path"
)

func newApp() *iris.Application {
	app := iris.New()
	//使用中间件  如果有的话
	//app.Use()
	//注册路由文件
	routes.InitApp(app)
	//定位yml文件
	currentDir, _ := os.Getwd()
	configFileDir := path.Join(currentDir, "config.yml")
	config.Init(configFileDir)
	//注册iris配置
	app.Configure(iris.WithConfiguration(config.Setting.Iris))

	//配置连接数据库
	database.ConnectDb(app)
	/**
	1. 声明一个result map
	2. user 赋值model
	3. db.Model(&user) 指定model ，如果只有First 则为result的model 会报错
	4. Frist(&result)   将查询结果赋值给result
	*/
	/*result := map[string]interface{}{}
	var user model.Admin
	db.Model(&user).First(&result)
	fmt.Printf("%+v", result)*/

	//注册html
	tmpl := iris.HTML("./views", ".html")
	tmpl.Reload(true)
	//定义静态统一文件目录路径
	tmpl.AddFunc("static", func() string {
		return "/static/admin"
	})
	app.RegisterView(tmpl)

	//注册全局中间件  目前没有用到，取自其他项目
	//app.UseGlobal(middleware.IrisRequestHandler)
	//配置静态文件
	app.HandleDir("/static", "./static")
	//错误页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("admin/pages-404.html")
	})
	return app
}

func main() {
	app := newApp()
	err := app.Run(iris.Addr(fmt.Sprintf("%s:%d", config.Setting.App.BindAddress, config.Setting.App.Port)))
	if err != nil {
		app.Logger().Print("Server failed to start")
	}
}
