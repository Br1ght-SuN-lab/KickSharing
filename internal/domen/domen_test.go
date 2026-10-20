package domen

import "testing"

func TestConstructorsAndGetters(t *testing.T) {
	bike := NewBike("Велосипед", "B-1", 8)
	if bike.Name() != "Велосипед" || bike.InventoryNumber() != "B-1" || bike.Simplicity() != 8 {
		t.Errorf("Bike: неверные поля: %q %q %d", bike.Name(), bike.InventoryNumber(), bike.Simplicity())
	}

	es := NewEScooter("Электросамокат", "ES-1", 5, 0.6)
	if es.DailyEnergy() != 0.6 {
		t.Errorf("EScooter.DailyEnergy = %.2f, ожидали 0.6", es.DailyEnergy())
	}

	eb := NewEBike("Электровелосипед", "EB-1", 7, 0.4)
	if eb.Simplicity() != 7 || eb.DailyEnergy() != 0.4 {
		t.Errorf("EBike: %d %.2f", eb.Simplicity(), eb.DailyEnergy())
	}

	//проверяю на сопоставление интерфейсам
	var _ Inventoried = NewBike("Велосипед", "B-2", 8)
	var _ Transport = NewEScooter("Велосипед", "B-2", 8, 20)
	var _ EnergyConsumer = NewChargingCabinet("Велосипед", "B-2", 25)
}
