package admin

import (
	"fmt"
	"go-hyper-blog/model"
	"gorm.io/gorm"
)

/*type LoginService interface {
	GetUser(username string, passowrd string)
}*/

type LoginService struct {
	//里边是具体要实现的方法
	DB *gorm.DB
}

func (l *LoginService) GetUser(username string, password string) []model.Admin {
	result := []model.Admin{}

	l.DB.Raw("select * from admin").Scan(&result)
	fmt.Println(result)
	return result
}
