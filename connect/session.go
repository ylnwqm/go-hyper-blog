package connect

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

const cookieNameForSessionID = "session_id_cookie"

type SessionCtx struct {
	Ctx iris.Context
}

func (s SessionCtx) GetSession(app *iris.Application) {
	sess := sessions.New(sessions.Config{
		Cookie: cookieNameForSessionID,
		// CookieSecureTLS: true,
		AllowReclaim: true,
	})
	app.Use(sess.Handler())
	sess.Start(s.Ctx)
}
