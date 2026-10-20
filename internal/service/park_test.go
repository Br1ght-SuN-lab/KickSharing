package service

import (
	"fmt"
	"testing"

	"kicksharing/internal/domen"
)

// FakeServiceCenter — заглушка техосмотра, чтобы проверять только логику Park
type FakeServiceCenter struct {
	err error
}

func (f *FakeServiceCenter) Check(t domen.Transport) error {
	return f.err
}


func TestAddTransportPassed(t *testing.T) {
	p := NewPark(&FakeServiceCenter{err: nil}) //гарантированно пройдет Check

	err := p.AddTransport(domen.NewBike("Велосипед", "B-1", 8))

	if err != nil {
		t.Errorf("Ошибка добавления в парк")
	}

	if p.Count() != 1 {
		t.Errorf("Парк пустой")
	}
}


func TestAddTransportFailed(t *testing.T) {
	p := NewPark(&FakeServiceCenter{err: fmt.Errorf("Предмет не прошёл проверку сервисом")}) //результат заранее задается

	err := p.AddTransport(domen.NewBike("Велосипед", "B-1", 8))

	if err == nil {
		t.Errorf("Ошибка на стороне Park")
	}

	if p.Count() != 0 {
		t.Errorf("Ошибочное добавление предмета")
	}
}	

