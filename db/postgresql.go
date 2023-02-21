package db

import (
	sql "database/sql"
	"fmt"
	"server-test/model"
	"sync"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// const (

// )

var db *sql.DB
var once sync.Once
var err error

// type dataBase struct {
// 	db *sql.DB
// }

func Init() {
	once.Do(func() {
		// host := os.Getenv("POSTGRES_HOST")
		// port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		// user := os.Getenv("POSTGRES_USER")
		// password := os.Getenv("POSTGRES_PASSWORD")
		// dbname := os.Getenv("POSTGRES_DB")

		host := "localhost"
		port := 5432
		user := "postgres"
		password := "1"
		dbname := "postgres"

		// log.Info("nhatnt", host)

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err = sql.Open("postgres", psqlInfo)

		log.Info("nhatnt", psqlInfo)

		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully connected!")
	})
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
		log.Error(err)
		return []model.DataInfo{}, err
	}
	defer rows.Close()
	var DataInfo []model.DataInfo

	for rows.Next() {
		err := rows.Scan(&Id, &Name, &FullName)
		if err != nil {
			log.Error(err)
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
		log.Error(err)
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
		log.Error(err)
		return "error", err
	}

	// log.Info("trongnhat = ", a)
	return "Update Done", nil
}
