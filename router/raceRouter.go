package racerouter

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	rabbitMQ "github.com/mottajunior/race-service/Infraestruture"
	. "github.com/mottajunior/race-service/models"
	. "github.com/mottajunior/race-service/repository"
	strategy "github.com/mottajunior/race-service/service"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

var dao = RaceDAO{}

func GetAll(w http.ResponseWriter, r *http.Request) {
	races, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, races)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	race, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid race ID")
		return
	}
	respondWithJson(w, http.StatusOK, race)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var race Race
	if err := json.NewDecoder(r.Body).Decode(&race); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	race.StatusCorrida = "Procurando corrida"
	race.ID = bson.NewObjectId()
	if err := dao.Create(race); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, race)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func UpdateState(w http.ResponseWriter, r* http.Request){
	params := mux.Vars(r)
	keys, ok := r.URL.Query()["status"]
	if !ok{
		log.Println("Parametro nao informado")
		return;
	}

	status := keys[0]
	var strategy  = getStrategy(status)
	if (!strategy.Run(params["id"])){
		respondWithError(w, http.StatusInternalServerError,"erro ao atulizar status da corrida")
		return
	}

	msg := createMessage(strategy)
	rabbitMQ.PostMessage(msg)
	fmt.Println("Mensagem postada na fila.")

	respondWithJson(w, http.StatusOK, params["id"])
}

func getStrategy(status string) strategy.SetStatusRaceStrategy{
	strategy, exists := strategy.AvaliableStatus[status]
	if !exists{
		log.Println("strategy não encontrado")
	}
	return strategy
}
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func createMessage(strategy strategy.SetStatusRaceStrategy) string{
	//objectType := reflect.TypeOf(strategy)
	//check type of strategy
	//request necessary for make message
	//return message
	return "message fake"

	//request example.
	//resp, err := http.Get("http://localhost:4000/api/v1/races")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer resp.Body.Close()
	//bodyBytes, err := ioutil.ReadAll(resp.Body)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println("retornou => " + string(bodyBytes))
}

