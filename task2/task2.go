package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Читаем входные данные
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разделяем строку на элементы
	elements := strings.Fields(input)

	// Определяем длину массива
	N := len(elements)

	// Печатаем сообщение
	fmt.Printf("Квадрат со стороной %d\n", N)

	// Формируем и печатаем квадрат
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(elements[j], " ")
		}
		fmt.Println()
	}
}
