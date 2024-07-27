package main

import (
	"bufio"
	"fmt"
	"os"
)

// решена +

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	// Чтение игрового поля
	field := make([]string, n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			field[i] = scanner.Text()
		}
	}

	// Функция для проверки наличия 5 подряд одинаковых символов
	checkWin := func(symbol byte) bool {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if field[i][j] != symbol {
					continue
				}
				// Проверка горизонтальных линий
				if j <= m-5 {
					if field[i][j+1] == symbol &&
						field[i][j+2] == symbol &&
						field[i][j+3] == symbol &&
						field[i][j+4] == symbol {
						return true
					}
				}
				// Проверка вертикальных линий
				if i <= n-5 {
					if field[i+1][j] == symbol &&
						field[i+2][j] == symbol &&
						field[i+3][j] == symbol &&
						field[i+4][j] == symbol {
						return true
					}
				}
				// Проверка диагоналей слева направо
				if i <= n-5 && j <= m-5 {
					if field[i+1][j+1] == symbol &&
						field[i+2][j+2] == symbol &&
						field[i+3][j+3] == symbol &&
						field[i+4][j+4] == symbol {
						return true
					}
				}
				// Проверка диагоналей справа налево
				if i <= n-5 && j >= 4 {
					if field[i+1][j-1] == symbol &&
						field[i+2][j-2] == symbol &&
						field[i+3][j-3] == symbol &&
						field[i+4][j-4] == symbol {
						return true
					}
				}
			}
		}
		return false
	}

	// Проверка выигрыша для символов 'X' и 'O'
	if checkWin('X') || checkWin('O') {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
