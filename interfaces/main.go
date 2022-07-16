package main

import (
	"errors"
	"fmt"
	"math/rand"
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

var ErrWrongType = errors.New("wrong type")
var ErrLowWeight = errors.New("weight is too low")
var ErrIsNotEdible = errors.New("wrong edible status")

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
		return
	}
	fmt.Printf("Total amount of feed needed: %d", totalFeed)
}

func calculateFoodConsumption(animals []animalInfoGetter) (int, error) {
	var totalFeed int
	for _, v := range animals {
		err := validateAnimals(v)
		if err != nil {
			switch {
			case errors.Is(err, ErrLowWeight):
				err = fmt.Errorf("validation failed: %w", err)
				err = fmt.Errorf("for %s: %w", v, err)
				return 0, err
			case errors.Is(err, ErrWrongType):
				err = fmt.Errorf("validation failed: %w", err)
				err = fmt.Errorf("for %s: %w", v, err)
				fmt.Println(err)
				continue
			case errors.Is(err, ErrIsNotEdible):
				err = fmt.Errorf("validation failed: %w", err)
				err = fmt.Errorf("for %s: %w", v, err)
				fmt.Println(err)
				continue
			default:
				fmt.Printf("unexpected error: %s\n", err)
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
	if err := checkWeight(animal); err != nil {
		err = fmt.Errorf("minimal weight is %d, but it weights %d: %w", animal.returnMinimalWeight(),
			animal.returnAnimalWeight(), err)
		return err
	}
	if err := checkType(animal); err != nil {
		err = fmt.Errorf("animal impersonate %s, but actually is %T: %w", animal.returnImpersonateStatus(),
			animal, err)
		return err
	}

	if err := checkIfEdible(animal); err != nil {
		err = fmt.Errorf("animal edible status is %v, but should be %v: %w", animal.returnEdibleStatus(),
			!animal.returnEdibleStatus(), err)
		return err
	}
	return nil
}

func checkType(animal animalInfoGetter) error {
	switch animal.(type) {
	case cat:
		if animal.returnImpersonateStatus() != "cat" {
			return ErrWrongType
		}
	case dog:
		if animal.returnImpersonateStatus() != "dog" {
			return ErrWrongType
		}
	case cow:
		if animal.returnImpersonateStatus() != "cow" {
			return ErrWrongType
		}
	}
	return nil
}

func checkWeight(animal animalInfoGetter) error {
	if animal.returnAnimalWeight() < animal.returnMinimalWeight() {
		return ErrLowWeight
	}
	return nil
}

func checkIfEdible(animal animalInfoGetter) error {
	if !animal.returnEdibleStatus() {
		return ErrIsNotEdible
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
