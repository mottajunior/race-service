package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Race struct {	
	Id      bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	IdMotorista		int `bson: "id_motorista" json: "id_motorista"`
	IdCliente		int `bson: "id_cliente" json: "id_cliente"`
	LocalOrigem		string `bson: "local_origem" json: "local_origem"`
	LocalDestino	string `bson: "local_destino" json: "local_destino"`
	Descricao		string `bson: "descricao" json: "descricao"`
	StatusCorrida	string `bson: "status_corrida" json: "status_corrida"`
	DtInicioCorrida time.Time `bson: "dt_inicio_corrida" json: "dt_inicio_corrida"`
	DtFinalCorrida	time.Time `bson: "dt_final_corrida" json: "dt_final_corrida"`
	ValorCorrida float64 `bson: "valor_corrida" json: "valor_corrida"`
}

