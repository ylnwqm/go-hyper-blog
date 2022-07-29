package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"go-hyper-blog/service/admin"
)

type LoginController struct {
	Ctx          iris.Context
	LoginService admin.LoginService
	Session      *sessions.Session
}

/**
登录页 admin/login
*/
func (l *LoginController) Get() mvc.Result {
	return mvc.View{
		Name: "admin/pages-login.html",
	}
}

//
func (l *LoginController) Post(ctx iris.Context) mvc.Result {
	println(ctx.Params())

	return mvc.View{
		Name: "admin/pages-login.html",
		Data: iris.Map{
			"Title": "Profile of ",
			"User":  1,
		},
	}
}
