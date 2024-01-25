package pattern

import "fmt"

/*
	Паттерн цепочка вызовов позволяет выполнять операции над объектом при помощи цепочки обработчиков, запрос
	передается по этой цепочке, пока не будет обработан.

	Применяется: когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее
	неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.

	Плюсы: каждый обработчик выполняет свою логику независимо, поэтому уменьшается связь между клиентом и обработчиком;
	реализует первый и второй приницп SOLID (единственной обязанности и закрытости/открытости); можно с легкостью добавлять
	новые обработчики в цепочку.

	Минусы: для запроса может не найтись обработчика.

	Реальный пример использования: представим, что мы обрабатываем полученный запрос. Прежде чем выполнить его обработку
	необходимо выполнить определенные проверки корректности полученного запроса. Причем если одна из проверок не прошла, то
	дальнейшие проверки выполнять нет смысла. В данном случае на помощь может прийти паттерн цепочка вызовов.
*/

type Handler interface {
	HandleRequest(request string)
	SetNext(handler Handler)
}

type FirstHandler struct {
	NextHandler Handler
}

func (f *FirstHandler) SetNext(h Handler) {
	f.NextHandler = h
}

func (f *FirstHandler) HandleRequest(request string) {
	if request == "first" {
		fmt.Println("First handler processed request")
		return
	}

	if f.NextHandler != nil {
		f.NextHandler.HandleRequest(request)
	} else {
		fmt.Println("No handlers to process request")
	}
}

type SecondHandler struct {
	NextHandler Handler
}

func (s *SecondHandler) SetNext(h Handler) {
	s.NextHandler = h
}

func (s *SecondHandler) HandleRequest(request string) {
	if request == "second" {
		fmt.Println("Second handler processed request")
		return
	}

	if s.NextHandler != nil {
		s.NextHandler.HandleRequest(request)
	} else {
		fmt.Println("No handlers to process request")
	}
}

type ThirdHandler struct {
	NextHandler Handler
}

func (t *ThirdHandler) SetNext(h Handler) {
	t.NextHandler = h
}

func (t *ThirdHandler) HandleRequest(request string) {
	if request == "third" {
		fmt.Println("Third handler processed request")
		return
	}

	if t.NextHandler != nil {
		t.NextHandler.HandleRequest(request)
	} else {
		fmt.Println("No handlers to process request")
	}
}
