package model

//先定义常量
const STATUS_ENABLE = 1
const STATUS_DISABLE = 1

//定义map常量
var Status = map[int]string{0: "否", 1: "是"}

//定义表结构数据
type Admin struct {
	Id     int
	Name   string
	Email  string
	Passwd string
	Status int
}
