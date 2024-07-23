package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// todo - вообще не понял эту задачу
func main() {
	// Чтение входных данных
	reader := bufio.NewReader(os.Stdin)

	// количество дней
	nStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	n, _ := strconv.Atoi(nStr)

	// цены за каждый день
	pricesStr, _ := reader.ReadString('\n')
	pricesStr = strings.TrimSpace(pricesStr)
	pricesStrs := strings.Split(pricesStr, " ")

	prices := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i], _ = strconv.Atoi(pricesStrs[i])
	}

	// Инициализация переменных
	maxProfit := 0
	currentArea := 1

	// Проходим массив цен с конца к началу
	maxPrice := 0
	for i := n - 1; i >= 0; i-- {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		}
		maxProfit += maxPrice * currentArea
		currentArea++
	}

	fmt.Println(maxProfit)
}
