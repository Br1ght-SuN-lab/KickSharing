package service

import (
	"testing"

	"kicksharing/internal/domen"
)


func TestDefaultServiceCenterCheck(t *testing.T) {
	sc := NewDefaultServiceCenter(20.0) // известный лимит расхода — 20 кВт·ч

	cases := []struct {
		name    string
		device  domen.Transport
		IsError bool //жду ошибку (true), не жду ошибку (false)
	}{
		// корректные устройства
		{"корректный велосипед", domen.NewBike("Велосипед", "B-1", 8), false},
		{"корректный электросамокат", domen.NewEScooter("Электро", "ES-1", 7, 0.6), false},

		// имя и номер
		{"пустое имя", domen.NewBike("", "B-2", 8), true},
		{"номер без дефиса", domen.NewBike("Вел", "B2", 8), true},
		{"номер из одних цифр", domen.NewBike("Вел", "123", 8), true},

		// простота: вне диапазона и на границах
		{"простота ниже диапазона (0)", domen.NewBike("Вел", "B-3", 0), true},
		{"простота выше диапазона (11)", domen.NewBike("Вел", "B-4", 11), true},
		{"простота на нижней границе (1)", domen.NewBike("Вел", "B-5", 1), false},
		{"простота на верхней границе (10)", domen.NewBike("Вел", "B-6", 10), false},

		// энергия: выше лимита, ровно на лимите, скрытый потребитель
		{"расход выше лимита", domen.NewEScooter("Электро", "ES-2", 7, 31.0), true},
		{"расход ровно на лимите", domen.NewEScooter("Электро", "ES-3", 7, 20.0), false},

		// неэлектрический транспорт — проверка энергии пропускается
		{"велосипед без потребления энергии", domen.NewBike("Вел", "B-7", 5), false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := sc.Check(c.device)
			if (err != nil) != c.IsError {
				t.Errorf("Check(%s) = %v, ожидали наличие ошибки = %v", c.name, err, c.IsError)
			}
		})
	}
}
