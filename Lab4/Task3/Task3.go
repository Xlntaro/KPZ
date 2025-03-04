package main

import (
	"fmt"
)

// EventManager керує підписками та подіями
type EventManager struct {
	events map[string][]func()
}

func NewEventManager() *EventManager {
	return &EventManager{events: make(map[string][]func())}
}

func (e *EventManager) Subscribe(eventType string, listener func()) {
	e.events[eventType] = append(e.events[eventType], listener)
}

func (e *EventManager) Dispatch(eventType string) {
	if listeners, found := e.events[eventType]; found {
		for _, listener := range listeners {
			listener()
		}
	}
}

// LightNode - базовий інтерфейс для всіх елементів
 type LightNode interface {
	Render() string
}

// LightTextNode - текстовий вузол
type LightTextNode struct {
	text string
}

func (t LightTextNode) Render() string {
	return t.text
}

// LightElementNode - HTML-елемент
type LightElementNode struct {
	tagName   string
	children  []LightNode
	events    *EventManager
}

func NewLightElementNode(tagName string) *LightElementNode {
	return &LightElementNode{
		tagName:  tagName,
		events:   NewEventManager(),
	}
}

func (e *LightElementNode) AddChild(child LightNode) {
	e.children = append(e.children, child)
}

func (e *LightElementNode) AddEventListener(eventType string, listener func()) {
	e.events.Subscribe(eventType, listener)
}

func (e *LightElementNode) DispatchEvent(eventType string) {
	e.events.Dispatch(eventType)
}

func (e *LightElementNode) Render() string {
	content := ""
	for _, child := range e.children {
		content += child.Render()
	}
	return fmt.Sprintf("<%s>%s</%s>", e.tagName, content, e.tagName)
}

func main() {
	h1 := NewLightElementNode("h1")
	h1.AddChild(LightTextNode{"Привіт, світ!"})

	button := NewLightElementNode("button")
	button.AddChild(LightTextNode{"Натисни мене!"})

	button.AddEventListener("click", func() {
		fmt.Println("Кнопка була натиснута!")
	})

	fmt.Println("Рендер HTML:")
	fmt.Println(h1.Render())
	fmt.Println(button.Render())

	fmt.Println("Симуляція натискання кнопки:")
	button.DispatchEvent("click")
}
