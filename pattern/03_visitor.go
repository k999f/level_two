package pattern

import "fmt"

/*
	Паттерн посетитель используется для добавления новых операций для элементов некоторой структуры объектов.

	Применяется: когда нужно применить операцию над всеми элементами сложной структуры, не изменяя классы этих элементов.

	Плюсы: позволяет (практически) не изменять существующие классы; объединяет родственные операции в одном классе;
	при помощи посетителя можно накапливать состояние при обходе элементов.

	Минусы: может привести к нарушению инкапсуляции (необходимо будет предоставлять доступ к полям и методам);
	если иерархия классов часто меняется, то паттерн не оправдан.

	Реальный пример использования: представим, что имеется некоторая сложная структура (например, дерево), которая
	состоит из элементов различных классов. И нам необходимо применять одну и ту же операцию ко всем элементам структуры
	(объектам различных классов).
*/

type AnimalInterface interface {
	GetAnimalColor() string
	Accept(Visitor)
}

type Cat struct {
	Name  string
	Color string
}

func (c *Cat) GetAnimalColor() string {
	return c.Color
}

func (c *Cat) Accept(v Visitor) {
	v.VisitForCat(c)
}

type Dog struct {
	Name  string
	Color string
}

func (d *Dog) GetAnimalColor() string {
	return d.Color
}

func (d *Dog) Accept(v Visitor) {
	v.VisitForDog(d)
}

type Visitor interface {
	VisitForCat(*Cat)
	VisitForDog(*Dog)
}

type AnimalData struct {
	Data string
}

func (a *AnimalData) VisitForCat(c *Cat) {
	a.Data = fmt.Sprintf("I'm %s cat and my name is %s", c.GetAnimalColor(), c.Name)
	fmt.Println(a.Data)
}

func (a *AnimalData) VisitForDog(d *Dog) {
	a.Data = fmt.Sprintf("I'm %s dog and my name is %s", d.GetAnimalColor(), d.Name)
	fmt.Println(a.Data)
}
