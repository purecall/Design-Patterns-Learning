package main

import "fmt"

// Decouple abstraction from implementation

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRender struct {
	Dpi int
}

func (r *RasterRender) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func main() {
	raster := RasterRender{}
	circle1 := NewCircle(&raster, 5)
	circle1.Draw()

	vector := VectorRenderer{}
	circle2 := NewCircle(&vector, 5)
	circle2.Draw()
}
