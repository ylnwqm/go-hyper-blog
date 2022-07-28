package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-hyper-blog/controller/admin"
)

func InitApp(app *iris.Application) {
	//如果是闭包，则需要声明route，如果是模板，则需要声明Party,并在子路由使用iris.Party
	/*app.PartyFunc("/admin", func(r iris.Party) {
		r.Get("/login", func(ctx iris.Context) {
			//引用了admin.LoginController 赋值之后调用方法
			login := &admin.LoginController{}
			login.Login()
		})
	})*/

	//这里很奇怪，不能自定义路由名称，找的是8大方法  比如我要找Login，就找不到 只能找Get  //咋弄...
	mvc.New(app.Party("/login")).Handle(new(admin.LoginController))
}
