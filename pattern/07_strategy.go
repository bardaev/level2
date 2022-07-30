package pattern

/*
Описание:
	Паттерн стратегия определяет семейство схожих алгоритмов.
	Он позволяет помещать алгоритмы в собственный клас и менять их во время выполнения.

Применение:
	Применяется когда у одной и той же проблемы может быть несколько решений.

Плюсы:
	- Динамическая смена алгоритмов
	- Не использует наследование

Минусы:
	- внедряет дополнительные классы, что может усложнить код
*/

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

// Сложение
type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

// Умножение
type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}
