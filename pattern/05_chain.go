package pattern

import (
	"fmt"
)

/*
Описание:
	Цепочка вызовов - это паттерн, позволяющий построить цепочку обработчиков. Запрос обрабатывается и передается от одного обработчика к другому.

Применение:
	Может использоваться при авторизациях. На вход подаются данные пользователя и по цепочке идет его обработка (проверка существования пользователя, проверка прав доступа и т.д.)

Плюсы:
	- Каждый обработчик в цепочке выполняется независимо от другого
	- Модификация обработчика не затрагивает другие

Минусы:
	- Запрос может быть никем не обработан
	- Запрос может не достичь цели
*/

type Handler interface {
	SendRequest(msg string)
}

type ConcreteHandler1 struct {
	Next Handler
}

func (h *ConcreteHandler1) SendRequest(msg string) {
	fmt.Println("Handler1: " + msg)
	if h.Next != nil {
		h.Next.SendRequest(msg)
	}
}

type ConcreteHandler2 struct {
	Next Handler
}

func (h *ConcreteHandler2) SendRequest(msg string) {
	fmt.Println("Handler2: " + msg)
	if h.Next != nil {
		h.Next.SendRequest(msg)
	}
}

type ConcreteHandler3 struct {
	Next Handler
}

func (h *ConcreteHandler3) SendRequest(msg string) {
	fmt.Println("Handler3: " + msg)
	if h.Next != nil {
		h.Next.SendRequest(msg)
	}
}
