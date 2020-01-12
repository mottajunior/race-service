package repository

import (
	"log"

	. "github.com/mottajunior/race-service/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RaceDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "races"
)

func (m *RaceDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *RaceDAO) GetAll() ([]Race, error) {
	var races []Race
	err := db.C(COLLECTION).Find(bson.M{}).All(&races)
	return races, err
}

func (m *RaceDAO) GetByID(id string) (Race, error) {
	var race Race
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&race)
	return race, err
}

func (m *RaceDAO) Create(race Race) error {
	err := db.C(COLLECTION).Insert(&race)
	return err
}

func (m *RaceDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *RaceDAO) Update(id string, race Race) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &race)
	return err
}
