package domen


//--------------------------Transport-----------------
type Bike struct {
	vehicleBase
}

func NewBike(name string, invNum string, simplicity int) *Bike {
	return &Bike{
		vehicleBase{baseUnit{name, invNum}, simplicity},
	}
}

type Scooter struct {
	vehicleBase
}

func NewScooter(name string, invNum string, simplicity int) *Scooter {
	return &Scooter{
		vehicleBase{baseUnit{name, invNum}, simplicity},
	}
}


//-----------------------Electric + Transport-----------
type EBike struct {
	vehicleBase
	electric
}

func NewEBike(name string, invNum string, simplicity int, dailyEnergy float64) *EBike {
	return &EBike{
		vehicleBase{baseUnit{name, invNum}, simplicity},
		electric{dailyEnergy},
	}
}

type EScooter struct {
	vehicleBase
	electric
}

func NewEScooter(name string, invNum string, simplicity int, dailyEnergy float64) *EScooter {
	return &EScooter{
		vehicleBase{baseUnit{name, invNum}, simplicity},
		electric{dailyEnergy},
	}
}