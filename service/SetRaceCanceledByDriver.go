package strategy


import (
	"log"
	"time"
  . "github.com/mottajunior/race-service/models"
  )
  

type SetRaceCanceledByDriver struct {}

func (strategy SetRaceCanceledByDriver)  Run(Id string) (race Race, sucess bool){
	race, err := dao.GetByID(Id)
	if err != nil {
	  log.Println("erro ao buscar corridas.")
	  return race, false
	}
	race.StatusCorrida = "Corrida cancelada pelo motorista"
	race.DtFinalCorrida = time.Now()
	if err := dao.Update(Id,race); err != nil {
	  log.Println("erro ao atualizar o status da corrida")
	  return race,false
	}
	return race,true
  }