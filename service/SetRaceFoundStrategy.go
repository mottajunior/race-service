package strategy

import (
  "log"
  "time"
)

type SetRaceFoundStrategy struct {}

func (strategy SetRaceFoundStrategy)  Run(Id string) bool {
  race, err := dao.GetByID(Id)
  if err != nil {
    log.Println("erro ao buscar corridas.")
    return false
  }
  race.StatusCorrida = "Corrida encontrada"
  race.DtInicioCorrida = time.Now()
  if err := dao.Update(Id,race); err != nil {
    log.Println("erro ao atualizar o status da corrida")
    return false
  }
  return true
}