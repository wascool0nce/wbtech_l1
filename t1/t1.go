package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
	ActionType string
}

func (h *Human) Birthday() {
	h.Age++
}

func (a Action) FavoriteAction() {
	fmt.Println(a.Name, "favorite action", a.ActionType)
}

func main() {

	a := Action{
		Human: Human{
			Name: "Maks",
			Age:  23,
		},
		ActionType: "running",
	}
	fmt.Println(a.Age)
	a.Birthday()
	fmt.Println(a.Age)
	a.FavoriteAction()
}
