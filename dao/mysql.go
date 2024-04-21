package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

// InitMySQL

func InitMYSQL() (err error) {
	DB, err = gorm.Open("mysql", "root:jh529529@(localhost)/todolist?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return

	}
	return DB.DB().Ping()
}
