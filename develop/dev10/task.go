package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func sendToConn(c *net.Conn) {
	conn := *c
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		_, err := fmt.Fprintf(conn, "%s\n", input)

		if err != nil {
			log.Fatal("Writing to socket error: ", err)
		}
	}
}

func readFromConn(c *net.Conn) {
	conn := *c
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		output := scanner.Text()
		fmt.Println(output)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal("Reading from socket error: ", err)
	}
}

func main() {
	// Создаем и парсим флаг
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "Server connection timeout")
	flag.Parse()

	args := flag.Args()

	// Проверяем количество аргументов
	if flag.NArg() == 0 {
		log.Fatal("Need host and port")
	}

	host := args[0]
	port := args[1]

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	address := fmt.Sprintf("%s:%s", host, port)

	// Создаем подключение с таймаутом
	conn, err := net.DialTimeout("tcp", address, timeout)

	if err != nil {
		time.Sleep(timeout)
		log.Fatal("Connection error: ", err)
	}

	// Запускаем горутины для чтения и отправки
	go sendToConn(&conn)
	go readFromConn(&conn)

	<-done
	conn.Close()
	fmt.Println("Program interrupted")
}
