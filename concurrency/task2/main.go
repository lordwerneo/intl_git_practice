package main

import "fmt"

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”
func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	// Ваша реалізація
	var sum int
	var ch = make(chan int)
	go func(ch chan int, intSlice [][]int) {
		for _, slice := range intSlice {
			result := 0
			for _, v := range slice {
				result += v
			}
			ch <- result
		}
		close(ch)
	}(ch, n)
	for value := range ch {
		sum += value
	}
	fmt.Printf("result: %d", sum)
}
