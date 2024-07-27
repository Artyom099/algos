package main

import (
	"fmt"
)

// решена +

func main() {
	var B, W int
	fmt.Scan(&B, &W)

	for n := 2; ; n++ {
		for m := 1; m <= n; m++ {
			if n >= m {
				// сумма черных плиток по периметру комнаты
				blackTiles := 2*n + 2*m - 4
				// количество внутренних белых плиток
				whiteTiles := (n - 2) * (m - 2)

				if whiteTiles < 0 {
					whiteTiles = 0
				}
				if blackTiles == B && whiteTiles == W {
					fmt.Printf("%d %d\n", n, m)
					return
				}
			}
		}
	}
}
