package mysqldb

import (
	"database/sql"
	"fmt"
	"server-test/config"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sql.DB

func InitMySqlDb(config *config.Config) (*gorm.DB, error) {
	fmt.Println("Go MySQL Tutorial")

	fmt.Println(config)
	dsn := config.Sever.ServerMySql.DBUserName + ":" + config.Sever.ServerMySql.DBPassword + "@tcp" +
		"(" + config.Sever.ServerMySql.DBHost + ":" + config.Sever.ServerMySql.DBPort + ")/" + config.Sever.ServerMySql.DBName

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
