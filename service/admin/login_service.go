package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"go-hyper-blog/database"
	"go-hyper-blog/model"
)

/*type LoginService interface {
	GetUser(username string, passowrd string)
}*/

type LoginService struct {
	//里边是具体要实现的方法
	Session *sessions.Session //里边的具体实现需要先将其声明为地址才可以
	Ctx     iris.Context
}

func (l *LoginService) GetUser(username string, password string) {
	var adminModel model.Admin
	//声明返回的map
	result := map[string]interface{}{}
	//使用adminmodel 并且赋值给result
	//database.GetDb().Model(&adminModel).Where("name = ? AND passwd = ?", username, password).Find(&result)
	//上面使用原生sql的方式，还有一种使用使用map 传递参数
	database.GetDb().Model(&adminModel).Where(map[string]interface{}{"name": username, "passwd": password}).Find(&result)
	if len(result) == 0 {
		result["message"] = "账号密码错误"
		//return result
	}

	println(1111111111111111)
	l.Ctx.SetCookieKV("userId", "家居安静安静")
	println(22222)
	//common.PrettyPrint(result)
	//如果登录成功，设置session
	//result["message"] = "登录成功"
}
