package main

import "fmt"

// 里氏替换原则：所有引用基类的地方必须能透明地使用其子类的对象

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

// USE POINTER !!!
func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// modified LSP
// If a function takes an interface and
// works with a type T that implements this
// interface, any structure that aggregates T
// should also be usable in that function.
type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// 注意，对Square进行Set时，同时改变了Width和Height
// 用"子类Square"去调用"父类Rectangle"的方法，会输出预期之外的面积
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

type Square2 struct {
	size int
}

func NewSquare2(size int) *Square2 {
	sq2 := Square2{}
	sq2.size = size
	return &sq2
}

func (s *Square2) Rectangle() *Rectangle {
	return &Rectangle{s.size, s.size}
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea,
		", but got ", actualArea, "\n")
}

func main_() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5) // 预期之外的结果
	UseIt(sq)

	sq2 := NewSquare2(10)
	tmp := sq2.Rectangle() // 提供一个将子类"适配"父类的函数？
	UseIt(tmp)
}
