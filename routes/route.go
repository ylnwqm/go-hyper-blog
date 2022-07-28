package routes

import (
	"github.com/kataras/iris/v12"
)

func InitApp(app *iris.Application) {
	//如果是闭包，则需要声明route，如果是模板，则需要声明Party,并在子路由使用iris.Party
	app.PartyFunc("/admin", func(r iris.Party) {
		r.Get("/login", func(ctx iris.Context) {
			// 绑定： {{.message}}　为　"Hello world!"
			println("就按就按家家爱")
			//ctx.ViewData("message", "Hello world!")
			// 渲染模板文件： ./views/hello.html
			//ctx.View("hello.html")
		})
	})
}
