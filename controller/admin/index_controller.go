package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type IndexController struct {
	Ctx     iris.Context
	Session *sessions.Session
}
