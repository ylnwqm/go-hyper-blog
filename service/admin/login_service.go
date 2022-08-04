package admin

import (
	"gorm.io/gorm"
)

/*type LoginService interface {
	GetUser(username string, passowrd string)
}*/

var DB *gorm.DB
var err error

type LoginService struct {
	//里边是具体要实现的方法
}

func (l *LoginService) GetUser(username string, password string) {

	println(DB)
	//admin := model.Admin{Name: "哈哈", Email: "aaaa", Passwd: "123456"}
	//DB.Create(&admin)

	//fmt.Printf("%v", admin)
	/*result := model.Admin
	fmt.Printf("值为 %#v", result)
	DB.First(&result)
	fmt.Printf("值为 %#v", result)*/
}
