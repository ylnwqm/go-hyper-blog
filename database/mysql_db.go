package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//这里最后跟类型string的意思是，这些参数的所有类型均为string 相当于定义所有参数的返回类型
func ConnectDb(host, port, username, password, dbname string) (*gorm.DB, error) {
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
