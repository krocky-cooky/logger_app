package model 

import (

	_"github.com/krocky-cooky/logger_app/crypto"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

type Log struct {
	gorm.Model 
	User User
	UserId uint
	Username string `binding:"required"`
	Logname string `binding:"required"`
	Guid string
	LogDataSets []LogDataSet `gorm"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;`
	Variables []Variable `gorm"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;`
}

type LogDataSet struct {
	gorm.Model 
	LogID uint
}

type LogDataValue struct {
	gorm.Model 
	Variable Variable `gorm"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;`
	VariableId uint 
	LogDataSet LogDataSet `gorm"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;`
	LogDataSetId uint
	Data string
} 



type Variable struct {
	gorm.Model 
	Typeid int
	Name string
	LogID uint 
}


func CreateLog(logname string, username string, varNames []string, varTypes []int) error {
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()

	guid := xid.New().String()

	log := Log{Username: username, Logname: logname, Guid:guid}
	
	if result := db.Create(&log); result.Error != nil {
		return result.Error
	}

	for i := 0; i < len(varNames); ++i {
		variable := Variable{Name: varNames[i], Typeid: varTypes[i], LogId: log.ID}
		if result := db.Create(&variable); result.Error != nil {
			return result.Error
		}
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

func GetLogDataSets(logId uint) ([]LogDataSet, error) {
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()

	var logDataSets []LogDataSet
	err = db.Find(&logDataSets, "log_id = ?", logId)

	return logDataSets, err 

}

func GetLogDataValues(logDataSetId uint) ([]LogDataValue, error) {
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()

	var logDataValues []LogDataValue 
	err = db.Find(&logDataValues. "log_data_set_id = ?", logDataSetId)
	return logDataValues, err 
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
	db.Model(&Log{}).Preload("LogDatas").Preload("Variables").Where("username = ?",username).First(&log,logid)
	logdatas := Log.LogDatas 
	

	return logdatas
}

func GetVariables(username string, logid uint) []Variable {
	db := gormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()

	var log Log 
	db.Where("username = ?",username).First(&log,logid)
	var variables []Variables 
	db.Where("log_id = ?",log.ID).Find(&variables)

	return variables 
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


