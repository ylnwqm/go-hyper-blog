package controller

import "github.com/kataras/iris/v12"

//这个有点类似于声明一个类 括号里引入需要的其他类
type AdController struct{}

//函数名前参数是将controller和方法做绑定
func (ad *AdController) ad(ctx iris.Context) {

}
