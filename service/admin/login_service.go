package admin

import (
	"fmt"
	"github.com/kataras/iris/v12/sessions"
	"go-hyper-blog/connect"
	"go-hyper-blog/model"
	"time"
)

/*type LoginService interface {
	GetUser(username string, passowrd string)
}*/

type LoginService struct {
	//里边是具体要实现的方法
	Session   *sessions.Session //里边的具体实现需要先将其声明为地址才可以
	StartTime time.Time
}

func (l *LoginService) GetUser(username string, password string) {
	var adminModel model.Admin
	//声明返回的map
	result := map[string]interface{}{}
	//使用adminmodel 并且赋值给result
	//connect.GetDb().Model(&adminModel).Where("name = ? AND passwd = ?", username, password).Find(&result)
	//上面使用原生sql的方式，还有一种使用使用map 传递参数
	connect.DB.Model(&adminModel).Where(map[string]interface{}{"name": username, "password": password}).Find(&result)
	if len(result) == 0 {
		result["message"] = "账号密码错误"
	}

	fmt.Printf("%+v", result)
}
