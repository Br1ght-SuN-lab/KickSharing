package domen

//обычный item
type baseUnit struct {
	name string
	invNum string
}

func (b baseUnit) Name() string { return b.name}
func (b baseUnit) InventoryNumber() string { return b.invNum}

//электричество
type electric struct {
	dailyEnergy float64 //кВт*ч - делаю вещественным
}

func (e electric) DailyEnergy() float64 { return e.dailyEnergy}


//транспорт
type vehicleBase struct {
	baseUnit
	simplicity int
}

func (v vehicleBase) Simplicity() int { return v.simplicity}


