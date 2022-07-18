package routes

import "github.com/kataras/iris/v12"

func InitApp(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
}
