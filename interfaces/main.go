package main

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type feedAmountGetter interface {
	feedPerMonth() int
}

type animalNameGetter interface {
	String() string
}

type animalWeightGetter interface {
	returnAnimalWeight() int
}

type animalMinimumWeightGetter interface {
	returnMinimalWeight() int
}

type edibleStatusGetter interface {
	returnEdibleStatus() bool
}

type impersonateStatusGetter interface {
	returnImpersonateStatus() string
}

type animalInfoGetter interface {
	feedAmountGetter
	animalNameGetter
	animalWeightGetter
	animalMinimumWeightGetter
	edibleStatusGetter
	impersonateStatusGetter
}

var wrongType = errors.New("wrong type")
var lowWeight = errors.New("weight is too low")
var isNotEdible = errors.New("is not edible")

func main() {
	rand.Seed(time.Now().UnixNano())
	var animals []animalInfoGetter

	for i := 0; i < rand.Intn(2)+1; i++ {
		animals = append(animals, cat{name: fmt.Sprintf("IAmCat%d", i), weight: int(rand.Intn(8) + 2),
			feedConsumption: 7, isEdible: false, minimalWeight: 3, impersonate: generateRandomResponse()})
	}
	for i := 0; i < rand.Intn(2)+1; i++ {
		animals = append(animals, dog{name: fmt.Sprintf("IAmDog%d", i), weight: int(rand.Intn(63) + 2),
			feedConsumption: 2, isEdible: false, minimalWeight: 3, impersonate: generateRandomResponse()})
	}
	for i := 0; i < rand.Intn(2)+1; i++ {
		animals = append(animals, cow{name: fmt.Sprintf("IAmCow%d", i), weight: int(rand.Intn(500) + 600),
			feedConsumption: 25, isEdible: true, minimalWeight: 600, impersonate: generateRandomResponse()})
	}

	totalFeed, err := calculateFoodConsumption(animals)
	if err != nil {
		err = fmt.Errorf("stop working, critical error: %w", err)
		fmt.Println(err)
	} else {
		fmt.Printf("Total amount of feed needed: %d", totalFeed)
	}
}

func calculateFoodConsumption(animals []animalInfoGetter) (int, error) {
	var totalFeed int
	for _, v := range animals {
		err := validateAnimals(v)
		if err != nil {
			if errors.Is(err, lowWeight) {
				err = fmt.Errorf("validation failed: %w", err)
				err = fmt.Errorf("for %s: %w", v, err)
				return 0, err
			} else if errors.Is(err, wrongType) {
				err = fmt.Errorf("validation failed: %w", err)
				err = fmt.Errorf("for %s: %w", v, err)
				fmt.Println(err)
				continue
			} else if errors.Is(err, isNotEdible) {
				err = fmt.Errorf("validation failed: %w", err)
				err = fmt.Errorf("for %s: %w", v, err)
				fmt.Println(err)
				continue
			}
		}
		requiredFeed := v.feedPerMonth()
		totalFeed += requiredFeed
		fmt.Printf("This is %s, it weights %d, and it needs %d of feed per month.\n", v.String(), v.returnAnimalWeight(), requiredFeed)
	}
	return totalFeed, nil
}

func validateAnimals(animal animalInfoGetter) error {
	err := checkWeight(animal)
	if err != nil {
		err = fmt.Errorf("minimal weight is %d, but it weights %d: %w", animal.returnMinimalWeight(),
			animal.returnAnimalWeight(), err)
		return err
	}
	err = checkType(animal)
	if err != nil {
		err = fmt.Errorf("animal impersonate %s, but actually is %s: %w", animal.returnImpersonateStatus(),
			reflect.TypeOf(animal).Name(), err)
		return err
	}
	err = checkIfEdible(animal)
	if err != nil {
		return err
	}
	return nil
}

func checkType(animal animalInfoGetter) error {
	if reflect.TypeOf(animal).Name() != animal.returnImpersonateStatus() {
		return wrongType
	}
	return nil
}

func checkWeight(animal animalInfoGetter) error {
	if animal.returnAnimalWeight() < animal.returnMinimalWeight() {
		return lowWeight
	}
	return nil
}

func checkIfEdible(animal animalInfoGetter) error {
	if !animal.returnEdibleStatus() {
		return isNotEdible
	}
	return nil
}

func generateRandomResponse() (iAm string) {
	switch rand.Intn(3) {
	case 0:
		iAm = "cat"
	case 1:
		iAm = "dog"
	case 2:
		iAm = "cow"
	}
	return iAm
}
