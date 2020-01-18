package strategy

type SetStatusRaceStrategy interface {
  Run(Id string) bool
}

var AvaliableStatus = map[string]SetStatusRaceStrategy{
  "SetRaceFoundStrategy":         SetRaceFoundStrategy{},
  "SetObjectInTransportStrategy": SetObjectInTransportStrategy{},
  "SetRaceFinishedStrategy":      SetRaceFinishedStrategy{},
}
