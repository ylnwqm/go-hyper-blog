package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-hyper-blog/controller/admin"
)

func InitApp(app *iris.Application) {
	//
	/**
	1. 根据iris的路由规则
		func(*Controller) Get() - GET:/user.
		func(*Controller) Post() - POST:/user.
		func(*Controller) GetLogin() - GET:/user/login
		func(*Controller) PostLogin() - POST:/user/login
		func(*Controller) GetProfileFollowers() - GET:/user/profile/followers
		func(*Controller) PostProfileFollowers() - POST:/user/profile/followers
		func(*Controller) GetBy(id int64) - GET:/user/{param:long}
		func(*Controller) PostBy(id int64) - POST:/user/{param:long}

	2. 使用mvc定义路由组，也需要定义PartyFunc，同时将iris.Party注入，并在使用时使用i iris.Party. new
	3. 如果是闭包，则需要声明route，如果是模板，则需要声明Party,并在子路由使用iris.Party
	*/

	/**
	admin模块
	*/
	mvc.New(app.PartyFunc("/admin", func(i iris.Party) {
		mvc.New(i.Party("/login")).Handle(new(admin.LoginController))
	}))
}
