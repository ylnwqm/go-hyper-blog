package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go-hyper-blog/configs"
	"go-hyper-blog/database"
	"go-hyper-blog/routes"
	"os"
	"path"
)

func newApp() *iris.Application {
	app := iris.New()
	//使用中间件  如果有的话
	//app.Use()
	//注册路由文件
	routes.InitApp(app)
	//定位yml文件
	currentDir, _ := os.Getwd()
	configFileDir := path.Join(currentDir, "config.yml")
	config.Init(configFileDir)
	//注册iris配置
	app.Configure(iris.WithConfiguration(config.Setting.Iris))

	//配置连接数据库
	db, err := database.ConnectDb(
		config.Setting.Mysql.Host,
		config.Setting.Mysql.Port,
		config.Setting.Mysql.User,
		config.Setting.Mysql.Password,
		config.Setting.Mysql.Database,
	)
	if err != nil {
		app.Logger().Fatalf("error while loading the tables: %v", err)
	}
	sqlDB, errDb := db.DB()
	if errDb != nil {
		app.Logger().Fatalf("can not close database: %v", errDb)
	}
	defer sqlDB.Close()
	/*db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:3306)/bis?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                     // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		app.Logger().Fatalf("error while loading the tables: %v", err)
	}*/

	/**
	1. 声明一个result map
	2. user 赋值model
	3. db.Model(&user) 指定model ，如果只有First 则为result的model 会报错
	4. Frist(&result)   将查询结果赋值给result
	*/
	/*result := map[string]interface{}{}
	var user model.Admin
	db.Model(&user).First(&result)
	fmt.Printf("%+v", result)*/

	//注册html
	tmpl := iris.HTML("./views", ".html")
	tmpl.Reload(true)
	//定义静态统一文件目录路径
	tmpl.AddFunc("static", func() string {
		return "/static/admin"
	})
	app.RegisterView(tmpl)

	//注册全局中间件  目前没有用到，取自其他项目
	//app.UseGlobal(middleware.IrisRequestHandler)
	//配置静态文件
	app.HandleDir("/static", "./static")
	//错误页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("admin/pages-404.html")
	})
	return app
}

func main() {
	app := newApp()
	err := app.Run(iris.Addr(fmt.Sprintf("%s:%d", config.Setting.App.BindAddress, config.Setting.App.Port)))
	if err != nil {
		app.Logger().Print("Server failed to start")
	}
}
