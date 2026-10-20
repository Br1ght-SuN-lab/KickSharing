package main

import (
	"fmt"
	"os"

	"go.uber.org/dig"

	"kicksharing/internal/cli"
	"kicksharing/internal/service"
)


// контейнер соберёт объекты сам, когда они понадобятся
func buildContainer() *dig.Container {
	c := dig.New()

	// провайдер обязан вернуть этот ServiceCenter интерфейс, а не *DefaultServiceCenter
	c.Provide(func() service.ServiceCenter {
		return service.NewDefaultServiceCenter(30.0) // 30.0 кВт·ч — лимит техосмотра
	})

	// Park просит ServiceCenter => DI-контейнер подставит то, что зарегистрировал выше
	c.Provide(service.NewPark)

	// ConsoleApp просит *Park — его даст предыдущий провайдер.
	c.Provide(newConsoleApp)

	return c
}


func newConsoleApp(park *service.Park) *cli.ConsoleApp {
	return cli.NewConsoleApp(os.Stdin, os.Stdout, park)
}


func main() {
	c := buildContainer()

	// Invoke просит у контейнера *ConsoleApp => а сам контейнер уже создает все вспомогательные объекты (Service, Park)
	err := c.Invoke(func(app *cli.ConsoleApp) error {
		return app.Run()
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", err)
		os.Exit(1)
	}
}
