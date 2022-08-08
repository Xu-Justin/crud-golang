package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/Xu-Justin/crud-golang/pkg/config"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age int `json:"age"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func Get() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetById(id int64) *User {
	var user User
	db.Where("ID=?", id).Find(&user)
	return &user
}

func Create(user *User) *User {
	db.NewRecord(user)
	db.Create(user)
	return user
}

func UpdateName(id int64, name string) *User {
	var user User
	status := db.Where("ID=?", id).Find(&user)
	if !errors.Is(status.Error, gorm.ErrRecordNotFound) {
		user.Name = name
		db.Save(&user)
	}
	return &user
}

func UpdateAge(id int64, age int) *User {
	var user User
	status := db.Where("ID=?", id).Find(&user)
	if !errors.Is(status.Error, gorm.ErrRecordNotFound) {
		user.Age = age
		db.Save(&user)
	}
	return &user
}

func Delete(id int64) *User {
	var user User
	status := db.Where("ID=?", id).Find(&user)
	if !errors.Is(status.Error, gorm.ErrRecordNotFound) {
		db.Delete(&user)
	}
	return &user
}