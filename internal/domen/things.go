package domen


type Helmet struct {
	baseUnit
}

func NewHelmet(name string, invNum string) *Helmet {
	return &Helmet{
		baseUnit{name, invNum},
	}
}

type DockStation struct {
	baseUnit
}

func NewDockStation(name string, invNum string) *DockStation {
	return &DockStation{
		baseUnit{name, invNum},
	}
}


type ChargingCabinet struct {
	baseUnit
	electric
}

func NewChargingCabinet(name string, invNum string, dailyEnergy float64) *ChargingCabinet {
	return &ChargingCabinet{
		baseUnit{name, invNum},
		electric{dailyEnergy},
	}
}
