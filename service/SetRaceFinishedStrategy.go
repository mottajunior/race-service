package strategy

import (
  . "github.com/mottajunior/race-service/repository"
  "log"
  . "github.com/mottajunior/race-service/models"
)
var dao = RaceDAO{}

type SetRaceFinishedStrategy struct {}


func (strategy SetRaceFinishedStrategy) Run(Id string) (race Race, sucess bool){
  race, err := dao.GetByID(Id)
  if err != nil {
    log.Println("erro ao buscar corridas.")
    return race, false
  }
  race.StatusCorrida = "Corrida finalizada"
  if err := dao.Update(Id,race); err != nil {
    log.Println("erro ao atualizar o status da corrida")
    return race, false
  }
  return race, true
}