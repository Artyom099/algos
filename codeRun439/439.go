package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// проходит 19 из 25 тестов
func main() {
	// Чтение входных данных
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	// Словарь для подсчета пар символов
	substrCounts := make(map[string]int)

	// Разбиение текста на слова и подсчет подстрок
	words := strings.Fields(text)
	for _, word := range words {
		wordLen := len(word)
		for i := 0; i < wordLen; i++ {
			for j := i + 1; j <= wordLen; j++ {
				substr := word[i:j]
				substrCounts[substr]++
			}
		}
	}
	fmt.Println("words", words)
	fmt.Println("substrCounts", substrCounts)

	// Поиск самой частой подстроки
	maxCount := 0
	var result string
	for substr, count := range substrCounts {
		if (count > maxCount || (count == maxCount && substr > result)) && len(substr) > 1 {
			maxCount = count
			result = substr
		}
	}

	// Вывод результата
	fmt.Println(result)
}
