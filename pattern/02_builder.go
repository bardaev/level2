package pattern

/*
Описание:
	Builder применяется для поэтапного построение объекта.
	Это достигается засчет набора методов, конструирующие объект.
	Данные методы должны возвращать конструирующий объект, что позволяется конструировать объект цепочками вызовов.

Применение:
	Может исполльзоваться в классах конфигурациях, когда у объекта может быть много настроек.
	В отличие от создания объекта с помощью констурктора, Builder может инкапсулировать логику проверки входных данных, что не может конструктор.

Плюсы:
	- Позволяет пошагово создавать объекты
	- Изолирует код сборки объекта

Минусы:
	- При модификации создаваемого класса нужно менять сам Builder
*/

type Processor string

const (
	Intel Processor = "Intel"
	AMD   Processor = "AMD"
)

type Memory string

const (
	Small  Memory = "8GB"
	Medium Memory = "16GB"
	Large  Memory = "32GB"
)

type Video string

const (
	Nvidia Video = "Nvidia"
	Radeon Video = "Radeon"
)

type Computer struct {
	processor Processor
	memory    Memory
	video     Video
}

type ComputerBuilder struct {
	computer *Computer
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{computer: &Computer{}}
}

func (c *ComputerBuilder) SetProcessor(p Processor) *ComputerBuilder {
	c.computer.processor = p
	return c
}

func (c *ComputerBuilder) SetMemory(m Memory) *ComputerBuilder {
	c.computer.memory = m
	return c
}

func (c *ComputerBuilder) SetVideo(v Video) *ComputerBuilder {
	c.computer.video = v
	return c
}

func (c *ComputerBuilder) Build() *Computer {
	return c.computer
}
