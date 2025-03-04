package main

import (
	"fmt"
	"strings"
)

// LightNode - базовий інтерфейс для всіх вузлів
type LightNode interface {
	OuterHTML() string
	InnerHTML() string
}

// LightTextNode - текстовий вузол
type LightTextNode struct {
	Text string
}

// OuterHTML і InnerHTML для текстового вузла просто повертає його текст
func (t LightTextNode) OuterHTML() string { return t.Text }
func (t LightTextNode) InnerHTML() string { return t.Text }

// Типи тегів
type DisplayType string
type ClosingType string

const (
	Block     DisplayType = "block"
	Inline    DisplayType = "inline"
	SelfClose ClosingType = "self-closing"
	Normal    ClosingType = "normal"
)

// LightElementNode - елементний вузол (може містити вкладені елементи)
type LightElementNode struct {
	TagName      string
	Display      DisplayType
	Closing      ClosingType
	CSSClasses   []string
	Children     []LightNode
}

// Генерує InnerHTML (вміст без зовнішнього тега)
func (e LightElementNode) InnerHTML() string {
	var innerContent strings.Builder
	for _, child := range e.Children {
		innerContent.WriteString(child.OuterHTML())
	}
	return innerContent.String()
}

// Генерує OuterHTML (включаючи зовнішній тег)
func (e LightElementNode) OuterHTML() string {
	var builder strings.Builder

	// Відкриваючий тег
	builder.WriteString(fmt.Sprintf("<%s", e.TagName))
	if len(e.CSSClasses) > 0 {
		builder.WriteString(fmt.Sprintf(" class=\"%s\"", strings.Join(e.CSSClasses, " ")))
	}
	builder.WriteString(">")

	// Вміст елемента
	if e.Closing == Normal {
		builder.WriteString(e.InnerHTML())
		builder.WriteString(fmt.Sprintf("</%s>", e.TagName))
	} else {
		builder.WriteString(" />")
	}

	return builder.String()
}

// Додає дочірній елемент
func (e *LightElementNode) AppendChild(child LightNode) {
	e.Children = append(e.Children, child)
}

// Додає CSS-клас
func (e *LightElementNode) AddClass(class string) {
	e.CSSClasses = append(e.CSSClasses, class)
}

func main() {
	// Створення сторінки
	body := LightElementNode{TagName: "body", Display: Block, Closing: Normal}

	// Створюємо div з класом container
	div := LightElementNode{TagName: "div", Display: Block, Closing: Normal}
	div.AddClass("container")

	// Додаємо заголовок
	h1 := LightElementNode{TagName: "h1", Display: Block, Closing: Normal}
	h1.AppendChild(LightTextNode{Text: "Welcome to LightHTML!"})

	// Додаємо параграф
	p := LightElementNode{TagName: "p", Display: Block, Closing: Normal}
	p.AppendChild(LightTextNode{Text: "This is a simple example of a custom markup language in Go."})

	// Додаємо до div
	div.AppendChild(h1)
	div.AppendChild(p)

	// Додаємо в body
	body.AppendChild(div)

	// Вивід готового HTML
	fmt.Println(body.OuterHTML())
}
