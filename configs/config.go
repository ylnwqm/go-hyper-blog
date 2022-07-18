package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/kataras/iris/v12"
	"log"
	"time"
)

//声明一个地址使用*  不是&
var Setting *Config

//定义初始化参数
const (
	DefaultBindAddress          = "0.0.0.0"
	DefaultListenPort           = 8080
	DefaultMysqlHost            = "127.0.0.1"
	DefaultMysqlPort            = "3306"
	DefaultMysqlDatabase        = "test"
	DefaultMysqlUser            = "root"
	DefaultMysqlPassword        = "123456"
	DefaultMysqlCharset         = "utf8"
	DefaultMysqlMaxIdleConns    = 20
	DefaultMysqlMaxOpenConns    = 200
	DefaultMysqlMaxConnLifetime = 60
)

//定义需要的配置
type Config struct {
	Iris  iris.Configuration `yaml:"iris"`
	App   AppConfig          `yaml:"app"`
	Log   LogConfig          `yaml:"log"`
	Mysql MysqlConfig        `yaml:"mysql"`
}

type AppConfig struct {
	Debug       bool   `yaml:"debug"`
	Timezone    string `yaml:"timezone"`
	BindAddress string `yaml:"bindAddress"` // Server listening address
	Port        int    `yaml:"port"`        //
}

type MysqlConfig struct {
	Host            string        `yaml:"host"`
	Port            string        `yaml:"port"`
	Database        string        `yaml:"database"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	Charset         string        `yaml:"charset"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxConnLifetime time.Duration `yaml:"maxConnLifeTime"`
	SqlLog          bool          `yaml:"sqlLog"`
}

type LogConfig struct {
	//日志级别
	Level string `yaml:"level" default:"debug"`
	//日志存储路径
	Path string `yaml:"path" default:"storage/logs"`
	//日志文件名称
	FileName string `yaml:"filename" default:"log"`
}

//首先定义初始化
/**
初始化接受path 配置文件路径，如果没有则初始化配置
*/
func Init(path string) {
	//所以这里setting赋值时 需要将Config转为地址
	Setting = &Config{}
	set := *Setting
	fmt.Printf("student=%+v\n", set)
	if path == "" {
		Setting.initDefaultConfig()
		return
	}
	//载入配置
	Setting.parse(path)
}

// Resolve profile  载入配置
//这里并没有解决configor.Load是如何载入配置的
func (config *Config) parse(path string) {
	err := configor.Load(config, path)
	if err != nil {
		log.Fatalf("Unable to resolve profile#%s", err.Error())
	}
}

func (config *Config) initDefaultConfig() {
	config.App.Debug = false
	config.App.Timezone = "UTC"
	config.App.BindAddress = DefaultBindAddress
	config.App.Port = DefaultListenPort
	//Default MySQL configuration
	config.Mysql.Host = DefaultMysqlHost
	config.Mysql.Port = DefaultMysqlPort
	config.Mysql.Database = DefaultMysqlDatabase
	config.Mysql.User = DefaultMysqlUser
	config.Mysql.Password = DefaultMysqlPassword
	config.Mysql.Charset = DefaultMysqlCharset
	config.Mysql.MaxIdleConns = DefaultMysqlMaxIdleConns
	config.Mysql.MaxOpenConns = DefaultMysqlMaxOpenConns
	config.Mysql.MaxConnLifetime = DefaultMysqlMaxConnLifetime
	config.Mysql.SqlLog = false
	config.Log.Path = "/storage/logs"
	config.Log.Level = "error"
	config.Log.FileName = "iris-demo-new"
}
