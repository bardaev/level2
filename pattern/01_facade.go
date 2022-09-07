package pattern

import "fmt"

/*
Описание:
	Фасад - паттерн, который инкапсулирует сложную логику и предоставляет упрощенный интерфейс.

Применение:
	Может применяться, когда, например, есть библиотека со множеством объектов,
	которые необходимо инициализировать, следить за правильным порядком зависимостей и т.п.,
	и вся эта логика оборачивается в простые методы, которые инкапсулируют всю сложную работу.

Плюсы:
	- Изоляция компонентов от сложной логики

Минусы:
	- Привязка объекта фасада ко всем классам программы
*/

type SubsystemA struct{}

func (s *SubsystemA) A1() {
	fmt.Println("A1")
}

type SubsystemB struct{}

func (s *SubsystemB) B1() {
	fmt.Println("B1")
}

type SubsystemC struct{}

func (s *SubsystemC) C1() {
	fmt.Println("C1")
}

type Facade struct {
	SubsystemA SubsystemA
	SubsystemB SubsystemB
	SubsystemC SubsystemC
}

func (f *Facade) Operation1() {
	f.SubsystemA.A1()
	f.SubsystemB.B1()
	f.SubsystemC.C1()
}

func (f *Facade) Operation2() {
	f.SubsystemB.B1()
	f.SubsystemC.C1()
}
