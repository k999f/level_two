package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func wget(siteURL string) (string, error) {
	response, err := http.Get(siteURL)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err := fmt.Errorf("Bad response status: %s", response.Status)
		return "", err
	}

	siteHTML, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(siteHTML), nil
}

func saveHTML(siteURL, siteHTML string) error {
	u, err := url.Parse(siteURL)

	if err != nil {
		return fmt.Errorf("Parsing URL for file name error: %s", err)
	}

	fileName := u.Host + u.Path
	fileName = strings.ReplaceAll(fileName, ".", "(d)")
	fileName = strings.ReplaceAll(fileName, "/", "(s)")
	fileName += ".html"
	file, err := os.Create(fileName)

	if err != nil {
		return fmt.Errorf("Creating file error: %s", err)
	}

	defer file.Close()

	file.Write([]byte(siteHTML))

	return nil
}

func main() {
	// Парсим название сайта
	args := os.Args

	// Проверяем количество аргументов
	if len(args) != 2 {
		log.Fatal("Wrong number of arguments")
	}

	// Получаем HTML
	siteURL := args[1]
	siteHTML, err := wget(siteURL)

	if err != nil {
		log.Fatal(err)
	}

	// Сохраняем HTML
	err = saveHTML(siteURL, siteHTML)

	if err != nil {
		log.Fatal(err)
	}
}
