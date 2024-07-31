package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите кнопки: ")
	buttonsStr, _ := reader.ReadString('\n')
	buttonsStr = strings.TrimSpace(buttonsStr)
	buttonsStrs := strings.Split(buttonsStr, " ")
	fmt.Println("buttonsStrs", buttonsStrs)

	fmt.Print("Введите число: ")
	nStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	nSlice := strings.Split(nStr, "")
	fmt.Println("nSlice: ", nSlice, "len: ", len(nSlice))

	numbersArr := make([]int, 3)
	for i := 0; i < 3; i++ {
		numbersArr[i], _ = strconv.Atoi(buttonsStrs[i])
	}
	fmt.Println("numbersArr-кнопки: ", numbersArr)

	nArr := make([]int, len(nSlice))
	for i := 0; i < len(nSlice); i++ {
		nArr[i], _ = strconv.Atoi(nSlice[i])
	}
	fmt.Println("nArr-цифры числа", nArr)

	var buttonsForAdd int = 0

	for _, n := range nArr {
		hasButton := slices.Contains(numbersArr, n)
		fmt.Println(hasButton)
		if !hasButton {
			buttonsForAdd++
		}
	}
	fmt.Println("buttonsForAdd", buttonsForAdd)
}
