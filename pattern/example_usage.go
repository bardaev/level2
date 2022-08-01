package pattern

import (
	"fmt"
	"log"
)

func Usage() {
	// 02_bulder
	var computerBuilder ComputerBuilder = *NewComputerBuilder()
	var computer Computer = *computerBuilder.SetProcessor(Intel).SetMemory(Large).SetVideo(Nvidia).Build()
	fmt.Println(computer)

	// 03_visitor
	backend := Developer{"Mark", "Zuckerberg", 1000, 32}
	boss := Director{"Bob", "Baggins", 2000, 40}

	backend.FullName()
	backend.Accept(CalculIncome{20})

	boss.FullName()
	boss.Accept(CalculIncome{10})

	// 04_command
	tv := &Tv{}
	onCommand := &OnCommand{
		device: tv,
	}
	offCommand := &OffCommand{
		device: tv,
	}
	onButton := &Button{
		command: onCommand,
	}
	onButton.press()
	offButton := &Button{
		command: offCommand,
	}
	offButton.press()

	// 05_chain
	var handler3 Handler = &ConcreteHandler3{Next: nil}
	var handler2 Handler = &ConcreteHandler2{Next: handler3}
	var handler1 Handler = &ConcreteHandler1{Next: handler2}

	handler1.SendRequest("Hello World!")

	// 06_factory_method
	mercedes, _ := GetCar("Mercedes")
	bmw, _ := GetCar("Bmw")
	fmt.Println(mercedes)
	fmt.Println(bmw)

	// 07_strategy
	add := Operation{Addition{}}
	add.Operator.Apply(5, 10)

	mult := Operation{Multiplication{}}
	mult.Operator.Apply(5, 10)

	// 08_state
	var thread Thread = *Newthread()

	err := thread.Create()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = thread.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = thread.Block()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = thread.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = thread.Terminate()
	if err != nil {
		log.Fatal(err.Error())
	}

}
