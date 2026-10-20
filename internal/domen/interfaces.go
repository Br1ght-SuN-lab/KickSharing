package domen


type EnergyConsumer interface {
	DailyEnergy() float64
}


type Inventoried interface {
	Name() string 
	InventoryNumber() string
}


type Transport interface {
	Inventoried
	Simplicity() int
}