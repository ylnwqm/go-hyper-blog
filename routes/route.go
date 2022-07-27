package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

func InitApp(app *iris.Application) {
	app.PartyFunc("/admin", func(r router.Party) {
		/*r.Handle("GET", "/login", func(context iris.Context) {
			context.JSON(iris.Map{"message": "pong"})
		})*/

		app.Get("/login", func(ctx iris.Context) {
			// 绑定： {{.message}}　为　"Hello world!"
			println("就按就按家家爱")
			ctx.ViewData("message", "Hello world!")
			// 渲染模板文件： ./views/hello.html
			ctx.View("hello.html")
		})
	})
}
