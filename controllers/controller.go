package controllers

import (
	"br.com.meli/m/v2/service"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the homePage!")

}

func listAll(w http.ResponseWriter, r *http.Request) {
	service.ListAll(w)
}



func HandleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/list", listAll)
	myRouter.HandleFunc("/get/{id}", getById)
	myRouter.HandleFunc("/add", addHouse).Methods("POST")
	myRouter.HandleFunc("/update/{id}", updateHouse).Methods("PUT")
	myRouter.HandleFunc("/delete/{id}", deleteHouse).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func deleteHouse(writer http.ResponseWriter, request *http.Request) {
	keystr := mux.Vars(request)["id"]
	key,_ := strconv.Atoi(keystr)
	service.DeleteHouse(key, writer)

}

func updateHouse(writer http.ResponseWriter, request *http.Request) {
	keystr := mux.Vars(request)["id"]
	key,_ := strconv.Atoi(keystr)
	reqBody,_ := ioutil.ReadAll(request.Body)

	service.UpdateHouse(key, reqBody)
}

func addHouse(writer http.ResponseWriter, request *http.Request) {
	reqBody,_ := ioutil.ReadAll(request.Body)
	service.AddHouse(reqBody)
}

func getById(writer http.ResponseWriter, request *http.Request) {
	keystr := mux.Vars(request)["id"]
	key,_ := strconv.Atoi(keystr)
	service.GetById(key, writer)
}

