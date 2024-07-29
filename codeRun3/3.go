package main

import (
	"bufio"
	"os"
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
	n, _ := strconv.Atoi(nStr)

}
