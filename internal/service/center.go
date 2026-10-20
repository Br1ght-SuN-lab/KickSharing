package service

import (
	"kicksharing/internal/domen"
	"regexp"
	"fmt"
)

//формат номер word-number
var invNumPattern = regexp.MustCompile(`^[A-Za-zА-Яа-я]+-\d+$`)

type ServiceCenter interface {
	Check(t domen.Transport) error
}

type DefaultServiceCenter struct {
	maxDailyEnergy float64
}

func NewDefaultServiceCenter(maxEnergy float64) *DefaultServiceCenter {
	return &DefaultServiceCenter{
		maxEnergy,
	}
}


func (s *DefaultServiceCenter) Check(t domen.Transport) error {
	if t.Name() == "" {
		return fmt.Errorf("пустое наименование")
	}
	if !invNumPattern.MatchString(t.InventoryNumber()) {
		return fmt.Errorf("неверный формат номера: %q", t.InventoryNumber())
	}
	if simple := t.Simplicity(); simple < 1 || simple > 10 {
		return fmt.Errorf("простота вне диапазона 1–10: %d", simple)
	}
	if ec, ok := t.(domen.EnergyConsumer); ok {
		if ec.DailyEnergy() > s.maxDailyEnergy {
			return fmt.Errorf("расход %.1f превышает лимит %.1f", ec.DailyEnergy(), s.maxDailyEnergy)
		}
	}
	return nil
}