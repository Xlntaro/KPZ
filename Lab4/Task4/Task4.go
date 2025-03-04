package main

import (
	"fmt"
	"strings"
)

// Стратегія завантаження зображень
// ImageLoaderStrategy - інтерфейс стратегії

type ImageLoaderStrategy interface {
	LoadImage(href string) string
}

// FileImageLoader - стратегія для локальних файлів
type FileImageLoader struct{}

func (f FileImageLoader) LoadImage(href string) string {
	return fmt.Sprintf("Loading image from file: %s", href)
}

// NetworkImageLoader - стратегія для мережевих зображень
type NetworkImageLoader struct{}

func (n NetworkImageLoader) LoadImage(href string) string {
	return fmt.Sprintf("Downloading image from URL: %s", href)
}

// LightImageNode - елемент LightHTML для зображень

type LightImageNode struct {
	src      string
	loader  ImageLoaderStrategy
}

func NewLightImageNode(href string) *LightImageNode {
	var loader ImageLoaderStrategy
	if strings.HasPrefix(href, "file://") {
		loader = FileImageLoader{}
	} else if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
		loader = NetworkImageLoader{}
	} else {
		panic("Unsupported image source")
	}
	return &LightImageNode{src: href, loader: loader}
}

func (img *LightImageNode) Render() string {
	return img.loader.LoadImage(img.src)
}

func main() {
	// Тестуємо різні варіанти
	fileImage := NewLightImageNode("file://images/photo.jpg")
	netImage := NewLightImageNode("https://example.com/photo.jpg")

	fmt.Println(fileImage.Render())
	fmt.Println(netImage.Render())
}
