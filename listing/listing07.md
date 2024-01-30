Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:

Программа выведет значения от *1* до *8*, а затем бесконечно будет выводить *0*. 

Такой вывод объясняется тем, что в функции ```merge()``` в бесконечном цикле читаются каналы *a* и *b* и полученные значения записываются в канал *c*. Когда каналы *a* и *b* закрываются, блок *select* в функции ```merge()``` все равно будет выполняться, поэтому будет происходить чтение из закрытых каналов типа *int* (это значит, что будем бесконечно получать нулевое значение типа *int*, которое равно *0*).
