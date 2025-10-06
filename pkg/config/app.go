package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectToDB() {
	d , err:=gorm.Open("mysql", "root:root@1@/golang_db?charset=utf8&parseTime=True&loc=Local")

	if err !=nil {
		panic("failed to connect database")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}