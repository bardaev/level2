package pattern

import "fmt"

/*
Описание:
	Состояние - это паттерн, который позволяет менять поведение оюъекта после смены его состояния.

Применение:
	Применяется когда объект может иметь много различных состояний.

Плюсы:
	- Меняет поведение объекта в зависимости от состояния

Минусы:
	- При добавлении нового состояния необходимо добавлять реализацию во всех поведениях
*/

// Методы, которые будут реализововать состояния и главный объект
type State interface {
	Create() error
	Run() error
	Block() error
	Terminate() error
}

// Объект с состояниями
type Thread struct {
	new        State
	runnable   State
	blocked    State
	terminated State

	currentState State
}

// Конструктор
func Newthread() *Thread {
	var t *Thread = &Thread{}

	new := &NewThread{
		thread: t,
	}
	runnable := &Runnable{
		thread: t,
	}
	blocked := &Blocked{
		thread: t,
	}
	terminated := &Terminate{
		thread: t,
	}

	t.setState(new)
	t.new = new
	t.runnable = runnable
	t.blocked = blocked
	t.terminated = terminated

	return t
}

func (t *Thread) Create() error {
	return t.currentState.Create()
}

func (t *Thread) Run() error {
	return t.currentState.Run()
}

func (t *Thread) Block() error {
	return t.currentState.Block()
}

func (t *Thread) Terminate() error {
	return t.currentState.Terminate()
}

func (t *Thread) setState(s State) {
	t.currentState = s
}

// Новый поток
type NewThread struct {
	thread *Thread
}

func (n *NewThread) Create() error {
	if n.thread.currentState != nil {
		return fmt.Errorf("Поток уже существует")
	}
	n.thread.setState(n.thread.new)
	return nil
}

func (n *NewThread) Run() error {
	n.thread.setState(n.thread.runnable)
	return nil
}

func (n *NewThread) Block() error {
	return fmt.Errorf("Поток не запущен")
}

func (n *NewThread) Terminate() error {
	return fmt.Errorf("Поток не запущен")
}

// Поток заблокирован
type Runnable struct {
	thread *Thread
}

func (n *Runnable) Create() error {
	return fmt.Errorf("Поток уже создан")
}

func (n *Runnable) Run() error {
	if n.thread.currentState == n.thread.runnable {
		return fmt.Errorf("Поток уже запущен")
	}
	n.thread.setState(n.thread.runnable)
	return nil
}

func (n *Runnable) Block() error {
	if n.thread.currentState == n.thread.runnable {
		n.thread.setState(n.thread.blocked)
		fmt.Println("Поток заблокирован")
		return nil
	}
	return fmt.Errorf("Поток не запущен")
}

func (n *Runnable) Terminate() error {
	n.thread.setState(n.thread.terminated)
	fmt.Println("Поток завершен")
	return nil
}

// Поток запущен
type Blocked struct {
	thread *Thread
}

func (n *Blocked) Create() error {
	return fmt.Errorf("Поток уже создан")
}

func (n *Blocked) Run() error {
	n.thread.setState(n.thread.runnable)
	return nil
}

func (n *Blocked) Block() error {
	return fmt.Errorf("Поток уже заблокирован")
}

func (n *Blocked) Terminate() error {
	n.thread.setState(n.thread.terminated)
	fmt.Println("Поток завершен")
	return nil
}

// Поток завершен
type Terminate struct {
	thread *Thread
}

func (n *Terminate) Create() error {
	return fmt.Errorf("Поток завершен")
}

func (n *Terminate) Run() error {
	return fmt.Errorf("Поток завершен")
}

func (n *Terminate) Block() error {
	return fmt.Errorf("Поток завершен")
}

func (n *Terminate) Terminate() error {
	return fmt.Errorf("Поток завершен")
}
