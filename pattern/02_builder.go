package pattern

/*
	Паттерн строитель используется для пошагового создания сложного объекта.

	Применяется: когда необходимо поэтапно создавать сложный объект;
	когда необходимо создавать разные представления объекта (например, животного - кошка, собака). Также
	можно использовать дополнительный объект - директор, который будет выполнять выозов методов строителя.

	Плюсы: позволяет пошагово создавать объект; позволяет изолировать логику создания объекта;
	позволяет использовать один и тот же код для различных представлений объекта.

	Минусы: появляются дополнительные классы, поэтому код программы усложняется; при внесении изменения
	в класс конструемого объекта также скорее всего придется вносить изменения в класс конкретного строителя.

	Реальный пример использования: представим, что мы создаем сложные объекты, построение которых выполняется за
	несколько операций (например, животное), а также существует нескольких видов объекта (например, кошка и собака).
*/

type Animal struct {
	Head string
	Body string
	Paws string
	Tail string
}

type AnimalBuilder interface {
	BuildHead()
	BuildBody()
	BuildPaws()
	BuildTail()
	GetAnimal() *Animal
}

type CatBuilder struct {
	Cat *Animal
}

func NewCatBuilder() *CatBuilder {
	return &CatBuilder{Cat: &Animal{}}
}

func (c *CatBuilder) BuildHead() {
	c.Cat.Head = "Cat's head"
}

func (c *CatBuilder) BuildBody() {
	c.Cat.Body = "Cat's body"
}

func (c *CatBuilder) BuildPaws() {
	c.Cat.Paws = "Cat's paws"
}

func (c *CatBuilder) BuildTail() {
	c.Cat.Tail = "Cat's tail"
}

func (c *CatBuilder) GetAnimal() *Animal {
	return c.Cat
}

type Director struct {
	b AnimalBuilder
}

func NewDirector(builder AnimalBuilder) *Director {
	return &Director{b: builder}
}

func (d *Director) Build() {
	d.b.BuildHead()
	d.b.BuildBody()
	d.b.BuildPaws()
	d.b.BuildTail()
}
