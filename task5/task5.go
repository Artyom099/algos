package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Создаем новый ридер для чтения данных из стандартного ввода
	reader := bufio.NewReader(os.Stdin)

	numbersStr, _ := reader.ReadString('\n')
	numbersStr = strings.TrimSpace(numbersStr)
	numbersStrs := strings.Split(numbersStr, " ")

	// преобразуем слайс в массив чисел
	numbersArr := make([]int, 3)
	for i := 0; i < 3; i++ {
		numbersArr[i], _ = strconv.Atoi(numbersStrs[i])
	}

	// Сортировка вставками
	for i := 1; i < 3; i++ {
		key := numbersArr[i]
		j := i - 1
		for j >= 0 && numbersArr[j] > key {
			numbersArr[j+1] = numbersArr[j]
			j = j - 1
		}
		numbersArr[j+1] = key
	}

	// Выводим среднее число
	fmt.Println(numbersArr[1])
}
