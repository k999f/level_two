package pattern

import "fmt"

/*
	Паттерн стратегия позволяет поместить в собственный класс алгоритмы, после чего алгоритмы можно
	взаимозаменять прямо во время исполнения программы.

	Применяется: когда необходимо использовать разные вариации алгоритма внутри одного объекта.

	Плюсы: позволяет заменять алгоритмы во время исполнения программы; изолирует код алгоритмов от других классов;
	реализует второй принцип SOLID (открытости/закрытости).

	Минусы: усложнение кода из-за появления новых классов; пользователю необходимо знать разницу между стратегиями.

	Реальный пример использования: представим, что в приложении у одной и той же проблемы может быть несколько решений.
	В данном случае паттерн стратегия позволяет объединять похожие алгоритмы в единое семейство и подключать тот или иной
	алгоритм в работу по необходимости.
*/

type Strategy interface {
	Execute()
}

type FirstStrategy struct {
	StrategyName string
}

func (s *FirstStrategy) Execute() {
	fmt.Printf("Executing %s strategy of first type\n", s.StrategyName)
}

type SecondStrategy struct {
	StrategyName string
}

func (s *SecondStrategy) Execute() {
	fmt.Printf("Executing %s strategy of second type\n", s.StrategyName)
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}
