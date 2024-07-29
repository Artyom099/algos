package main

//	func main() {
//		v := []int{} //{3, 4, 1, 2, 5}
//		v = append(v, 4, 1, 2, 5, 3)
//		v = ap(v)
//		sr(v)
//		fmt.Println(v)
//	}
//
//	func ap(arr []int) []int {
//		arr = append(arr, 10)
//		return arr // +
//	}
//
//	func sr(arr []int) {
//		sort.Ints(arr)
//	}

//func main() {
//	var foo []int //1 2 3 5
//	var bar []int //1 2 3 4
//
//	foo = append(foo, 1)
//	foo = append(foo, 2)
//	foo = append(foo, 3)
//
//	bar = append([]int{}, foo...)
//	bar = append(foo, 4)
//
//	foo = append(foo, 5)
//	foo = append(foo, 6)
//
//	fmt.Println(foo, bar)
//}

//func main() {
//	c := []string{"A", "B", "D", "E"}
//	// b := c[1:2] // c[1:2] = "B"
//
//	b := append([]string{}, c[1:2]...)
//	b = append(b, "TT")
//
//	fmt.Println(c)
//	fmt.Println(b)
//}

//func main() {
//	m := make(map[string]int)
//
//	// считаем количество повторений каждого слова, записали в словарь
//	for _, word := range []string{"hello", "world", "from", "the",
//		"best", "language", "in", "the", "world"} {
//		m[word]++
//	}
//	fmt.Println(m)
//
//	for k, v := range m {
//		fmt.Println(k, v)
//	}
//}

//func main() {
//	mutate := func(a *[]int) {
//		(*a)[0] = 0
//		*a = append(*a, 1) //append создает новый слайс, не изменяя старый
//		fmt.Println(*a)
//		//return a
//	}
//
//	a := []int{1, 2, 3, 4}
//	mutate(&a)
//	fmt.Println(a)
//}

//func a(s []int) {
//	s = append(s, 37)
//}
//func b(m map[int]int) {
//	m[3] = 33
//}
//func main() {
//	s := make([]int, 3, 8)
//	m := make(map[int]int, 8)
//	println(3)
//
//	// add to slice
//	a(s)
//	println(s[3]) // panic: runtime error: index out of range [3] with length 3
//
//	// add to map
//	b(m)
//	println(m[3]) // 37
//}

// не понимаю append
//func main() {
//	a := []int{1, 2}
//	a = append(a, 3)
//	b := append(a, 4)
//	fmt.Println(a)
//	c := append(a, 5)
//	fmt.Println(a)
//
//	fmt.Println(b)
//	fmt.Println(c)
//}

//func main() {
//	a := []int{1, 2}
//	a = append(a, 3)
//	a = append(a, 7)
//	b := append(a, 4)
//	c := append(a, 5)
//
//	fmt.Println(b)
//	fmt.Println(c)
//}

//func main() {
//	foo := make([]int, 0, 4)
//	foo = append(foo, 1, 2, 3)
//	bar := append(foo, 4)
//	baz := append(foo, 5)
//
//	fmt.Println(bar)
//	fmt.Println(baz)
//}
