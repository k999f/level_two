package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// UnpackString выполняет распаковку строки
func UnpackString(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	var result string
	sRune := []rune(s)
	var prevRune rune
	errorStirng := errors.New("Incorrect string")

	// Проверяем является ли первый символ цифрой
	if unicode.IsDigit(sRune[0]) {
		return "", errorStirng
	}

	// Проходимся в цикле по всем символам
	for _, currentRune := range sRune {
		// Если текущий символ - число
		if unicode.IsDigit(currentRune) {
			// Проверяем, что не идут 2 числа подряд
			if unicode.IsDigit(prevRune) {
				return "", errorStirng
			}

			// Приводим rune к int
			repeatNum, err := strconv.Atoi(string(currentRune))
			if err != nil {
				return "", err
			}

			// Повторяем предыдущий символ нужное количество раз
			result += strings.Repeat(string(prevRune), repeatNum)

			prevRune = currentRune

			continue
		}

		/* Если текущий символ не число и предыдущий тоже не число, а также prevRune не пуста, то добавляем
		предыдущий символ в результат */
		if !unicode.IsDigit(prevRune) && prevRune != 0 {
			result += string(prevRune)
		}

		prevRune = currentRune
	}

	// Если в конце строки стояла буква, добавляем ее в результат
	if !unicode.IsDigit(prevRune) {
		result += string(prevRune)
	}

	return result, nil
}

func main() {
	res, err := UnpackString("a4bc2d5e")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(res)
}
