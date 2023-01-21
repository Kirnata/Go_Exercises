package main

import "fmt"

type Action struct {
	Human
}

func (h Human) sayName() {
	fmt.Printf("Hello! My name %s)\n", h.name)
}

type Human struct {
	name string
	age  int
}

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)

func main() {
	a := Action{Human{name: "Nataly", age: 22}}
	a.sayName()
}
