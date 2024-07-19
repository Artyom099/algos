package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	//todo - почему-то считываем только символы до пробела, после пробела не видит их
	fmt.Scan(&str)
	//fmt.Fscan(os.Stdin, &str)

	slice := strings.SplitAfter(str, " ")
	fmt.Println(str, slice, len(slice), cap(slice))
	fmt.Printf("Квадрат со стороной %d", cap(slice))
}
