package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// 创建几个代表预定义员工类型的常量
const (
	Developer = iota
	Manager
)

// 此工厂函数以role为唯一的参数
func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 60000}
	case Manager:
		return &Employee{"", "manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
