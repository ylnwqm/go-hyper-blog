package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
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

const cookieNameForSessionID = "session_id_cookie"

func main() {
	app := iris.New()
	sess := sessions.New(sessions.Config{
		Cookie: cookieNameForSessionID,
		// CookieSecureTLS: true,
		AllowReclaim: true,
	})
	app.Use(sess.Handler())
	app.Get("/login", login)
	app.Listen(":8888")
	/*app := newApp()
	err := app.Run(iris.Addr(fmt.Sprintf("%s:%d", config.Setting.App.BindAddress, config.Setting.App.Port)))
	if err != nil {
		app.Logger().Print("Server failed to start")
	}*/
}

func login(ctx iris.Context) {
	println(111111111111)
	session := sessions.Get(ctx)

	// Authentication goes here
	// ...
	println(22222)
	// Set user as authenticated
	session.Set("authenticated", "哈哈哈哈哈")
	auth := session.GetString("authenticated")
	fmt.Printf("%+v", auth)
}
