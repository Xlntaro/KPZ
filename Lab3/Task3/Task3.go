package main

import "fmt"

// Renderer - інтерфейс для рендерингу
type Renderer interface {
	Render(shape string)
}

// VectorRenderer - Векторний рендерер
type VectorRenderer struct{}

func (v VectorRenderer) Render(shape string) {
	fmt.Println("Drawing", shape, "as vector graphics")
}

// RasterRenderer - Растровий рендерер
type RasterRenderer struct{}

func (r RasterRenderer) Render(shape string) {
	fmt.Println("Drawing", shape, "as pixels")
}

// Shape - базовий клас для фігур
type Shape struct {
	renderer Renderer
	name     string
}

func (s Shape) Draw() {
	s.renderer.Render(s.name)
}

// Circle - Клас Коло
type Circle struct {
	Shape
}

func NewCircle(renderer Renderer) Circle {
	return Circle{Shape{renderer, "Circle"}}
}

// Square - Клас Квадрат
type Square struct {
	Shape
}

func NewSquare(renderer Renderer) Square {
	return Square{Shape{renderer, "Square"}}
}

// Triangle - Клас Трикутник
type Triangle struct {
	Shape
}

func NewTriangle(renderer Renderer) Triangle {
	return Triangle{Shape{renderer, "Triangle"}}
}

func main() {
	// Використання векторного рендеру
	vectorRenderer := VectorRenderer{}
	circleVector := NewCircle(vectorRenderer)
	squareVector := NewSquare(vectorRenderer)
	triangleVector := NewTriangle(vectorRenderer)

	circleVector.Draw()
	squareVector.Draw()
	triangleVector.Draw()

	// Використання растрового рендеру
	rasterRenderer := RasterRenderer{}
	circleRaster := NewCircle(rasterRenderer)
	squareRaster := NewSquare(rasterRenderer)
	triangleRaster := NewTriangle(rasterRenderer)

	circleRaster.Draw()
	squareRaster.Draw()
	triangleRaster.Draw()
}
