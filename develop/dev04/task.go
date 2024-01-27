package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func getAnagramSets(words *[]string) *map[string]*[]string {
	anagramsMap := make(map[string][]string)
	resultAnagramsMap := make(map[string]*[]string)

	// Создаем набор, где ключ - сортированное по буквам слово
	for _, word := range *words {
		word = strings.ToLower(word)

		splitWord := strings.Split(word, "")
		sort.Strings(splitWord)

		sortedWord := strings.Join(splitWord, "")

		// Проверяем наличие слова в наборе
		if !slices.Contains(anagramsMap[sortedWord], word) {
			anagramsMap[sortedWord] = append(anagramsMap[sortedWord], word)
		}
	}

	// Удаляем множества из одного элемента
	for k, v := range anagramsMap {
		if len(v) == 1 {
			delete(anagramsMap, k)
		}
	}

	/* Создаем набор, где ключ - первое встреченное слово в множестве, а значение - указатель на
	   сортированный слайс слов из множества */
	for _, v := range anagramsMap {
		anagramSlice := new([]string)

		newK := v[0]

		for _, word := range v {
			*anagramSlice = append(*anagramSlice, word)
		}

		sort.Strings(*anagramSlice)

		resultAnagramsMap[newK] = anagramSlice
	}

	return &resultAnagramsMap
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	anagramSets := getAnagramSets(&words)

	for k, v := range *anagramSets {
		fmt.Printf("Word: %s, anagrams: %v\n", k, *v)
	}
}
