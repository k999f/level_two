package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	// Единый done-канал
	singleCh := make(chan interface{})

	// Чтобы не закрыть канал повторно
	var o sync.Once

	// Закрываем канал, если в функцию ничего не было передано
	if len(channels) == 0 {
		close(singleCh)
		return singleCh
	}

	for _, ch := range channels {
		// На каждый канал запускаем горутину
		go func(ch <-chan interface{}) {
			select {
			// Если прочитали из канала, то закрываем основной
			case <-ch:
				o.Do(func() {
					close(singleCh)
				})
			// Если закрыли основной канал, то завершаем селект
			case <-singleCh:
				break
			}
		}(ch)
	}

	return singleCh
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("Done after %v", time.Since(start))
}
