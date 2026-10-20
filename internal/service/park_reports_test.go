package service

import (
	"testing"

	"kicksharing/internal/domen"
)

func TestAddItemAndCount(t *testing.T) {
	p := NewPark(&FakeServiceCenter{})
	if err := p.AddItem(domen.NewHelmet("Шлем", "H-1")); err != nil {
		t.Fatalf("AddItem: %v", err)
	}
	if err := p.AddItem(domen.NewHelmet("Шлем2", "H-1")); err == nil {
		t.Error("дубль номера H-1 должен отклоняться")
	}
	if p.Count() != 1 {
		t.Errorf("Count = %d, ждали 1", p.Count())
	}
}

func TestTotalEnergy(t *testing.T) {
	p := NewPark(&FakeServiceCenter{})
	_ = p.AddTransport(domen.NewEScooter("ЭС", "ES-1", 7, 0.6))
	_ = p.AddTransport(domen.NewBike("Вел", "B-1", 8)) // не потребляет
	_ = p.AddItem(domen.NewChargingCabinet("Шкаф", "C-1", 2.0))
	if got := p.TotalEnergy(); got != 2.6 {
		t.Errorf("TotalEnergy = %.2f, ждали 2.60", got)
	}
}

func TestBeginnerAndAllSorted(t *testing.T) {
	p := NewPark(&FakeServiceCenter{})
	_ = p.AddTransport(domen.NewScooter("Сам", "A-1", 5)) // отсеется из новичков
	_ = p.AddTransport(domen.NewBike("Вел", "B-1", 8))
	_ = p.AddItem(domen.NewHelmet("Шлем", "H-1"))

	beg := p.BeginnerAmount()
	if len(beg) != 1 || beg[0].InventoryNumber() != "B-1" {
		t.Errorf("BeginnerAmount неверно: %+v", beg)
	}

	all := p.All()
	want := []string{"A-1", "B-1", "H-1"} // отсортировано
	if len(all) != 3 {
		t.Fatalf("All вернул %d, ждали 3", len(all))
	}
	for i, num := range want {
		if all[i].InventoryNumber() != num {
			t.Errorf("All[%d] = %s, ждали %s", i, all[i].InventoryNumber(), num)
		}
	}
}
