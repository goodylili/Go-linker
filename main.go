package main

import (
	"Golinker/config"
	"Golinker/controllers"
	"Golinker/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func initRouters() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", controllers.SetLink).Methods("POST")
	router.HandleFunc("/api/v1/{id}", utils.FetchHandleRequestsRedirects).Methods("GET")
	router.HandleFunc("/api/v1/", controllers.Ping).Methods("GET")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalln("There was a problem with the routers")
	}

}

func main() {
	config.Migrations()
	initRouters()
}
