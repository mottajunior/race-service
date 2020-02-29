package strategy

import (
  "log"
  "time"
. "github.com/mottajunior/race-service/models"
)

type SetRaceFoundStrategy struct {}

func (strategy SetRaceFoundStrategy)  Run(Id string) (race Race, sucess bool) {
  race, err := dao.GetByID(Id)
  if err != nil {
    log.Println("erro ao buscar corridas.")
    return race, false
  }
  race.StatusCorrida = "Corrida encontrada"
  race.DtInicioCorrida = time.Now()
  if err := dao.Update(Id,race); err != nil {
    log.Println("erro ao atualizar o status da corrida")
    return race,false
  }
  return race,true
}