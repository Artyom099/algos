package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// работает

func main() {
	// Чтение ввода
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	// Словарь для подсчета пар символов
	pairCounts := make(map[string]int)

	// Разбиение текста на слова
	words := strings.Fields(text)

	// Подсчет пар символов в каждом слове
	for _, word := range words {
		for i := 0; i < len(word)-1; i++ {
			// выбираем пару соседних символов
			pair := word[i : i+2]
			// увеличиваем счетчик этой пары на 1
			pairCounts[pair]++
		}
	}

	// Поиск самой часто встречающейся пары символов
	maxCount := 0
	var result string
	for pair, count := range pairCounts {
		if count > maxCount || (count == maxCount && pair > result) {
			maxCount = count
			result = pair
		}
	}

	// Вывод результата
	fmt.Println(result)
}
