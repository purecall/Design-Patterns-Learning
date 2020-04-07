package main

import "fmt"

/*
	工厂方法是粒度很小的设计模式，因为模式的表现只是一个抽象的方法。
	在工厂和产品中间增加接口，工厂不再负责产品的创建，由接口根据不同条件返回具体的类实例
	完全遵循"开放-关闭原则"，可扩展
*/

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old\n", p.name, p.age)
}

type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Sorry, I am too tired to talk")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("pure", 20) // Hi, my name is pure, I am 20 years old
	//p := NewPerson("pure", 123) // Sorry, I am too tired to talk
	p.SayHello()
}
