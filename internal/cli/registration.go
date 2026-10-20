package cli 

import (
	"kicksharing/internal/domen"
)

//ВАЖНО, что здесь нельзя сразу объекты создавать, потому что пользователь будет в RunTime вводить данные
//наша задача по мере поступления заполнять наш парк, поэтому делаем логику создания объекта

type transportFactory func(name, num string, simplicity int, energy float64) domen.Transport


var transportRegistry = map[string]transportFactory {
	"bike": func(n, num string, s int, _ float64) domen.Transport {return domen.NewBike(n,num,s)},
	"ebike": func(n, num string, s int, e float64) domen.Transport {return domen.NewEBike(n,num,s,e)},
	"scooter": func(n, num string, s int, _ float64) domen.Transport {return domen.NewScooter(n,num,s)},
	"escooter": func(n, num string, s int, e float64) domen.Transport {return domen.NewEScooter(n,num,s,e)},
}


type itemFactory func(name, num string, energy float64) domen.Inventoried


var itemRegistry = map[string]itemFactory {
	"chargingcabinet": func(n, num string, e float64) domen.Inventoried {return domen.NewChargingCabinet(n,num,e)},
	"helmet": func(n, num string, _ float64) domen.Inventoried {return domen.NewHelmet(n,num)},
	"dockstation": func(n, num string, _ float64) domen.Inventoried {return domen.NewDockStation(n,num)},
}