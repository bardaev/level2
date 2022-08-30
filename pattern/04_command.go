package pattern

import "fmt"

/*
Описание:
	Объект команды заключает в себе само действие и его параметры
	У данного паттерна есть 4 основные абстракции:
		- Приемник - содержит бизнесс логику
		- Комманда - инкапсулирует приемник для вызова его действия
		- Вызывающий комманды - объект к которому привязывается команда
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

type Database struct{}

func (d Database) Select() {
	fmt.Println("Select record")
}

func (d Database) Insert() {
	fmt.Println("Insert record")
}

func (d Database) Update() {
	fmt.Println("Update record")
}

func (d Database) Delete() {
	fmt.Println("Delete record")
}

type Command interface {
	Execute()
}

type SelectCommand struct {
	Db Database
}

func (s SelectCommand) Execute() {
	s.Db.Select()
}

type InsertCommand struct {
	Db Database
}

func (i InsertCommand) Execute() {
	i.Db.Insert()
}

type UpdateCommand struct {
	Db Database
}

func (u UpdateCommand) Execute() {
	u.Db.Update()
}

type DeleteCommand struct {
	Db Database
}

func (d DeleteCommand) Execute() {
	d.Db.Delete()
}

type Developerr struct {
	Select Command
	Insert Command
	Update Command
	Delete Command
}

func (d Developerr) SelectRecord() {
	d.Select.Execute()
}

func (d Developerr) InsertRecord() {
	d.Insert.Execute()
}

func (d Developerr) UpdateRecord() {
	d.Update.Execute()
}

func (d Developerr) DeleteRecord() {
	d.Delete.Execute()
}
