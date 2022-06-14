package main

import (
	"fmt"
	"math"
)

const (
	applePrice = 5.99
	pearPrice  = 7.00
	purse      = 23.00
)

func main() {
	fmt.Printf("We need %v of money to buy 9 apples and 8 pears.\n", (applePrice*9)+(pearPrice)*8)
	fmt.Printf("We can buy %v pears.\n", math.Floor(purse/pearPrice))
	fmt.Printf("We can buy %v apples.\n", math.Floor(purse/applePrice))
	fmt.Printf("Can we buy 2 apples and 2 pears? - %v \n", (applePrice*2)+(pearPrice*2) <= purse)
}
