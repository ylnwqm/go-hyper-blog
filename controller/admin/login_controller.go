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

func (l *LoginController) Login() mvc.Result {

	// 绑定： {{.message}}　为　"Hello world!"
	/*println("就按就按家家爱")
	ctx.ViewData("message", "Hello world!")
	// 渲染模板文件： ./views/hello.html
	ctx.View("amdin/pages-login.html")*/

	println("用到了Login控制器")

	return mvc.View{
		Name: "amdin/pages-login.html",
		Data: iris.Map{
			"Title": "Profile of ",
			"User":  1,
		},
	}
}
