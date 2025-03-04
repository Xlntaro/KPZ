package main

import (
	"fmt"
	"runtime"
	"strings"
)

// LightNode - базовий інтерфейс для HTML-елементів
type LightNode interface {
	OuterHTML() string
}

// LightTextNode - текстовий вузол
type LightTextNode struct {
	Text string
}

// OuterHTML для текстового вузла просто повертає його текст
func (t LightTextNode) OuterHTML() string { return t.Text }

// LightElementNode - елементний вузол
type LightElementNode struct {
	TagName    string
	Attributes string
	Children   []LightNode
}

// OuterHTML - створює HTML-код для вузла
func (e LightElementNode) OuterHTML() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("<%s%s>", e.TagName, e.Attributes))

	for _, child := range e.Children {
		builder.WriteString(child.OuterHTML())
	}

	builder.WriteString(fmt.Sprintf("</%s>", e.TagName))
	return builder.String()
}

// AppendChild додає дочірній елемент
func (e *LightElementNode) AppendChild(child LightNode) {
	e.Children = append(e.Children, child)
}

// FlyweightFactory - зберігає унікальні HTML-елементи
type FlyweightFactory struct {
	elements map[string]*LightElementNode
}

// GetElement повертає або створює новий вузол
func (f *FlyweightFactory) GetElement(tagName string) *LightElementNode {
	if el, exists := f.elements[tagName]; exists {
		return el
	}
	newEl := &LightElementNode{TagName: tagName}
	f.elements[tagName] = newEl
	return newEl
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{elements: make(map[string]*LightElementNode)}
}

func parseTextToHTML(text string, factory *FlyweightFactory) *LightElementNode {
	lines := strings.Split(text, "\n")
	root := &LightElementNode{TagName: "body"}

	for i, line := range lines {
		line = strings.TrimRight(line, "\r") // Видалити можливі `\r` в Windows

		var node *LightElementNode
		if i == 0 {
			node = factory.GetElement("h1")
		} else if len(line) < 20 {
			node = factory.GetElement("h2")
		} else if strings.HasPrefix(line, " ") {
			node = factory.GetElement("blockquote")
		} else {
			node = factory.GetElement("p")
		}

		newNode := *node // Клонуємо об'єкт, щоб не змінювати загальний шаблон
		newNode.Children = []LightNode{LightTextNode{Text: line}}
		root.AppendChild(&newNode)
	}
	return root
}

func main() {
	text := `The Great Adventure
Chapter One

    It was a dark and stormy night.
Suddenly, a shot rang out.
The maid screamed.

    A door slammed.
The clock struck twelve.

Silence.`

	factory := NewFlyweightFactory()
	htmlTree := parseTextToHTML(text, factory)

	fmt.Println(htmlTree.OuterHTML())

	// Обрахунок пам'яті
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("Memory usage: %d KB\n", mem.Alloc/1024)
}
