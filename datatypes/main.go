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
	applesToBuy := 9
	pearsToBuy := 8
	fmt.Printf("We need %v of money to buy 9 apples and 8 pears.\n", (applePrice*float64(applesToBuy))+(pearPrice)*float64(pearsToBuy))
	fmt.Printf("We can buy %v pears.\n", math.Floor(purse/pearPrice))
	fmt.Printf("We can buy %v apples.\n", math.Floor(purse/applePrice))
	applesToBuy = 2
	pearsToBuy = 2
	fmt.Printf("Can we buy 2 apples and 2 pears? - %v \n", (applePrice*float64(applesToBuy))+(pearPrice*float64(pearsToBuy)) <= purse)
}
