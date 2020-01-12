package racerouter

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/mottajunior/race-service/models"
	. "github.com/mottajunior/race-service/repository"
	"gopkg.in/mgo.v2/bson"
)

var dao = RaceDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

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
	race.ID = bson.NewObjectId()
	if err := dao.Create(race); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, race)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var race Race
	if err := json.NewDecoder(r.Body).Decode(&race); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(params["id"], race); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": race.Descricao + " atualizado com sucesso!"})
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
