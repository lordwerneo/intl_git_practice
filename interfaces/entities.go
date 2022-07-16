package main

import (
	"fmt"
	"reflect"
)

type cat struct {
	name            string
	impersonate     string
	weight          int
	feedConsumption int
	minimalWeight   int
	isEdible        bool
}

func (c cat) feedPerMonth() int {
	return c.weight * c.feedConsumption
}

func (c cat) String() string {
	return fmt.Sprintf("%s named %s", reflect.TypeOf(c).Name(), c.name)
}

func (c cat) returnAnimalWeight() int {
	return c.weight
}

func (c cat) returnMinimalWeight() int {
	return c.minimalWeight
}

func (c cat) returnEdibleStatus() bool {
	return c.isEdible
}

func (c cat) returnImpersonateStatus() string {
	return c.impersonate
}

type dog struct {
	name            string
	impersonate     string
	weight          int
	feedConsumption int
	minimalWeight   int
	isEdible        bool
}

func (d dog) feedPerMonth() int {
	return d.weight * d.feedConsumption
}

func (d dog) String() string {
	return fmt.Sprintf("%s named %s", reflect.TypeOf(d).Name(), d.name)
}

func (d dog) returnAnimalWeight() int {
	return d.weight
}

func (d dog) returnMinimalWeight() int {
	return d.minimalWeight
}

func (d dog) returnEdibleStatus() bool {
	return d.isEdible
}

func (d dog) returnImpersonateStatus() string {
	return d.impersonate
}

type cow struct {
	name            string
	impersonate     string
	weight          int
	feedConsumption int
	minimalWeight   int
	isEdible        bool
}

func (c cow) feedPerMonth() int {
	return c.weight * c.feedConsumption
}

func (c cow) String() string {
	return fmt.Sprintf("%s named %s", reflect.TypeOf(c).Name(), c.name)
}

func (c cow) returnAnimalWeight() int {
	return c.weight
}

func (c cow) returnMinimalWeight() int {
	return c.minimalWeight
}

func (c cow) returnEdibleStatus() bool {
	return c.isEdible
}

func (c cow) returnImpersonateStatus() string {
	return c.impersonate
}
