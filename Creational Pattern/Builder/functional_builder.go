package main

import "fmt"

type Person struct {
	name, position string
}

type personMod func(*Person)
type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

// extend PersonBuilder
func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, fn := range b.actions {
		fn(&p)
	}
	return &p
}

func main() {
	b := PersonBuilder{}
	p := b.Called("pure").WorksAsA("Js").Build()
	fmt.Println(*p)
}
