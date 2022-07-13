package model 

import (

	_"github.com/krocky-cooky/logger_app/crypto"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

type LogData struct {
	gorm.Model 
	Username string `form:"username" binding:"required"`
	LogID uint 
	Log Log `gorm"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;`
	Data float64 
} 

type Log struct {
	gorm.Model 
	Username string `binding:"required"`
	Logname string `binding:"required"`
	Guid string

}


func CreateLog(logname string, username string) error {
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()

	guid := xid.New().String()

	
	if result := db.Create(&Log{Username: username, Logname: logname, Guid:guid}); result.Error != nil {
		return result.Error
	}
	
	return nil
}

func GetLogs(username string) ([]Log, error) {
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()

	var logs []Log 
	err = db.Find(&logs, "username = ?",username).Error 
	
	return logs, err 
}

func RegisterLog(username string, logid uint, guid string, postdata float64) error {
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()

	var log Log
	db.Where("username = ?",username).Where("id = ?",logid).Where("guid = ?",guid).First(&log)
	result := db.Create(&LogData{Username: username, LogID: log.ID, Data: postdata})
	if result.Error != nil{
		return result.Error
	}
	return nil

}

func GetLogDatas(username string, logid uint) []LogData {
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()

	var log Log
	db.Where("username = ?",username).First(&log,logid)
	var logdatas []LogData
	db.Where("log_id = ?", log.ID).Find(&logdatas)
	return logdatas
}

func GetLogGuid(username string, logid uint) string {
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()
	var log Log 
	db.Where("username = ?",username).First(&log,logid)
	return log.Guid
}
