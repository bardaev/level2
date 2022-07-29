package pattern

import "fmt"

func Usage() {
	// 02_bulder
	var computerBuilder ComputerBuilder = *NewComputerBuilder()
	var computer Computer = *computerBuilder.SetProcessor(Intel).SetMemory(Large).SetVideo(Nvidia).Build()
	fmt.Println(computer)

	// 05_chain
	var handler3 Handler = &ConcreteHandler3{Next: nil}
	var handler2 Handler = &ConcreteHandler2{Next: handler3}
	var handler1 Handler = &ConcreteHandler1{Next: handler2}

	handler1.SendRequest("Hello World!")
}
