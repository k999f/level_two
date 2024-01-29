package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func shell(args []string) error {
	argsErr := errors.New("Wrong number of arguments")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return argsErr
		}

		err := os.Chdir(args[1])

		if err != nil {
			return err
		}

	case "pwd":
		if len(args) != 1 {
			return argsErr
		}

		workDirPath, err := os.Getwd()

		if err != nil {
			return err
		}

		fmt.Println(workDirPath)
	case "echo":
		if len(args) < 2 {
			return argsErr
		}

		fmt.Println(strings.Join(args[1:], " "))
	case "kill":
		if len(args) < 2 {
			return argsErr
		}

		for _, pid := range args[1:] {
			err := exec.Command("kill", pid).Run()

			if err != nil {
				return err
			}
		}
	case "ps":
		if len(args) != 1 {
			return argsErr
		}

		output, err := exec.Command("ps").Output()

		if err != nil {
			return err
		}

		fmt.Println(string(output))

	default:
		command := exec.Command(args[0], args[1:]...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			return
		}

		// Парсим команду
		input := scanner.Text()
		args := strings.Fields(input)

		// Проверяем ввод на пустоту
		if len(args) == 0 {
			fmt.Println("Error. No command")
			continue
		}

		// Выполняем команду
		err := shell(args)
		if err != nil {
			fmt.Println(err)
		}
	}
}
