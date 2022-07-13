package model

import (
	"os"
	_"log"
	"fmt"

	_"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_"github.com/joho/godotenv"
)

func gormConnect() *gorm.DB {
	
	//DBMS := os.Getenv("MYSQL_DBMS")
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	DBNAME := os.Getenv("MYSQL_DATABASE")
	
	CONNECT := USER + ":" + PASS + "@tcp(db)/" + DBNAME + "?parseTime=true"
	dialector := mysql.Open(CONNECT)
	fmt.Print(CONNECT)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		fmt.Print("error\n")
		panic(err.Error())
	}
	fmt.Print("db connected\n")
	return db
}

func DbInit() {
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()
	
	db.AutoMigrate(&User{},&Log{},&LogData{})
}





