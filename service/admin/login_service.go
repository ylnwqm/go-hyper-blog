package admin

import (
	"fmt"
	"go-hyper-blog/database"
	"go-hyper-blog/model"
)

/*type LoginService interface {
	GetUser(username string, passowrd string)
}*/

type LoginService struct {
	//里边是具体要实现的方法
}

func (l *LoginService) GetUser(username string, password string) {

	//admin := model.Admin{Name: "哈哈", Email: "aaaa", Passwd: "123456"}
	//DB.Create(&admin)

	//fmt.Printf("%v", admin)
	var result model.Admin
	fmt.Printf("值为 %#v", result)
	database.GetDb().First(&result)
	fmt.Printf("值为 %#v", result)
}
