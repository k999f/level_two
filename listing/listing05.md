Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:

Программа выведет:
```
error
```
Функция ```test()``` возвращает значение nil типа *customError, которое приводится к интерфейсу *error* (потому что у *customError* есть метод ```Error()```, следовательно, удовлетворяет интерфейсу).

В данном случае возвращаемый интерфейс будет иметь указатель на *itable* для типа *customError. Соответственно, его сравнение с *nil* будет *false*.