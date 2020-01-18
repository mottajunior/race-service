package strategy

import (
  "log"
)

type SetObjectInTransportStrategy struct {}

func (strategy SetObjectInTransportStrategy)  Run(Id string) bool {
  race, err := dao.GetByID(Id)
  if err != nil {
    log.Println("erro ao buscar corridas.")
    return false
  }
  race.StatusCorrida = "Objeto em transporte"
  if err := dao.Update(Id,race); err != nil {
    log.Println("erro ao atualizar o status da corrida")
    return false
  }
  return true
}