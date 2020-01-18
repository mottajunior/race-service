package strategy

import (
  . "github.com/mottajunior/race-service/repository"
  "log"
)
var dao = RaceDAO{}

type SetRaceFinishedStrategy struct {}


func (strategy SetRaceFinishedStrategy) Run(Id string) bool{
  race, err := dao.GetByID(Id)
  if err != nil {
    log.Println("erro ao buscar corridas.")
    return false
  }
  race.StatusCorrida = "Corrida finalizada"
  if err := dao.Update(Id,race); err != nil {
    log.Println("erro ao atualizar o status da corrida")
    return false
  }
  return true
}