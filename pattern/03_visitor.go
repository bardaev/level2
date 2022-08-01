package pattern

import "fmt"

/*
Описание:
	Паттерн посетитель позволяет добавить к структуре поведение без изменения самой структуры.

Применение:
	Когда над объектом необходимо выполнять некоторые несвязанные между собой операции, не затрагивая его исходную структуру.

Плюсы:
	- Упрощает добавление операций

Минусы:
	- Не подходит для классов, которые часто изменяются
	- Может привести к нарушению инкапсуляции элементов
*/

type Employee interface {
	FullName()
	// Должен быть метод, который принимает объект типа Visitor
	Accept(Visitor)
}

type Developer struct {
	FirstName string
	LastName  string
	Income    float32
	Age       int
}

func (d Developer) FullName() {
	fmt.Println("Developer ", d.FirstName, " ", d.LastName)
}

type Director struct {
	FirstName string
	LastName  string
	Income    float32
	Age       int
}

func (d Director) FullName() {
	fmt.Println("Director", d.FirstName, " ", d.LastName)
}

type Visitor interface {
	VisitDeveloper(d Developer)
	VisitorDirector(d Director)
}

type CalculIncome struct {
	bonusRate int
}

// Методы Visitor
func (c CalculIncome) VisitDeveloper(d Developer) {
	fmt.Println(d.Income + d.Income*float32(c.bonusRate)/100)
}

func (c CalculIncome) VisitorDirector(d Director) {
	fmt.Println(d.Income + d.Income*float32(c.bonusRate)/100)
}

// Реализация метода, принимающего объект Visitor
func (d Developer) Accept(v Visitor) {
	v.VisitDeveloper(d)
}

func (d Director) Accept(v Visitor) {
	v.VisitorDirector(d)
}
