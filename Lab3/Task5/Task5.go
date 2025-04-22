package main

import (
	"fmt"
	"strings"
)

// Інтерфейси вузлів

type LightNode interface {
	OuterHTML() string
	InnerHTML() string
	Accept(visitor Visitor)
	OnCreated()
}

// Текстовий вузол

type LightTextNode struct {
	Text string
}

func (t *LightTextNode) OuterHTML() string      { return t.Text }
func (t *LightTextNode) InnerHTML() string      { return t.Text }
func (t *LightTextNode) Accept(visitor Visitor) {}
func (t *LightTextNode) OnCreated() {
	fmt.Println("[hook] Text node created:", t.Text)
}

// Типи для елементів

type DisplayType string
type ClosingType string

const (
	Block     DisplayType = "block"
	Inline    DisplayType = "inline"
	SelfClose ClosingType = "self-closing"
	Normal    ClosingType = "normal"
)

// Елемент

type LightElementNode struct {
	TagName    string
	Display    DisplayType
	Closing    ClosingType
	CSSClasses []string
	Children   []LightNode
	State      RenderState
}

func (e *LightElementNode) OuterHTML() string {
	if e.State != nil && !e.State.ShouldRender() {
		return ""
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("<%s", e.TagName))
	if len(e.CSSClasses) > 0 {
		builder.WriteString(fmt.Sprintf(" class=\"%s\"", strings.Join(e.CSSClasses, " ")))
	}
	builder.WriteString(">")

	if e.Closing == Normal {
		builder.WriteString(e.InnerHTML())
		builder.WriteString(fmt.Sprintf("</%s>", e.TagName))
	} else {
		builder.WriteString(" />")
	}
	return builder.String()
}

func (e *LightElementNode) InnerHTML() string {
	var content strings.Builder
	for _, child := range e.Children {
		content.WriteString(child.OuterHTML())
	}
	return content.String()
}

func (e *LightElementNode) AppendChild(child LightNode) {
	e.Children = append(e.Children, child)
}

func (e *LightElementNode) AddClass(class string) {
	e.CSSClasses = append(e.CSSClasses, class)
}

func (e *LightElementNode) OnCreated() {
	fmt.Println("[hook] Element created:", e.TagName)
}

func (e *LightElementNode) Accept(visitor Visitor) {
	visitor.Visit(e)
	for _, child := range e.Children {
		child.Accept(visitor)
	}
}

// Ітератор

type NodeIterator interface {
	HasNext() bool
	Next() LightNode
}

// DFS

type DFSIterator struct {
	stack []LightNode
}

func NewDFSIterator(root LightNode) *DFSIterator {
	return &DFSIterator{stack: []LightNode{root}}
}

func (it *DFSIterator) HasNext() bool {
	return len(it.stack) > 0
}

func (it *DFSIterator) Next() LightNode {
	node := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]

	if elem, ok := node.(*LightElementNode); ok {
		for i := len(elem.Children) - 1; i >= 0; i-- {
			it.stack = append(it.stack, elem.Children[i])
		}
	}
	return node
}

// BFS

type BFSIterator struct {
	queue []LightNode
}

func NewBFSIterator(root LightNode) *BFSIterator {
	return &BFSIterator{queue: []LightNode{root}}
}

func (it *BFSIterator) HasNext() bool {
	return len(it.queue) > 0
}

func (it *BFSIterator) Next() LightNode {
	node := it.queue[0]
	it.queue = it.queue[1:]

	if elem, ok := node.(*LightElementNode); ok {
		it.queue = append(it.queue, elem.Children...)
	}
	return node
}

// State Pattern

type RenderState interface {
	ShouldRender() bool
}

type VisibleState struct{}
type HiddenState struct{}

func (v *VisibleState) ShouldRender() bool { return true }
func (h *HiddenState) ShouldRender() bool  { return false }

// Command Pattern

type Command interface {
	Execute()
	Undo()
}

type AddClassCommand struct {
	Element *LightElementNode
	Class   string
}

func (c *AddClassCommand) Execute() {
	c.Element.AddClass(c.Class)
}

func (c *AddClassCommand) Undo() {
	for i := len(c.Element.CSSClasses) - 1; i >= 0; i-- {
		if c.Element.CSSClasses[i] == c.Class {
			c.Element.CSSClasses = append(c.Element.CSSClasses[:i], c.Element.CSSClasses[i+1:]...)
			break
		}
	}
}

// Visitor Pattern

type Visitor interface {
	Visit(node *LightElementNode)
}

type TagCounterVisitor struct {
	Counts map[string]int
}

func NewTagCounterVisitor() *TagCounterVisitor {
	return &TagCounterVisitor{Counts: make(map[string]int)}
}

func (v *TagCounterVisitor) Visit(node *LightElementNode) {
	v.Counts[node.TagName]++
}

// main

func main() {
	body := &LightElementNode{TagName: "body", Display: Block, Closing: Normal, State: &VisibleState{}}

	div := &LightElementNode{TagName: "div", Display: Block, Closing: Normal, State: &VisibleState{}}
	div.AddClass("container")

	h1 := &LightElementNode{TagName: "h1", Display: Block, Closing: Normal, State: &VisibleState{}}
	h1.AppendChild(&LightTextNode{Text: "Welcome to LightHTML!"})

	p := &LightElementNode{TagName: "p", Display: Block, Closing: Normal, State: &VisibleState{}}
	p.AppendChild(&LightTextNode{Text: "This is a simple example of a custom markup language in Go."})

	div.AppendChild(h1)
	div.AppendChild(p)
	body.AppendChild(div)

	fmt.Println("DFS Traversal:")
	dfs := NewDFSIterator(body)
	for dfs.HasNext() {
		node := dfs.Next()
		fmt.Println("-", node.OuterHTML())
	}

	fmt.Println("\nBFS Traversal:")
	bfs := NewBFSIterator(body)
	for bfs.HasNext() {
		node := bfs.Next()
		fmt.Println("-", node.OuterHTML())
	}

	fmt.Println("\nLifecycle hooks:")
	body.OnCreated()
	div.OnCreated()
	h1.OnCreated()
	p.OnCreated()

	fmt.Println("\nTag counts:")
	visitor := NewTagCounterVisitor()
	body.Accept(visitor)
	for tag, count := range visitor.Counts {
		fmt.Printf("Tag <%s>: %d\n", tag, count)
	}

	fmt.Println("\nAddClassCommand demo:")
	cmd := &AddClassCommand{Element: div, Class: "new-style"}
	cmd.Execute()
	fmt.Println("After Execute:", div.OuterHTML())
	cmd.Undo()
	fmt.Println("After Undo:", div.OuterHTML())

	fmt.Println("\nState pattern demo:")
	p.State = &HiddenState{}
	fmt.Println("Paragraph should be hidden now:")
	fmt.Println(p.OuterHTML())

	fmt.Println("\nFinal HTML:")
	fmt.Println(body.OuterHTML())
}
