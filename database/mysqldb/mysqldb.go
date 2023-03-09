package mysqldb

import (
	"fmt"
	"server-test/config"
	"server-test/logs"
	"server-test/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitMySqlDb(config *config.Config) (*gorm.DB, error) {
	// fmt.Println("Go MySQL Tutorial")

	// fmt.Println(config)
	dsn := config.Sever.ServerMySql.DBUserName + ":" + config.Sever.ServerMySql.DBPassword + "@tcp" +
		"(" + config.Sever.ServerMySql.DBHost + ":" + config.Sever.ServerMySql.DBPort + ")/" + config.Sever.ServerMySql.DBName

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logs.Logger.Fatal("InitMySqlDb - Open Mysql error: ", err)
		fmt.Println(err)
	}

	db.AutoMigrate(&model.DataPost{})

	return db, err
}

func PostData(Id int, Name, FullName string) error {
	temp := model.DataPost{Id: Id, Name: Name, FullName: FullName}
	err := db.Create(temp).Error

	if err != nil {
		logs.Logger.Error("PostData - Create table mysql error: ", err)
		return err
	}
	return nil
}

func GetData(info *[]model.DataPost) error {

	fmt.Println("db= ", info)
	err := db.Find(info).Error
	if err != nil {
		logs.Logger.Error("PostData - Find data mysql error: ", err)
		return err
	}
	return nil
}
