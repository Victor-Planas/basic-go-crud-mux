package service

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"br.com.meli/m/v2/entities"
	"br.com.meli/m/v2/repositories"
)

func ListAll(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(repositories.HouseRepository)
}

func DeleteHouse(key int, w http.ResponseWriter) {
	_, ok := repositories.HouseRepository[key]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - No such house was found for this key!"))
	}
	delete(repositories.HouseRepository, key)

}

func UpdateHouse(key int, reqBody []byte, w http.ResponseWriter) {
	house, ok := repositories.HouseRepository[key]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - No such house was found for this key!"))
	} else {
		json.Unmarshal([]byte(reqBody), &house)
		repositories.HouseRepository[key] = house
	}
}

func AddHouse(reqBody []byte, w http.ResponseWriter) {
	house := entities.House{}
	json.Unmarshal([]byte(reqBody), &house)
	newKey := rand.Int()
	_, ok := repositories.HouseRepository[newKey]
	if !ok {
		repositories.HouseRepository[rand.Int()] = house
	} else {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("409 - There is already object with generated key, please try again!"))
	}

}

func GetById(key int, writer http.ResponseWriter) {
	house, ok := repositories.HouseRepository[key]
	if !ok {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - No such house was found for this key!"))
	} else {
		json.NewEncoder(writer).Encode(house)
	}

}
