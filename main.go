package main

import (
	"fmt"
	"level_two/pattern"
)

func main() {
	// Фасад
	fmt.Println("Facade:")
	f := pattern.NewFacade()
	f.FacadePrint()

	fmt.Println()

	// Строитель
	fmt.Println("Builder:")
	b := pattern.NewCatBuilder()
	d := pattern.NewDirector(b)
	d.Build()
	c := b.GetAnimal()
	fmt.Printf("%+v\n", *c)

	fmt.Println()

	// Посетитель
	fmt.Println("Visitor:")
	с1 := pattern.Cat{Name: "First cat", Color: "orange"}
	с2 := pattern.Cat{Name: "Second cat", Color: "grey"}
	d1 := pattern.Dog{Name: "First dog", Color: "black"}
	d2 := pattern.Dog{Name: "Second dog", Color: "white"}

	data := &pattern.AnimalData{}
	с1.Accept(data)
	с2.Accept(data)
	d1.Accept(data)
	d2.Accept(data)

	fmt.Println()

	// Команда
	fmt.Println("Command:")
	c3 := &pattern.Cat1{}
	catControl := &pattern.CatControl{}
	sleepCommand := &pattern.SleepCommand{Cat: c3}
	wakeUpCommand := &pattern.WakeUpCommand{Cat: c3}

	catControl.SetCommand(sleepCommand)
	catControl.ExecuteCommand()

	catControl.SetCommand(wakeUpCommand)
	catControl.ExecuteCommand()

	fmt.Println()

	// Цепочка вызовов
	fmt.Println("Chain of responsibility:")
	firstHandler := &pattern.FirstHandler{}
	secondHandler := &pattern.SecondHandler{}
	thirdHandler := &pattern.ThirdHandler{}

	firstHandler.SetNext(secondHandler)
	secondHandler.SetNext(thirdHandler)

	firstHandler.HandleRequest("first")
	firstHandler.HandleRequest("second")
	firstHandler.HandleRequest("third")
	firstHandler.HandleRequest("invalid")

	fmt.Println()

	// Фабричный метод
	fmt.Println("Factory method:")
	c4, _ := pattern.CreateAnimal("cat")
	d4, _ := pattern.CreateAnimal("dog")

	fmt.Println(c4.GetName(), c4.GetColor())
	fmt.Println(d4.GetName(), d4.GetColor())

	fmt.Println()

	// Стратегия
	fmt.Println("Strategy:")
	context := &pattern.Context{}
	s1 := &pattern.FirstStrategy{StrategyName: "name1"}
	s2 := &pattern.SecondStrategy{StrategyName: "name2"}

	context.SetStrategy(s1)
	context.ExecuteStrategy()

	context.SetStrategy(s2)
	context.ExecuteStrategy()

	fmt.Println()

	// Состояние
	fmt.Println("State:")
	context1 := &pattern.Context1{}
	state1 := &pattern.FirstState{}
	state2 := &pattern.SecondState{}

	context1.SetState(state1)
	context1.FirstAction()
	context1.SecondAction()

	context1.SetState(state2)
	context1.FirstAction()
	context1.SecondAction()
}
