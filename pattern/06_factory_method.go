package pattern

import "errors"

/*
	Паттерн фабричный метод позволяет создать общий интерфейс для создания объектов в классе, позволяя
	подклассам изменять тип создаваемых объектов.

	Применяется: когда заранее неизвестны типы и зависимости объектов, с которыми должен работать код.

	Плюсы: универсальный код для создания объектов различных типов; нет привязки к конкретным классам продуктов;
	позволяет легко добавлять новые продукты в код; реализует второй принцип SOLID (открытости/закрытости).

	Минусы: для каждого класса продукта надо создать свой подкласс создателя.

	Реальный пример использования: представим, что в программе необходимо создавать объекты различных типов, причем
	в дальнейшем будут добавлено новые типы, объекты которых также необходимо будет создавать. Чтобы не усложнять
	модификацию кода при появлении новых типов, можно использовать фабричный метод, который будет отвечать за содание
	новых объектов.
*/

type AnimalInterface2 interface {
	SetName(name string)
	SetColor(color string)
	GetName() string
	GetColor() string
}

type Animal2 struct {
	name  string
	color string
}

func (a *Animal2) SetName(name string) {
	a.name = name
}

func (a *Animal2) SetColor(color string) {
	a.color = color
}

func (a *Animal2) GetName() string {
	return a.name
}

func (a *Animal2) GetColor() string {
	return a.color
}

type Cat2 struct {
	Animal2
}

type Dog2 struct {
	Animal2
}

func newCat() AnimalInterface2 {
	return &Cat2{
		Animal2: Animal2{
			name:  "Cat's name",
			color: "Cat's color",
		},
	}
}

func newDog() AnimalInterface2 {
	return &Dog2{
		Animal2: Animal2{
			name:  "Dog's name",
			color: "Dog's color",
		},
	}
}

func CreateAnimal(animalType string) (AnimalInterface2, error) {
	switch animalType {
	case "cat":
		return newCat(), nil
	case "dog":
		return newDog(), nil
	default:
		return nil, errors.New("Wrong animal type")
	}
}
