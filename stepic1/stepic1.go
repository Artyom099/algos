package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ф-я вывода уникальных значений отсортированного массива
func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string

	for in.Scan() {
		txt := in.Text()

		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file not sorted")
		}
		prev = txt

		fmt.Fprintln(output, txt)
	}
	return nil
}

func main() {
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}

// ф-я вывода уникальных значений
func main2() {
	in := bufio.NewScanner(os.Stdin)

	alreadySeen := make(map[string]bool)

	for in.Scan() {
		txt := in.Text()

		// если эту строку уже видели, пропускаем ее
		if _, found := alreadySeen[txt]; found {
			continue
		}
		alreadySeen[txt] = true

		fmt.Println(txt)
	}
}
