package pattern

import "fmt"

/*
	Паттерн состояние позволяет во время выполнения программы менять поведение объекта в зависимости от его состояния.

	Применяется: когда необходимо менять поведение объекта в зависимости от его текущего состояния.

	Плюсы: избавляет от необходимости использовать множество условных операторов; концентрирует в одном месте код, котоырй
	связан с состояниями.

	Минусы: усложнение кода из-за появления новых классов; приведет к усложнению кода, если состояний мало и они редко
	меняются.

	Реальный пример использования: представим, что программа может находиться в одном из нескольких состояний, которые всё
	время сменяют друг друга. Набор этих состояний, а также переходов между ними, предопределён и конечен. Находясь в разных
	состояниях, программа может по-разному реагировать на одни и те же события, которые происходят с ней.
*/

type State interface {
	firstAction()
	secondAction()
}

type Context1 struct {
	currentState State
	firstState   State
	secondState  State
}

func (c *Context1) SetState(state State) {
	c.currentState = state
}

func (c *Context1) FirstAction() {
	c.currentState.firstAction()
}

func (c *Context1) SecondAction() {
	c.currentState.secondAction()
}

type FirstState struct {
	context *Context
}

func (s *FirstState) firstAction() {
	fmt.Println("First state, first action")
}

func (s *FirstState) secondAction() {
	fmt.Println("First state, second action")
}

type SecondState struct {
	context *Context
}

func (s *SecondState) firstAction() {
	fmt.Println("Second state, first action")
}

func (s *SecondState) secondAction() {
	fmt.Println("Second state, second action")
}
