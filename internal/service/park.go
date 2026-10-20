package service

import (
	"fmt"
	"kicksharing/internal/domen"
	"sort"
)

type Park struct {
	units map[string]domen.Inventoried
	checker ServiceCenter

}


func NewPark(checker ServiceCenter) *Park {
	return &Park{
		units: make(map[string]domen.Inventoried),
		checker: checker,
	}
}


//нужно всегда проверять номера на уникальность => отдельный метод
func (p *Park) add(u domen.Inventoried) error {
	if _, ok := p.units[u.InventoryNumber()]; ok {
		return fmt.Errorf("номер %s уже занят", u.InventoryNumber())
	}
	p.units[u.InventoryNumber()] = u
	return nil
}

func (p *Park) AddTransport(transport domen.Transport) error {
	err := p.checker.Check(transport)
	if err != nil {
		return err
	}

	return p.add(transport)
}


func (p *Park) AddItem(item domen.Inventoried) error {
	return p.add(item)
}


func (p *Park) Count() int {
	return len(p.units)
}


func (p *Park) BeginnerAmount() []domen.Transport {
	result := make([]domen.Transport, 0)
	for _, u := range p.units {
		vc, ok := u.(domen.Transport)
		if ok && vc.Simplicity() >= 6 {
			result = append(result, vc)
		}
	}

	//для тестов нужно сортировка
	sort.Slice(result, func(i, j int) bool {
		return result[i].InventoryNumber() < result[j].InventoryNumber()
	})

	return result 
}


func (p *Park) TotalEnergy() float64 {
	totalEnergy := float64(0)
	for _, unit := range p.units {
		ec, ok := unit.(domen.EnergyConsumer)
		if ok {
			totalEnergy += ec.DailyEnergy()
		}
	}

	return totalEnergy
}


func (p *Park) All() []domen.Inventoried {
	result := make([]domen.Inventoried, len(p.units))
	i := 0
	for _, u := range p.units {
		result[i] = u
		i++
	}

	sort.Slice(result, func (i, j int) bool {
		return result[i].InventoryNumber() < result[j].InventoryNumber()
	})

	return result 
}