package strategy

import (. "github.com/mottajunior/race-service/models")

type SetStatusRaceStrategy interface {
  Run(Id string) (race Race, sucess bool)
}

var AvaliableStatus = map[string]SetStatusRaceStrategy{
  "SetRaceFoundStrategy":         SetRaceFoundStrategy{},
  "SetObjectInTransportStrategy": SetObjectInTransportStrategy{},
  "SetRaceFinishedStrategy":      SetRaceFinishedStrategy{},
  "SetRaceCanceledByDriver":      SetRaceCanceledByDriver{},
}
