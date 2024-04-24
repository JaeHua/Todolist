package dao

import (
	"ToolWeb/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

// InitMySQL

func InitMYSQL() (err error) {
	config := conf.InitConf()

	DB, err = gorm.Open("mysql", config.SqlConn)
	if err != nil {
		return

	}
	return DB.DB().Ping()
}
