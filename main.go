package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	. "github.com/mottajunior/race-service/config"
	. "github.com/mottajunior/race-service/repository"
	racerouter "github.com/mottajunior/race-service/router"
)

var dao = RaceDAO{}
var config = Config{}

func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	//TODO: Verificar quais endpoints necessarios, e acessar endpoints de (obter device token - enviar notificacao)
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/races", racerouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/races/{id}", racerouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/races/client/{id}", racerouter.GetByClientId).Methods("GET")
	r.HandleFunc("/api/v1/races/driver/{id}", racerouter.GetByDriverId).Methods("GET")
	r.HandleFunc("/api/v1/races", racerouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/races/{id}", racerouter.UpdateState).Methods("PUT")
	r.HandleFunc("/api/v1/races/{id}", racerouter.Delete).Methods("DELETE")	

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))

}




