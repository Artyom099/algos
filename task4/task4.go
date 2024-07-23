package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Чтение входных данных
	reader := bufio.NewReader(os.Stdin)

	// Считываем первую строку и преобразуем её в целое число - количество друзей
	mStr, _ := reader.ReadString('\n')
	mStr = strings.TrimSpace(mStr)
	M, _ := strconv.Atoi(mStr)

	// Создаем слайсы для хранения значений n и T
	nValues := make([]*big.Int, M+1)
	tValues := make([]int, M+1)

	// Считываем следующие M+1 строк
	for i := 0; i <= M; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			fmt.Println("Неверный формат строки:", line)
			return
		}

		// Читаем значение n и преобразуем его в big.Int - количество точек
		n := new(big.Int)
		n.SetString(parts[0], 10)

		// Читаем значение T и преобразуем его в int
		T, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Ошибка чтения числа T:", err)
			return
		}

		// Сохраняем значения в слайсы
		nValues[i] = n
		tValues[i] = T
	}

	// Пример использования считанных значений
	for i := 0; i <= M; i++ {
		fmt.Printf("Строка %d: n = %s, T = %d\n", i+1, nValues[i].String(), tValues[i])
	}
}
