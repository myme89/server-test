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

// func GetGroupTypeInfo() ([]model.GroupTypeInfo, error) {

// 	rows, err := db.Query(`SELECT * FROM public.group_type `)

// 	var (
// 		ID          int
// 		Tittle      string
// 		Description string
// 	)

// 	if err != nil {
// 		log.Error(err)
// 		return []model.GroupTypeInfo{}, err
// 	}

// 	defer rows.Close()
// 	var groupTypeAllInfo []model.GroupTypeInfo
// 	for rows.Next() {
// 		err := rows.Scan(&ID, &Tittle, &Description)
// 		if err != nil {
// 			log.Error(err)
// 			return []model.GroupTypeInfo{}, err
// 		}
// 		groupTypeOneInfo := model.GroupTypeInfo{ID: ID, Tittle: Tittle, Description: Description}

// 		groupTypeAllInfo = append(groupTypeAllInfo, groupTypeOneInfo)

// 	}
// 	return groupTypeAllInfo, nil
// }

// func GetRestaurantInfo() ([]model.RestaurantInfo, error) {

// 	rows, err := db.Query(`SELECT * FROM public.restaurant `)

// 	var (
// 		ID               int
// 		IDGroupType      int
// 		Tittle           string
// 		Rating           int
// 		Genre            string
// 		Address          string
// 		ShortDescipstion string
// 		Dishes           []model.MenuRestaurantInfo
// 		Long             float32
// 		Lat              float32
// 	)

// 	if err != nil {
// 		log.Error(err)
// 		return []model.RestaurantInfo{}, err
// 	}

// 	defer rows.Close()
// 	var RestaurantAllInfo []model.RestaurantInfo
// 	for rows.Next() {
// 		err := rows.Scan(&ID, &IDGroupType, &Tittle, &Rating, &Genre, &Address, &ShortDescipstion, &Long, &Lat)
// 		if err != nil {
// 			log.Error(err)
// 			return []model.RestaurantInfo{}, err
// 		}
// 		Dishes, err = GetMenuOneRestaurant(ID)
// 		RestaurantOneInfo := model.RestaurantInfo{ID: ID, IDGroupType: IDGroupType, Tittle: Tittle,
// 			Rating: Rating, Genre: Genre, Address: Address,
// 			ShortDescription: ShortDescipstion, Dishes: Dishes, Long: Long, Lat: Lat}

// 		RestaurantAllInfo = append(RestaurantAllInfo, RestaurantOneInfo)

// 	}
// 	return RestaurantAllInfo, nil
// }

// func GetMenuOneRestaurant(IDRestaurant int) ([]model.MenuRestaurantInfo, error) {
// 	rows, err := db.Query(`SELECT menu_restaurant.id,  menu_restaurant.name, menu_restaurant.description, menu_restaurant.price, menu_restaurant.image
// 							FROM menu_restaurant
// 							WHERE menu_restaurant.id_restaurant= $1`, IDRestaurant)

// 	var (
// 		ID          int
// 		Name        string
// 		Description string
// 		Price       int
// 		Image       string
// 	)

// 	if err != nil {
// 		log.Error(err)
// 		return []model.MenuRestaurantInfo{}, err
// 	}

// 	defer rows.Close()

// 	var MenuRestaurant []model.MenuRestaurantInfo
// 	for rows.Next() {
// 		err := rows.Scan(&ID, &Name, &Description, &Price, &Image)
// 		if err != nil {
// 			log.Error(err)
// 			return []model.MenuRestaurantInfo{}, err
// 		}
// 		RestaurantOneInfo := model.MenuRestaurantInfo{ID: ID, Name: Name, Description: Description, Price: Price, Image: Image}

// 		MenuRestaurant = append(MenuRestaurant, RestaurantOneInfo)
// 	}
// 	return MenuRestaurant, err
// }

// // GetPassword
// func GetUserHashPassword(username string) (string, error) {
// 	var hashPassword string
// 	log.Info("start query db for user password")
// 	rows, err := db.Query(`SELECT password FROM public.users WHERE user_name = $1 `, username)
// 	if err != nil {
// 		log.Info("GetUserHashPassword: select db fail")
// 		return "", err
// 	}
// 	defer rows.Close()
// 	if rows.Next() {
// 		if err := rows.Scan(&hashPassword); err != nil {
// 			log.Info("GetUserHashPassword: Scan password fail")
// 			return "", err
// 		}
// 	}
// 	return hashPassword, nil
// }

// //Get Group Users

// func GetUserGroup(username string) (string, error) {
// 	var (
// 		userGroup string
// 		err       error
// 	)
// 	rows, err := db.Query(`SELECT groups.group_name
// 							FROM public.users
// 							JOIN public.group_users AS groups
// 							ON users.id_group_users = groups.id
// 							WHERE users.user_name = $1`, username)
// 	if err != nil {
// 		log.Error(err)
// 		return "", err
// 	}
// 	if rows.Next() {
// 		if err = rows.Scan(&userGroup); err != nil {
// 			log.Error(err)
// 			return "", err
// 		}
// 	}
// 	defer rows.Close()
// 	return userGroup, err
// }

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
