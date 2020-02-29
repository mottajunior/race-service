package racerouter

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	rabbitMQ "github.com/mottajunior/race-service/Infraestruture"
	. "github.com/mottajunior/race-service/models"
	. "github.com/mottajunior/race-service/repository"
	strategy "github.com/mottajunior/race-service/service"	
	"io/ioutil"
	"log"
	"net/http"
	"strconv"		
	"strings"
)


type ClienteType struct{
	Nome string
	Email string
	Senha string
	IdProfile int
	Cidade string
	Endereco string
	NumeroResidencia string
	Bairro string
	Telefone string
	Cnh string
	Cep string
	DeviceToken string
	Perfil string
	Veiculos string
	Id int
}

var dao = RaceDAO{}

func GetAll(w http.ResponseWriter, r *http.Request) {
	races, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, races)
}

func GetByClientId(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	race, err := dao.GetByClientID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid client ID")
		return
	}
	respondWithJson(w, http.StatusOK, race)

}

func GetByDriverId(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	race, err := dao.GetByDriverID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid client ID")
		return
	}
	respondWithJson(w, http.StatusOK, race)

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


// []salvar =>
// []encontrar pelos parametros =>
// []retornar o encontrado com id certo pro alex


	defer r.Body.Close()
	var race Race
	if err := json.NewDecoder(r.Body).Decode(&race); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	race.StatusCorrida = "Procurando corrida"		
	
	if err := dao.Create(race); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	fmt.Println("dados da corrida.",race)		

	savedRace,err := dao.GetByDescricao(race.Descricao)

	if err != nil {
		fmt.Println("deu erro")
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	


	fmt.Println("corrida salva, vai enviar notificacao.")
	SendNotificationForAllDrivers(savedRace)
	respondWithJson(w, http.StatusCreated, savedRace)

}

func SendNotificationForAllDrivers(race Race){
	var msg string
	ClientId :=   strconv.Itoa(race.IdCliente)
	Client := GetClientById(ClientId)
	DriversDeviceTokens := GetAllDriversToken()
	

	for _,token  := range DriversDeviceTokens{
		msg += token + "#!#"
	}

	//remove last " #!# "
	msg = strings.TrimSuffix(msg, "#!#")

	
	//tokens #!# tokens #!# tokens/message/race id.
	msg += "/cliente "+Client.Nome+ " deseja um corrida do local: "+race.LocalOrigem + " ate o destino: "+ race.LocalDestino+" deseja aceitar?/"+ race.Id.Hex()


	rabbitMQ.PostMessage(msg)
	fmt.Println("Mensagem postada na fila.", msg)

}


func GetAllDriversToken() []string{
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get("http://localhost:5000/api/usuario/token/all/motoristas")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}
	StringJson := string(bodyBytes)
	var arr []string

	//convert json to string array
	_ = json.Unmarshal([]byte(StringJson), &arr)
	return arr
}

func GetClientById(id string) ClienteType{
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get("http://localhost:5000/api/usuario/" + id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}
	StringJson :=  string(bodyBytes)

	//convert string json, to object
	var client ClienteType
	json.Unmarshal([]byte(StringJson), &client)
	return client


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
	race,sucess := strategy.Run(params["id"])

	if (status == "SetRaceFoundStrategy"){
		motorista, _ := r.URL.Query()["motorista"]		
		race.IdMotorista, _ = strconv.Atoi(motorista[0])		
		dao.Update(params["id"],race)
	}


	if (!sucess){
		respondWithError(w, http.StatusInternalServerError,"erro ao atulizar status da corrida")
		return
	}

	if (status != "SetObjectInTransportStrategy"){
		ClientId :=   strconv.Itoa(race.IdCliente)
		Client := GetClientById(ClientId)

		msg := createMessage(race)
		SendNotificationForOneUser(msg,Client.DeviceToken)
	}
	
	respondWithJson(w, http.StatusOK, params["id"])
}

func SendNotificationForOneUser(msg string, DeviceToken string){
	FormattedMessage := DeviceToken + "/" + msg
	rabbitMQ.PostMessage(FormattedMessage)
	fmt.Println("Mensagem postada na fila ==> " + FormattedMessage)

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

func createMessage(race Race) string {
	ClientId := strconv.Itoa(race.IdMotorista)
	Client := GetClientById(ClientId)
	if (race.StatusCorrida == "Corrida encontrada") {
		return "Motorista: " + Client.Nome + " Esta a caminho, para a coleta do objeto"
	} else if (race.StatusCorrida == "Corrida finalizada") {
		return "Motorista: " + Client.Nome + " Finalizou a corrida, com sucesso."
	}else{
		return "Motorista: " + Client.Nome + " Cancelou a corrida."
	}
}


