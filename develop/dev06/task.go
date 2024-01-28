package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Flags - структура с флагами
type Flags struct {
	F int
	D string
	S bool
}

var flags Flags

func getPrettyString(lines []string) string {
	var result string
	for _, v := range lines {
		result += v + "\n"
	}
	if len(result) != 0 {
		result = result[:len(result)-1]
	}
	return result
}

func cutLines(lines []string) []string {
	var result []string
	// Проходимся по всем строкам
	for _, line := range lines {
		// Создаем мапу, где ключ - номер колонки, а значение - слово
		mapCols := make(map[int]string)
		lineCols := strings.Split(line, flags.D)

		// Заполняем мапу
		for i, v := range lineCols {
			mapCols[i+1] = v
		}

		// Если в строке больше одной колонки и установлен флаг s (то есть в строке были разделители)
		// или если флага s нет
		if (flags.S && len(mapCols) > 1) || !flags.S {
			v, ok := mapCols[flags.F]
			if ok {
				result = append(result, v)
			}
		}
	}
	return result
}

func main() {
	// Создаем и парсим флаги
	flag.IntVar(&flags.F, "f", -1, "Select fields (columns)")
	flag.StringVar(&flags.D, "d", " ", "Use a different delimiter")
	flag.BoolVar(&flags.S, "s", false, "Only delimited lines")
	flag.Parse()

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	// Читаем строки, пока не введем пустую строку или CTRL+D
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			lines = append(lines, text)
		} else {
			fmt.Println("Input stopped")
			break
		}
	}

	result := cutLines(lines)
	fmt.Println(getPrettyString(result))
}
