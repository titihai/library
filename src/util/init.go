package util

import (
	"github.com/jinzhu/gorm"
	//
	_ "github.com/go-sql-driver/mysql"
	"github.com/kdada/tinygo/config"
	"github.com/sirupsen/logrus"
)

//Logger 全局日志器
var Logger *logrus.Logger
var db *gorm.DB

func init() {
	// 读取配置
	cfg, err := config.NewConfig("ini", "./config/library.cfg")
	checkErr(err)
	//section := cfg.GlobalSection()
	// 初始化数据库连接
	dbSec, ok := cfg.Section("dbSource")
	if !ok {
		panic("config dbsource section is not exist")
	}
	librarySource, err := dbSec.String("library")
	checkErr(err)
	db, err = gorm.Open("mysql", librarySource)
	checkErr(err)

	// 邮件发送配置
	mailSec, ok := cfg.Section("mail")
	if !ok {
		panic("config mail section is not exist")
	}
	EmailFrom, err = mailSec.String("EmailFrom")
	checkErr(err)
	EmailAccount, err = mailSec.String("EmailAccount")
	checkErr(err)
	EmailPassword, err = mailSec.String("EmailPassword")
	checkErr(err)
	EmailPort, err = mailSec.Int("EmailPort")
	checkErr(err)
	EmailServer, err = mailSec.String("EmailServer")
	checkErr(err)
	// 初始化全局日志器
	Logger = logrus.New()
	Logger.Level = logrus.DebugLevel
	//Logger.Formatter = &logrus.JSONFormatter{}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
