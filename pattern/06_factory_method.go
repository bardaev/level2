package pattern

import "fmt"

/*
Описание:
	Позволяет создавать объекты без указания точного типа.

Применение:
	Можно применять при расширении библиотек.
	Для создания нового объекта необходимо всего лишь указать его в методе где его конструирует фабричный метод.
	А пользователь будет пользоваться все тем же интерфейсом для создания объектов.

Плюсы:
	- Упрощает добавление новых сущностей
	- Реализует принцип открытости/закрытости
	- Выделяет создание объекта в одно место

Минусы:
	- С увеличением сущностей увеличивается метод для их создания
*/

func GetCar(name string) (iCar, error) {
	if name == "Mercedes" {
		return newMercedes(), nil
	} else if name == "Bmw" {
		return newBmw(), nil
	}
	return nil, fmt.Errorf("wrong type")
}

type iCar interface {
	setName(n string)
	getName() string
	setSpeed(s int)
	getSpeed() int
}

type Car struct {
	name  string
	speed int
}

func (c *Car) setName(n string) {
	c.name = n
}

func (c *Car) getName() string {
	return c.name
}

func (c *Car) setSpeed(s int) {
	c.speed = s
}

func (c *Car) getSpeed() int {
	return c.speed
}

type Mercedes struct {
	Car
}

func newMercedes() iCar {
	return &Mercedes{
		Car{name: "Mercedes", speed: 250},
	}
}

type Bmw struct {
	Car
}

func newBmw() iCar {
	return &Bmw{
		Car{name: "Bmw", speed: 300},
	}
}
