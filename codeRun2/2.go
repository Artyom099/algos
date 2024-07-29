package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// работает +

func main() {
	reader := bufio.NewReader(os.Stdin)

	numbersStr, _ := reader.ReadString('\n')
	fmt.Println("numbersStr-1: ", numbersStr)

	// удаляем начальные и конечные пробелы
	numbersStr = strings.TrimSpace(numbersStr)
	fmt.Println("numbersStr-2: ", numbersStr)

	numbersStrs := strings.Split(numbersStr, " ")
	fmt.Println("numbersStrs: ", numbersStrs)

	// преобразуем слайс в массив чисел
	numbersArr := make([]int, len(numbersStrs))
	for i := 0; i < len(numbersStrs); i++ {
		numbersArr[i], _ = strconv.Atoi(numbersStrs[i])
	}
	fmt.Println("numbersArr: ", numbersArr)

	// идем со 2го элемента
	prev := numbersArr[0]
	for _, el := range numbersArr[1:] {
		if el <= prev {
			fmt.Println("NO")
			return
		}
		prev = el
	}
	fmt.Println("YES")

}
