package main

import (
	"log"

	"github.com/reiver/go-telnet"
)

func runServer() {
	// Отвечаем на сообщение этим же сообщением
	var handler telnet.Handler = telnet.EchoHandler

	err := telnet.ListenAndServe(":5555", handler)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

func main() {
	runServer()
}
