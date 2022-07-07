package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cat struct {
	weight          int
	feedConsumption int
}

func (c cat) feedPerMonth() int {
	return c.weight * c.feedConsumption
}

func (c cat) String() string {
	return "Cat"
}

func (c cat) returnAnimalWeight() int {
	return c.weight
}

type dog struct {
	weight          int
	feedConsumption int
}

func (d dog) feedPerMonth() int {
	return d.weight * d.feedConsumption
}

func (d dog) String() string {
	return "Dog"
}

func (d dog) returnAnimalWeight() int {
	return d.weight
}

type cow struct {
	weight          int
	feedConsumption int
}

func (c cow) feedPerMonth() int {
	return c.weight * c.feedConsumption
}

func (c cow) String() string {
	return "Cow"
}

func (c cow) returnAnimalWeight() int {
	return c.weight
}

type feedAmountGetter interface {
	feedPerMonth() int
}

type animalNameGetter interface {
	String() string
}

type animalWeightGetter interface {
	returnAnimalWeight() int
}

type animalInfoGetter interface {
	feedAmountGetter
	animalNameGetter
	animalWeightGetter
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var animals []animalInfoGetter

	for i := 0; i < rand.Intn(5)+1; i++ {
		animals = append(animals, cat{weight: int(rand.Intn(7) + 3), feedConsumption: 7})
	}
	for i := 0; i < rand.Intn(5)+1; i++ {
		animals = append(animals, dog{weight: int(rand.Intn(62) + 3), feedConsumption: 2})
	}
	for i := 0; i < rand.Intn(5)+1; i++ {
		animals = append(animals, cow{weight: int(rand.Intn(400) + 700), feedConsumption: 25})
	}

	totalFeed := calculateFoodConsumption(animals)
	fmt.Printf("Total amount of feed needed: %d", totalFeed)
}

func calculateFoodConsumption(animals []animalInfoGetter) int {
	var totalFeed int
	feedPerAnimal := make(map[string]int)

	for _, v := range animals {
		requiredFeed := v.feedPerMonth()
		totalFeed += requiredFeed
		feedPerAnimal[v.String()] += requiredFeed
		fmt.Printf("This is %s, it weights %d, and it needs %d of feed per month.\n", v.String(), v.returnAnimalWeight(), requiredFeed)
	}

	for k, v := range feedPerAnimal {
		fmt.Printf("%ss need a total of %d of feed.\n", k, v)
	}
	return totalFeed
}
