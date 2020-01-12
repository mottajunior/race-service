package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

//TODO: create logs with this struct, and alter status in race collections.
type RaceLogs struct {
	ID			bson.ObjectId `bson: "_id" json: "id"`
	CreatedAt	time.Time `bson: "dt_inicio_corrida" json: "dt_inicio_corrida"`
	IdRace		int `bson: "id_cliente" json: "id_cliente"`
	StatusAtual string `bson: "status_atual" json: "status_atual"`
}