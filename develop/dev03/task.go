package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Flags - структура с флагами
type Flags struct {
	K int
	N bool
	R bool
	U bool
	M bool
	B bool
	C bool
	H bool
}

var flags Flags

func createFlags() {
	flag.IntVar(&flags.K, "k", 0, "Specifying the column to sort")
	flag.BoolVar(&flags.N, "n", false, "Sort by numeric value")
	flag.BoolVar(&flags.R, "r", false, "Sort in reverse order")
	flag.BoolVar(&flags.U, "u", false, "Do not print duplicate lines")
	flag.BoolVar(&flags.M, "M", false, "Sort by month name")
	flag.BoolVar(&flags.B, "b", false, "Ignore trailing spaces")
	flag.BoolVar(&flags.C, "c", false, "Check if data is sorted")
	flag.BoolVar(&flags.H, "h", false, "Sort by numeric value taking into account suffixes")
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

func getPrettyString(lines []string) string {
	var result string
	for _, v := range lines {
		result += v + "\n"
	}
	result = result[:len(result)-1]
	return result
}

func getMainAndSuffix(s string) (string, string) {
	parts := strings.Split(s, ".")
	if len(parts) > 1 {
		return parts[0], parts[1]
	}
	return s, ""
}

func noFlags() bool {
	switch {
	case flags.K > 0:
		return false
	case flags.N:
		return false
	case flags.R:
		return false
	case flags.U:
		return false
	case flags.M:
		return false
	case flags.B:
		return false
	case flags.C:
		return false
	case flags.H:
		return false
	default:
		return true
	}
}

func sortNoFlags(lines []string) []string {
	less := func(i, j int) bool {
		return lines[i] < lines[j]
	}

	sort.Slice(lines, less)

	return lines
}

func sortFlagK(lines []string) []string {
	kFlagCorrect := true
	var less func(i, j int) bool

	// Проверяем, что во всех строках есть количество слов равное k
	for i := 0; i < len(lines); i++ {
		numFields := len(strings.Fields(lines[i]))
		if numFields < flags.K {
			kFlagCorrect = false
		}
	}

	// Если слов нужное количество - сортируем, иначе сортируем по первой колонке
	if kFlagCorrect {
		less = func(i, j int) bool {
			iFields := strings.Fields(lines[i])
			jFields := strings.Fields(lines[j])

			return iFields[flags.K-1] < jFields[flags.K-1]
		}
	} else {
		fmt.Println("Some lines has less words than k. Sorting with k = 1")
		less = func(i, j int) bool {
			return lines[i] < lines[j]
		}
	}

	sort.Slice(lines, less)

	return lines
}

func sortFlagN(lines []string) []string {
	less := func(i, j int) bool {
		iNum, iErr := strconv.Atoi(lines[i])
		jNum, jErr := strconv.Atoi(lines[j])

		if iErr == nil && jErr == nil {
			return iNum < jNum
		}

		return lines[i] < lines[j]
	}

	sort.Slice(lines, less)

	return lines
}

func sortFlagR(lines []string) []string {
	less := func(i, j int) bool {
		return lines[i] > lines[j]
	}

	sort.Slice(lines, less)

	return lines
}

func sortFlagU(lines []string) []string {
	linesMap := make(map[string]struct{})
	var uniqueLines []string

	for _, line := range lines {
		_, ok := linesMap[line]
		if !ok {
			linesMap[line] = struct{}{}
			uniqueLines = append(uniqueLines, line)
		}
	}

	return uniqueLines
}

func sortFlagM(lines []string) []string {
	less := func(i, j int) bool {
		iFields := strings.Fields(lines[i])
		jFields := strings.Fields(lines[j])

		iMonth, iErr := time.Parse("January", iFields[0])
		jMonth, jErr := time.Parse("January", jFields[0])

		if iErr == nil && jErr == nil {
			return iMonth.Before(jMonth)
		}

		return lines[i] < lines[j]
	}

	sort.Slice(lines, less)

	return lines
}

func sortFlagB(lines []string) []string {
	less := func(i, j int) bool {
		return strings.TrimSpace(lines[i]) < strings.TrimSpace(lines[j])
	}

	sort.Slice(lines, less)

	return lines
}

func sortFlagC(lines []string) bool {
	less := func(i, j int) bool {
		return lines[i] < lines[j]
	}

	return sort.SliceIsSorted(lines, less)
}

func sortFlagH(lines []string) []string {
	// Будем считать, что суффикс - все, что идет у числа после точки
	less := func(i, j int) bool {
		iMain, iSuffix := getMainAndSuffix(lines[i])
		jMain, jSuffix := getMainAndSuffix(lines[i])

		if iMain != jMain {
			return iMain < jMain
		}

		iSuffixNum, iErr := strconv.Atoi(iSuffix)
		jSuffixNum, jErr := strconv.Atoi(jSuffix)

		if iErr == nil && jErr == nil {
			return iSuffixNum < jSuffixNum
		}

		return lines[i] < lines[j]
	}

	sort.Slice(lines, less)

	return lines
}

func sortLines(lines []string) []string {
	if noFlags() {
		lines = sortNoFlags(lines)
	}

	if flags.K > 0 {
		lines = sortFlagK(lines)
	}

	if flags.N {
		lines = sortFlagN(lines)
	}

	if flags.R {
		lines = sortFlagR(lines)
	}

	if flags.U {
		lines = sortFlagU(lines)
	}

	if flags.M {
		lines = sortFlagM(lines)
	}

	if flags.B {
		lines = sortFlagB(lines)
	}

	if flags.C {
		sorted := sortFlagC(lines)
		if sorted {
			fmt.Println("Slice is sorted")
		} else {
			fmt.Println("Slice not sorted")
		}
	}

	if flags.H {
		lines = sortFlagH(lines)
	}

	return lines
}

func main() {
	// Создаем и парсим флаги
	createFlags()
	flag.Parse()

	// Проверяем, что указано имя файла
	if flag.NArg() == 0 {
		log.Fatal("Specify the name of the file to sort")
	}

	// Открываем файл
	file, err := os.Open(flag.Arg(0))
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

	lines = sortLines(lines)

	getPrettyString(lines)
	fmt.Println(getPrettyString(lines))
}
