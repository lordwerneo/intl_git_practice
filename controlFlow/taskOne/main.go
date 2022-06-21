package main

import "fmt"

func main() {
	arr := []int{3, 4, 4, 3, 6, 3}
	uniqueElements := make(map[int]bool)
	result := make([]int, 0)

	for _, v := range arr {
		if _, ok := uniqueElements[v]; ok {
			continue
		}
		uniqueElements[v] = true
		result = append(result, v)
	}

	fmt.Println(result)

}
