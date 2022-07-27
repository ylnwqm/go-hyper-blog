package admin

import (
	"github.com/kataras/iris/v12"
	"go-hyper-blog/service/admin"
)

type LoginController struct {
	LoginService admin.LoginService
}

func (l *LoginController) login(ctx iris.Context) {
	println("用到了Login控制器")
	ctx.View("hello.html")
}
