package service

import (
	"br.com.meli/m/v2/entities"
	"br.com.meli/m/v2/repositories"
	"encoding/json"
	"math/rand"
	"net/http"
)

func ListAll(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(repositories.HouseRepository)
}

func DeleteHouse(key int, w http.ResponseWriter) {
	_,ok := repositories.HouseRepository[key]
	if ok == false {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - No such house was found for this key!"))
	}
	delete(repositories.HouseRepository, key)

}

func UpdateHouse(key int, reqBody []byte) {
	house := repositories.HouseRepository[key]
	json.Unmarshal([]byte(reqBody), &house)
	repositories.HouseRepository[key] = house
}

func AddHouse(reqBody []byte) {
	house := entities.House{}
	json.Unmarshal([]byte(reqBody), &house)
	repositories.HouseRepository[rand.Int()] = house
}

func GetById(key int, writer http.ResponseWriter) {
	json.NewEncoder(writer).Encode(repositories.HouseRepository[key])
}