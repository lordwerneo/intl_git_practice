package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string
	var min, max int32

	dividedInput := strings.Split(input, " ")

	for i, v := range dividedInput {
		if i == 0 {
			intOfValue, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			min, max = int32(intOfValue), int32(intOfValue)
			continue
		}

		intOfValue, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		if int32(intOfValue) < min {
			min = int32(intOfValue)
		}
		if int32(intOfValue) > max {
			max = int32(intOfValue)
		}
	}

	result = fmt.Sprintf("%d %d", max, min)
	fmt.Println(result)
}
