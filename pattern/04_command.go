package pattern

import "fmt"

/*
Описание:
	У данного паттерна есть 4 основные абстракции:
		- Приемник - содержит бизнесс логику (Класс Tv)
		- Комманда - инкапсулирует приемник для вызова его действия (Класс OnCommand, OffCommand)
		- Вызывающий комманды - объект к которому привязывается команда (Класс Button)
		- Клиент - в нем создается вызывающий комманды с приемником

Применение:
	Как следует из примера ниже, данный паттерн применяется в случаях когда необходимо отделить код обработчика от кнопки.
	Это позволяет отделить бизнес логику от кнопки.
	В противном случае пришлось бы писать на каждую кнопку свой обработчик, что породит дублирующий код на схожие действия.

Плюсы:
	- Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют

Минусы:
	- Усложняет код программы из-за множества дополнительных классов
*/

type Button struct {
	command command
}

func (b *Button) press() {
	b.command.execute()
}

type command interface {
	execute()
}

type OnCommand struct {
	device device
}

func (c *OnCommand) execute() {
	c.device.On()
}

type OffCommand struct {
	device device
}

func (c *OffCommand) execute() {
	c.device.On()
}

type device interface {
	On()
	Off()
}

type Tv struct {
	isRunning bool
}

func (t *Tv) On() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) Off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
