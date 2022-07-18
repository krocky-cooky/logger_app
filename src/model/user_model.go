package model

import (
	"github.com/krocky-cooky/logger_app/crypto"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model 
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
}

func CreateUser(username string, password string) error {
	passwordEncrypt, _ := crypto.PasswordEncrypt(password)
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()

	if result := db.Create(&User{Username: username, Password: passwordEncrypt}); result.Error != nil {
		return result.Error
	}
	
	return nil
}

func GetUser(username string) (User ,error) {
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()

	var user User
	err = db.First(&user, "username = ?",username).Error
	return user, err
}

func GetAllUser() ([]User, error) {
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()
	var users []User
	err = db.Find(&users).Error
	return users, err 
}
