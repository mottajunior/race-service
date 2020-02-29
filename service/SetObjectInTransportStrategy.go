package strategy

import (
  "log"
  . "github.com/mottajunior/race-service/models"
)

type SetObjectInTransportStrategy struct {}

func (strategy SetObjectInTransportStrategy)  Run(Id string) (race Race, sucess bool) {
  race, err := dao.GetByID(Id)
  if err != nil {
    log.Println("erro ao buscar corridas.")
    return race,false
  }
  race.StatusCorrida = "Objeto em transporte"
  if err := dao.Update(Id,race); err != nil {
    log.Println("erro ao atualizar o status da corrida")
    return race,false
  }
  return race,true
}