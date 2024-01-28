package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита grepLines ===

Реализовать утилиту фильтрации (man grepLines)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Flags - структура с флагами
type Flags struct {
	A       int
	B       int
	Context int
	Count   bool
	I       bool
	V       bool
	F       bool
	N       bool
}

var flags Flags

func createFlags() {
	flag.IntVar(&flags.A, "A", 0, "Print +N lines after match")
	flag.IntVar(&flags.B, "B", 0, "Print +N lines before match")
	flag.IntVar(&flags.Context, "C", 0, "Print ±N lines around match")
	flag.BoolVar(&flags.Count, "c", false, "Number of lines")
	flag.BoolVar(&flags.I, "i", false, "Ignore case")
	flag.BoolVar(&flags.V, "v", false, "Instead of matching, exclude")
	flag.BoolVar(&flags.F, "F", false, "Exact match to string, not a pattern")
	flag.BoolVar(&flags.N, "n", false, "Print line numbers")
}

func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, scanner.Err()
}

func getPrettyString(lines []string, linesIdx []int) string {
	var result string
	for i, v := range lines {
		if len(linesIdx) != 0 {
			result += strconv.Itoa(linesIdx[i]+1) + ". " + v + "\n"
		} else {
			result += v + "\n"
		}
	}
	if len(result) != 0 {
		result = result[:len(result)-1]
	}
	return result
}

func grepLines(lines []string, pattern string) ([]string, []int) {
	var re *regexp.Regexp
	var matchedLines []string
	var matchedLinesIdx []int

	oldPattern := pattern

	// Игнорируем регистр, если установлен флаг
	if flags.I {
		pattern = "(?i)" + pattern
	}

	re, err := regexp.Compile(pattern)

	if err != nil {
		log.Fatal("Error compiling regexp: ", err)
	}

	for i, line := range lines {
		matched := false

		// Проверяем точное совпадение со строкой (если установлен флаг) или с регулярным выражением
		if flags.F {
			if flags.I {
				if strings.ToLower(line) == strings.ToLower(oldPattern) {
					matched = true
				}
			} else {
				if line == pattern {
					matched = true
				}
			}
		} else {
			if re.Match([]byte(line)) {
				matched = true
			}
		}

		// Инвертируем результат совпадения, если установлен флаг
		if flags.V {
			matched = !matched
		}

		if matched {
			matchedLines = append(matchedLines, line)
			matchedLinesIdx = append(matchedLinesIdx, i)

			before := flags.B
			after := flags.A

			// Если указаны A, B, C, то выбираем наибольшее
			if flags.Context > before {
				before = flags.Context

			}
			if flags.Context > after {
				after = flags.Context
			}

			// Добавляем N строк до совпадения при возможности
			if before > 0 && i-before >= 0 {
				matchedLines = append(matchedLines, lines[i-before:i]...)
				var beforeLinesNumbers []int
				for j := i - before; j < i; j++ {
					beforeLinesNumbers = append(beforeLinesNumbers, j)
				}
				matchedLinesIdx = append(matchedLinesIdx, beforeLinesNumbers...)
			}

			// Добавляем N строк после совпадения при возможности
			if after > 0 && i+after < len(lines) {
				matchedLines = append(matchedLines, lines[i+1:i+1+after]...)
				var afterLinesNumbers []int
				for j := i + 1; j < i+1+after; j++ {
					afterLinesNumbers = append(afterLinesNumbers, j)
				}
				matchedLinesIdx = append(matchedLinesIdx, afterLinesNumbers...)
			}
		}
	}

	return matchedLines, matchedLinesIdx
}

func main() {
	// Создаем и парсим флаги
	createFlags()
	flag.Parse()

	// Проверяем, что указан паттерн и имя файла
	if flag.NArg() < 2 {
		log.Fatal("Specify pattern and name of the file")
	}

	// Парсим паттерн
	pattern := flag.Arg(0)

	// Открываем файл
	file, err := os.Open(flag.Arg(1))
	if err != nil {
		log.Fatal("File opening error: ", err)
	}

	// Читаем файл
	lines, err := readLines(file)
	if err != nil {
		log.Fatal("File reading error: ", err)
	}

	// Закрываем файл
	err = file.Close()
	if err != nil {
		log.Fatal("File closing error: ", err)
	}

	// Получаем подходящие строки
	matchedLines, matchedLinesIdx := grepLines(lines, pattern)

	// Выводим количество совпадений, если установлен флаг
	if flags.Count {
		fmt.Println("Matches: ", len(matchedLines))
	}

	// Добавляем номера строк, если установлен флаг
	if flags.N {
		fmt.Println(getPrettyString(matchedLines, matchedLinesIdx))
	} else {
		fmt.Println(getPrettyString(matchedLines, nil))
	}
}
