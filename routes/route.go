package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

func InitApp(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	//PartyFunc类似于声明路由组api   r.hander声明下一级路由
	app.PartyFunc("/api", func(r router.Party) {
		r.Handle("GET", "/ad", func(context iris.Context) {
			context.JSON(iris.Map{"message": "pong"})
		})
	})
}
