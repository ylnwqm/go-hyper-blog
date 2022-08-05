package database

import (
	"github.com/kataras/iris/v12"
	config "go-hyper-blog/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//这里最后跟类型string的意思是，这些参数的所有类型均为string 相当于定义所有参数的返回类型
func ConnectDb(app *iris.Application) (*gorm.DB, error) {
	dsn := config.Setting.Mysql.User + ":" + config.Setting.Mysql.Password + "@tcp(" + config.Setting.Mysql.Host + ":" + config.Setting.Mysql.Port + ")/" + config.Setting.Mysql.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		app.Logger().Fatalf("error while loading the tables: %v", err)
	}
	sqlDB, errDb := db.DB()
	if errDb != nil {
		app.Logger().Fatalf("can not close database: %v", errDb)
	}
	defer sqlDB.Close()
	return db, err
}

func GetDb() *gorm.DB {
	return DB
}
