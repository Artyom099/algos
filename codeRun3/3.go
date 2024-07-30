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
	numbersStr, _ := reader.ReadString('\n')
	numbersStr = strings.TrimSpace(numbersStr)
	numbersStrs := strings.Split(numbersStr, " ")

	numbersArr := make([]int, 3)
	for i := 0; i < 3; i++ {
		numbersArr[i], _ = strconv.Atoi(numbersStrs[i])
	}

	nStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	nSlice := strings.Split(numbersStr, " ")
	fmt.Println(numbersArr)
	fmt.Println(nSlice)

	var buttonsForAdd int = 0

	for n := range nSlice {
		hasButton := slices.Contains(numbersArr, n)
		fmt.Println(hasButton)
		if !hasButton {
			buttonsForAdd++
		}
	}
	fmt.Println(buttonsForAdd)
}
