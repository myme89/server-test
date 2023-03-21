package db

import (
	sql "database/sql"
	"fmt"
	"server-test/logs"
	"server-test/server-database/config"
	"server-test/server-database/model"
	"sync"

	_ "github.com/lib/pq"
)

// const (

// )

var db *sql.DB
var once sync.Once
var err error

// type dataBase struct {
// 	db *sql.DB
// }

func Init(config *config.Config) (*sql.DB, error) {
	once.Do(func() {

		// host := "localhost"
		// port := 5432
		// user := "postgres"
		// password := "1"
		// dbname := "postgres"

		// config := config.GetConfig()

		host := config.Sever.ServerPostgersSql.DBHost
		port := config.Sever.ServerPostgersSql.DBPort
		user := config.Sever.ServerPostgersSql.DBUserName
		password := config.Sever.ServerPostgersSql.DBPassword
		dbname := config.Sever.ServerPostgersSql.DBName

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err = sql.Open("postgres", psqlInfo)

		// log.Info("nhatnt", psqlInfo)

	})
	return db, err
}

// get, insert, update data

func GetData() ([]model.DataInfo, error) {
	rows, err := db.Query(`SELECT * FROM public.data_test `)

	var (
		Id       string
		Name     string
		FullName string
	)

	if err != nil {
		logs.Logger.Error("GetData - Query data Postgres error: ", err)
		return []model.DataInfo{}, err
	}
	defer rows.Close()
	var DataInfo []model.DataInfo

	for rows.Next() {
		err := rows.Scan(&Id, &Name, &FullName)
		if err != nil {
			logs.Logger.Error("GetData - Scan data Postgres error: ", err)
			return []model.DataInfo{}, err
		}
		TempDataInfo := model.DataInfo{Name: Name, FullName: FullName}
		DataInfo = append(DataInfo, TempDataInfo)
	}
	return DataInfo, nil
}

func PostData(id int, name, fullName string) (string, error) {

	_, err := db.Exec(`INSERT INTO data_test (id, name, full_name)
	VALUES ($1, $2, $3)`, id, name, fullName)

	if err != nil {
		// log.Error(err)
		logs.Logger.Error("PostData - Scan data Postgres error: ", err)
		return "error", err
	}

	// log.Info("trongnhat a = ", a)
	return "Post Done", nil
}

func UpdateData(oldName, newName, newFullname string) (string, error) {

	_, err := db.Exec(`UPDATE data_test
					   SET name = $1, full_name= $2
					   WHERE name = $3;`, newName, newFullname, oldName)

	if err != nil {
		// log.Error(err)
		logs.Logger.Error("PostData - Exec data Postgres error: ", err)
		return "error", err
	}

	// log.Info("trongnhat = ", a)
	return "Update Done", nil
}
