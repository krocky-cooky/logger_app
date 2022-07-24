package model 

import (
	"fmt"

	_"github.com/krocky-cooky/logger_app/crypto"
	"gorm.io/gorm"
	"github.com/rs/xid"
)

type Log struct {
	gorm.Model 
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
	VariableID uint 
	LogDataSet LogDataSet `gorm"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;`
	LogDataSetID uint
	Data string
} 



type Variable struct {
	gorm.Model 
	Typeid int
	Name string
	LogID uint 
}


func CreateLog(logname string, username string, varNames []string, varTypes []int, db *gorm.DB) error {
	guid := xid.New().String()

	log := Log{Username: username, Logname: logname, Guid:guid}
	
	if result := db.Create(&log); result.Error != nil {
		return result.Error
	}

	for i := 0; i < len(varNames); i++ {
		variable := Variable{Name: varNames[i], Typeid: varTypes[i], LogID: log.ID}
		if result := db.Create(&variable); result.Error != nil {
			return result.Error
		}
	}
	
	return nil
}

func GetLogs(username string, db *gorm.DB) ([]Log, error) {
	

	var logs []Log 
	err := db.Find(&logs, "username = ?",username).Error 
	
	return logs, err 
}

func GetLog(logId uint, db *gorm.DB) (Log, error) {
	var log Log 
	err := db.Preload("LogDataSets").Preload("Variables").First(&log, logId).Error
	return log, err 
} 

func GetLogDataSets(logId uint, db *gorm.DB) ([]LogDataSet, error) {

	var logDataSets []LogDataSet
	err := db.Find(&logDataSets, "log_id = ?", logId).Error

	return logDataSets, err 

}

func GetLogDataValues(logDataSetId uint, db *gorm.DB) ([]LogDataValue, error) {
	var logDataValues []LogDataValue 
	err := db.Find(&logDataValues, "log_data_set_id = ?", logDataSetId).Error
	return logDataValues, err 
}

func GetVariable(variableId uint, db *gorm.DB) (Variable, error) {
	var variable Variable 
	err := db.First(&variable, variableId).Error
	return variable, err 
}

func GetVariables(logId uint, db *gorm.DB) ([]Variable, error) {
	var variables []Variable
	err := db.Where("log_id = ?", logId).Find(&variables).Error
	return variables, err 
}

func RegisterLog(username string, logid uint, guid string, postdata map[string]string, db *gorm.DB) error {

	var log Log
	err := db.Where("username = ?",username).Where("id = ?",logid).Where("guid = ?",guid).First(&log).Error
	if err != nil {
		return err 
	}
	logDataSet := LogDataSet{LogID: log.ID}
	
	if result := db.Create(&logDataSet);result.Error != nil{
		return result.Error
	}

	for key,val := range postdata {
		var variable Variable 
		err := db.Where("log_id = ?", logid).Where("name = ?", key).First(&variable).Error
		if err != nil{
			return err
		}
		logDataValue := LogDataValue{
			VariableID: variable.ID,
			LogDataSetID: logDataSet.ID,
			Data: val,
		}

		if result := db.Create(&logDataValue); result.Error != nil {
			return result.Error
		}
		
	}
	fmt.Print("log registered !!!!!")
	
	return nil

}





func GetLogGuid(username string, logid uint, db *gorm.DB) string {
	var log Log 
	db.Where("username = ?",username).First(&log,logid)
	return log.Guid
}


