package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("Hello <strong>%s</strong>!", "World")
	})

	app.Listen(":8080")
}
