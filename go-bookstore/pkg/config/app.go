package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	if d, err := gorm.Open("mysql", "root:Aa123456@tcp(127.0.0.1:3306)/simplerest?charset=UTF8&parseTime=True&loc=Local"); err != nil {
		panic(err)
	} else {
		db = d
	}

}

func GetDB() *gorm.DB {
	return db
}
