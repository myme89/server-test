package services

import (
	"encoding/json"
	"net/http"
	"server-test/db"
	"server-test/model"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// func GetGroupType(response http.ResponseWriter, resquest *http.Request) {

// 	groupType, err := db.GetGroupTypeInfo()

// 	if err != nil {
// 		log.Error(err)
// 		response.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	var InfoGroup []model.GroupTypeInfo

// 	for i := 0; i < len(groupType); i++ {
// 		InfoGroup = append(InfoGroup, model.GroupTypeInfo{ID: groupType[i].ID, Tittle: groupType[i].Tittle, Description: groupType[i].Description})
// 	}
// 	response.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(response).Encode(InfoGroup)
// }

// func GetRestaurantInfo(response http.ResponseWriter, resquest *http.Request) {

// 	restaurantInfo, err := db.GetRestaurantInfo()

// 	if err != nil {
// 		log.Error(err)
// 		response.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	var InfoRestaurant []model.RestaurantInfo

// 	for i := 0; i < len(restaurantInfo); i++ {
// 		InfoRestaurant = append(InfoRestaurant, model.RestaurantInfo{ID: restaurantInfo[i].ID, IDGroupType: restaurantInfo[i].IDGroupType,
// 			Tittle: restaurantInfo[i].Tittle, Rating: restaurantInfo[i].Rating, Genre: restaurantInfo[i].Genre, Address: restaurantInfo[i].Address,
// 			ShortDescription: restaurantInfo[i].ShortDescription, Dishes: restaurantInfo[i].Dishes, Long: restaurantInfo[i].Long, Lat: restaurantInfo[i].Lat})
// 	}

// 	response.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(response).Encode(InfoRestaurant)
// }

// // GetLoginToken get jwt token
// func GetLoginToken(username string) (string, error) {
// 	//generate JWT
// 	var jwtString string
// 	// mySigningKey, err := helpers.ConfigGet("jwt", "signkey")
// 	// if err != nil {
// 	// 	log.Error("cannot get sign key for JWT , set to default")
// 	// 	mySigningKey = "weriwoxcr342f234"
// 	// }

// 	mySigningKey := "weriwoxcr342f234"
// 	var mapClaim jwt.StandardClaims
// 	mapClaim.ExpiresAt = time.Now().Add(time.Hour * 24 * 7).Unix()
// 	mapClaim.Issuer = username
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaim)
// 	jwtString, err := token.SignedString([]byte(mySigningKey))
// 	if err != nil {
// 		return "", err
// 	}
// 	return jwtString, nil
// }

//get, update, insert data

func GetData(response http.ResponseWriter, resquest *http.Request) {

	// InfoGroup := "nguyen trong nhat"

	dataInfo, err := db.GetData()

	if err != nil {
		log.Error(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var DataInfo []model.DataInfo

	for i := 0; i < len(dataInfo); i++ {
		DataInfo = append(DataInfo, model.DataInfo{Name: dataInfo[i].Name, FullName: dataInfo[i].FullName})
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(DataInfo)
}

func PostData(response http.ResponseWriter, resquest *http.Request) {

	Id := resquest.Header.Get("ID")
	Name := resquest.Header.Get("Name")
	FullName := resquest.Header.Get("FullName")
	id, _ := strconv.Atoi(Id)

	log.Info("nhatnt", id)
	log.Info("nhatnt", Name)
	log.Info("nhatnt", FullName)

	InfoGroup, err := db.PostData(id, Name, FullName)
	// InfoGroup := "data received : "

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(InfoGroup)
}

func UpdateData(response http.ResponseWriter, resquest *http.Request) {

	oldName := resquest.Header.Get("old_name")
	newName := resquest.Header.Get("new_name")
	newFullName := resquest.Header.Get("new_full_name")

	InfoGroup, err := db.UpdateData(oldName, newName, newFullName)
	// InfoGroup := "data received : "

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(InfoGroup)
}
