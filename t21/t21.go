package main

import (
	"fmt"
)

type Dog struct{}

func (dog *Dog) WoofWoof() {
	fmt.Println("Гав-Гав")
}

type Cat struct{}

func (cat *Cat) MeowMeow(isCall bool) {
	if isCall {
		fmt.Println("Мяу-мяу")
	}
}

type AnimalAdapter interface {
	Reaction()
}

type DogAdapter struct {
	*Dog
}

func (adapter *DogAdapter) Reaction() {
	adapter.WoofWoof()
}

func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

type CatAdapter struct {
	*Cat
}

func (adapter *CatAdapter) Reaction() {
	adapter.MeowMeow(true)
}

func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

func main() {
	animalCall := [2]AnimalAdapter{NewDogAdapter(&Dog{}), NewCatAdapter(&Cat{})}
	for _, member := range animalCall {
		member.Reaction()
	}
}
