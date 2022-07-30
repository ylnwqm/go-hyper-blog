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
func (l *LoginController) Post() mvc.Result {
	username := l.Ctx.FormValue("username")
	passwd := l.Ctx.FormValue("passwd")

	l.LoginService.GetUser(username, passwd)
	/*if check == false {
		return mvc.Response{
			Path: "admin/login",
		}
	}*/

	//加入session
	//println(u)

	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title": "Profile of ",
			"User":  1,
		},
	}
}
