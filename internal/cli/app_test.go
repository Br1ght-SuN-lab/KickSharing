package cli

import (
	"bytes"
	"strings"
	"testing"

	"kicksharing/internal/service"
)

// Интеграционный тест консольного слоя: подаём команды через строку,
// перехватываем вывод в буфер и проверяем ключевые строки.
func TestConsoleAppRun(t *testing.T) {
	input := strings.Join([]string{
		"-help",
		"-add-transport bike Велосипед B-1 8",
		"-add-transport escooter Бракованный ES-2 7 9.0", // отклонит техосмотр
		"-add-transport bike Плохая B-3 notanumber",       // ошибка парсинга
		"-add-transport unknownkind Икс X-1 5",            // неизвестный вид
		"-add-item helmet Шлем H-1",
		"-bogus-command",                                   // неизвестная команда
		"-count",
		"-total-energy",
		"-beginner-amount",
		"-list",
		"-exit",
	}, "\n")

	var out bytes.Buffer
	park := service.NewPark(service.NewDefaultServiceCenter(5.0))
	app := NewConsoleApp(strings.NewReader(input), &out, park)

	if err := app.Run(); err != nil {
		t.Fatalf("Run вернул ошибку: %v", err)
	}

	got := out.String()
	for _, want := range []string{
		"Список команд",              // -help отработал
		"принято: B-1",               // транспорт принят
		"не принято:",                // бракованный отклонён техосмотром
		"Неизвестный тип транспорта", // unknownkind
		"Неизвестная команда",        // -bogus-command
		"принято: H-1",               // вещь принята
		"Всего на балансе: 2",        // B-1 + H-1
	} {
		if !strings.Contains(got, want) {
			t.Errorf("вывод не содержит %q", want)
		}
	}
}
