package services

import (
	"encoding/json"
	"net/http"
	"server-test/database/db"
	"server-test/model"
	"strconv"

	log "github.com/sirupsen/logrus"
)

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
