package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Функция, возвращающая двумерный срез найденных индексов вхождения символов
// из chars в последних словах предложений.
// Сравнение символов для поиска происходит без учета регистра.
func symbolFinder(sentences []string, chars []rune) [][]int {
	var result [][]int

	for _, sentence := range sentences {
		words := strings.Fields(sentence)
		lastWordLower := strings.ToLower(words[len(words)-1])

		var tmpValues []int
		for _, char := range chars {
			charLower := unicode.ToLower(char)
			index := strings.IndexRune(lastWordLower, charLower)
			tmpValues = append(tmpValues, index)
		}
		result = append(result, tmpValues)
	}
	return result
}

func main() {
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'H', 'E', 'L', 'П', 'М'}

	result := symbolFinder(sentences, chars)

	// Вывод результата найденных индексов в удобной форме.
	for i, sentence := range sentences {
		fmt.Printf("%d \"%s\" - %v\n", i+1, sentence, result[i])
	}
}
