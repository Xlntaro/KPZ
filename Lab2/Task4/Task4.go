package main

import (
	"fmt"
)

// Virus — структура вірусу
type Virus struct {
	Name     string
	Type     string
	Age      int
	Weight   float64
	Children []*Virus
}

// Clone створює глибоку копію вірусу, включаючи всіх його дітей
func (v *Virus) Clone() *Virus {
	// Створюємо новий екземпляр вірусу
	clone := &Virus{
		Name:   v.Name + "_clone",
		Type:   v.Type,
		Age:    v.Age,
		Weight: v.Weight,
	}

	// Клонування дітей (рекурсивно)
	for _, child := range v.Children {
		clonedChild := child.Clone()
		clone.Children = append(clone.Children, clonedChild)
	}

	return clone
}

// PrintHierarchy рекурсивно виводить ієрархію вірусів
func (v *Virus) PrintHierarchy(level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	fmt.Printf("%s- %s (Вид: %s, Вік: %d, Вага: %.2f)\n", indent, v.Name, v.Type, v.Age, v.Weight)
	for _, child := range v.Children {
		child.PrintHierarchy(level + 1)
	}
}

func main() {
	// Створення сімейства вірусів (3 покоління)
	parent := &Virus{Name: "Alpha", Type: "Corona", Age: 5, Weight: 1.2}
	child1 := &Virus{Name: "Beta", Type: "Corona", Age: 2, Weight: 0.8}
	child2 := &Virus{Name: "Gamma", Type: "Corona", Age: 1, Weight: 0.5}
	grandchild := &Virus{Name: "Delta", Type: "Corona", Age: 0, Weight: 0.3}

	// Формування ієрархії
	child1.Children = append(child1.Children, grandchild)
	parent.Children = append(parent.Children, child1, child2)

	// Вивід оригінальної ієрархії
	fmt.Println("🌿 Оригінальна ієрархія вірусів:")
	parent.PrintHierarchy(0)

	// Клонування
	cloneParent := parent.Clone()

	// Вивід клонованої ієрархії
	fmt.Println("\n🧬 Клонована ієрархія вірусів:")
	cloneParent.PrintHierarchy(0)
}
