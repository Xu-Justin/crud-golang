package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Justin123!"
	dbName := "crudgolang"
	d, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err) 
	} else {
		db = d
		fmt.Println("Connected to DB.") 
	}
}

func GetDB() *gorm.DB {
	return db
}
