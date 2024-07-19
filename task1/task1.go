package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	// Количество масок в коробке и упаковке
	masksPerBox := 24 * 16
	masksPerPack := 24

	// Количество коробок
	boxes := N / masksPerBox
	remainingAfterBoxes := N % masksPerBox

	// Количество упаковок
	packs := remainingAfterBoxes / masksPerPack

	// Количество отдельных масок
	masks := remainingAfterBoxes % masksPerPack

	// Проверки на то, можно ли уменьшить стоимость:

	// если стоимость необходимых масок 11, то выгоднее купить упаковку
	if float64(masks)*0.55 > 11 {
		packs++
		masks = 0
	}

	// если стоимость необходимых масок и упаковок больше 160, то выгоднее купить коробку
	if float64(masks)*0.55+float64(packs*11) > 160 {
		boxes++
		packs = 0
	}

	// Выводим результат
	fmt.Println(masks, packs, boxes)
}
