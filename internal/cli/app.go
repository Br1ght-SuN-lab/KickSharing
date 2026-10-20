package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"kicksharing/internal/service"
	"strconv"
)

// ConsoleApp — интерактивный слой. Не знает, откуда пришёл ввод
// (клавиатура, файл, что угодно) и куда идёт вывод — работает через io.Reader/io.Writer.
// Вся бизнес-логика живёт в park, ConsoleApp только читает команды и печатает результат.
type ConsoleApp struct {
	scanner *bufio.Scanner
	out io.Writer
	park *service.Park
}


//-------------------Создание приложения-------------------
func NewConsoleApp(in io.Reader, out io.Writer, park *service.Park) *ConsoleApp {
	return &ConsoleApp{
		scanner: bufio.NewScanner(in),
		out: out,
		park: park,
	}
}

// readLine читает следующую строку ввода, обрезая пробелы по краям.
func (a *ConsoleApp) readLine() (string, bool) {
	if !a.scanner.Scan() {
		return "", false
	}
	return strings.TrimSpace(a.scanner.Text()), true
}

//-------------------Сообщение-helper-------------------
func (a *ConsoleApp) printHelp() {
	fmt.Fprintln(a.out, "Список команд:")
	fmt.Fprintln(a.out, "-help")
	fmt.Fprintln(a.out, "    показать эту справку")
	fmt.Fprintln(a.out, "-add-transport <вид> <имя> <инв.номер> <простота> [энергия]")
	fmt.Fprintln(a.out, "    добавить транспорт")
	fmt.Fprintln(a.out, "    вид: bike | scooter | ebike | escooter")
	fmt.Fprintln(a.out, "    простота: целое 1-10; энергия (кВт·ч) — только для ebike/escooter")
	fmt.Fprintln(a.out, "-add-item <вид> <имя> <инв.номер> [энергия]")
	fmt.Fprintln(a.out, "    добавить вещь на баланс")
	fmt.Fprintln(a.out, "    вид: helmet | dockstation | chargingcabinet")
	fmt.Fprintln(a.out, "    энергия (кВт·ч) — только для chargingcabinet")
	fmt.Fprintln(a.out, "-count")
	fmt.Fprintln(a.out, "    количество единиц на балансе")
	fmt.Fprintln(a.out, "-total-energy")
	fmt.Fprintln(a.out, "    суммарный суточный расход энергии (кВт·ч)")
	fmt.Fprintln(a.out, "-beginner-amount")
	fmt.Fprintln(a.out, "    транспорт для новичков (простота >= 6)")
	fmt.Fprintln(a.out, "-list")
	fmt.Fprintln(a.out, "    вся техника и вещи с инвентарными номерами")
	fmt.Fprintln(a.out, "-exit")
	fmt.Fprintln(a.out, "    выход из программы")
	fmt.Fprintln(a.out, "Имя не должно содержать пробелов — используйте дефис (например, Велосипед-городской).")
}


//-------------------Парсинг Int-------------------
func (a *ConsoleApp) parseInt(s string) (int, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		err := fmt.Errorf("%s должно быть целым числом", s)
		return 0, err
	}
	return v, nil
}


//-------------------Парсинг Float-------------------
func (a *ConsoleApp) parseFloat(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		err := fmt.Errorf("%s должно быть вещественным числом", s)
		return 0.0, err
	}

	return v, nil
}


//-------------------добавление транспорт-------------------
func (a *ConsoleApp) cmdAddTransport(args []string) {
	if len(args) < 4 {
		fmt.Fprintln(a.out, "Недостаточно аргументов для команды -add-transport")
		return
	}

	kind := args[0]
	make, ok := transportRegistry[kind]
	if !ok {
		fmt.Fprintf(a.out, "Неизвестный тип транспорта: %s\n", kind)
		return
	}

	name := args[1]
	invNum := args[2]
	simplicity, err := a.parseInt(args[3])
	if err != nil {
		fmt.Fprintf(a.out, "%v\n", err)
		return
	}

	energy := 0.0
	if len(args) >= 5 {
		energy, err = a.parseFloat(args[4])
		if err != nil {
			fmt.Fprintf(a.out, "%v\n", err)
			return
		}
	}

	t := make(name, invNum, simplicity, energy)

	if err := a.park.AddTransport(t); err != nil {
		fmt.Fprintf(a.out, "не принято: %v\n", err)
	} else {
		fmt.Fprintf(a.out, "принято: %s\n", invNum)
	}
}


//-------------------------добавление item---------------------
func (a *ConsoleApp) cmdAddItem(args []string) {
	if len(args) < 3 {
		fmt.Fprintln(a.out, "Недостаточно аргументов для команды -add-item")
		return
	}

	kind := args[0]
	make, ok := itemRegistry[kind]
	if !ok {
		fmt.Fprintf(a.out, "Неизвестный тип объекта: %s\n", kind)
		return
	}

	name := args[1]
	invNum := args[2]

	energy := 0.0
	if len(args) >= 4 {
		parsed, err := a.parseFloat(args[3])
		if err != nil {
			fmt.Fprintf(a.out, "%v\n", err)
			return
		}
		energy = parsed
	}

	t := make(name, invNum, energy)

	if err := a.park.AddItem(t); err != nil {
		fmt.Fprintf(a.out, "не принято: %v\n", err)
	} else {
		fmt.Fprintf(a.out, "принято: %s\n", invNum)
	}
}


//-------------------отчёты-------------------
func (a *ConsoleApp) printCount() {
	fmt.Fprintf(a.out, "Всего на балансе: %d\n", a.park.Count())
}

func (a *ConsoleApp) printTotalEnergy() {
	fmt.Fprintf(a.out, "Суммарный расход энергии: %.2f кВт·ч\n", a.park.TotalEnergy())
}

func (a *ConsoleApp) printBeginners() {
	list := a.park.BeginnerAmount()
	fmt.Fprintf(a.out, "Для новичков (%d):\n", len(list))
	for _, t := range list {
		fmt.Fprintf(a.out, "  %s — %s (простота %d)\n", t.InventoryNumber(), t.Name(), t.Simplicity())
	}
}

func (a *ConsoleApp) printAll() {
	all := a.park.All()
	fmt.Fprintf(a.out, "Всего единиц: %d\n", len(all))
	for _, u := range all {
		fmt.Fprintf(a.out, "  %s — %s\n", u.InventoryNumber(), u.Name())
	}
}


func (a *ConsoleApp) Run() error {
	fmt.Fprintln(a.out, "Добро пожаловать в Kicksharing!")
	fmt.Fprintln(a.out, "Введите одну из команд (-help, -add-transport, -add-item, -count, -total-energy, -beginner-amount, -list, -exit):")

	for {
		line, ok := a.readLine()
		if !ok {
			return nil
		}

		fields := strings.Fields(line)
		if len(fields) == 0 {
			fmt.Fprintln(a.out, "Пустая команда, введите команду -help")
			continue
		}

		cmd, args := fields[0], fields[1:]
		switch cmd {
		case "-help":
			a.printHelp()
		case "-add-transport":
			a.cmdAddTransport(args)
		case "-add-item":
			a.cmdAddItem(args)
		case "-count":
			a.printCount()
		case "-total-energy":
			a.printTotalEnergy()
		case "-beginner-amount":
			a.printBeginners()
		case "-list":
			a.printAll()
		case "-exit":
			return nil
		default:
			fmt.Fprintf(a.out, "Неизвестная команда: %s , воспользуйтесь -help\n", cmd)
		}
	}
}
