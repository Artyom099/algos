package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//

func main() {
	reader := bufio.NewReader(os.Stdin)
	numbersStr, _ := reader.ReadString('\n')
	numbersStr = strings.TrimSpace(numbersStr)
	numbersStrs := strings.Split(numbersStr, " ")

	numbersArr := make([]int, len(numbersStrs))
	for i := 0; i < len(numbersStrs); i++ {
		numbersArr[i], _ = strconv.Atoi(numbersStrs[i])
	}

	var maxInt, midInt, minInt int = -100_00_000, -100_00_000, -100_00_000

	for _, el := range numbersArr {
		if el > maxInt {
			minInt = midInt
			midInt = maxInt
			maxInt = el
		} else if el > midInt {
			minInt = midInt
			midInt = el
		} else if el > minInt {
			minInt = el
		}
	}
	fmt.Println(maxInt, midInt, minInt)

	//numCPU := runtime.NumCPU()
	//fmt.Printf("Number of logical CPUs available: %d\n", numCPU)
}
