package main

import "fmt"

// Memento зберігає стан документа
type Memento struct {
	state string
}

// TextDocument – документ із текстом
type TextDocument struct {
	text string
}

// SaveState створює Memento
func (d *TextDocument) SaveState() *Memento {
	return &Memento{state: d.text}
}

// RestoreState відновлює стан з Memento
func (d *TextDocument) RestoreState(m *Memento) {
	d.text = m.state
}

// SetText змінює вміст документа
func (d *TextDocument) SetText(text string) {
	d.text = text
}

// GetText отримує поточний текст
func (d *TextDocument) GetText() string {
	return d.text
}

// TextEditor керує документом та його історією змін
type TextEditor struct {
	document *TextDocument
	history  []*Memento
}

// NewTextEditor створює редактор із документом
func NewTextEditor() *TextEditor {
	return &TextEditor{
		document: &TextDocument{},
		history:  []*Memento{},
	}
}

// Save зберігає стан документа
func (e *TextEditor) Save() {
	e.history = append(e.history, e.document.SaveState())
}

// Undo скасовує останню зміну
func (e *TextEditor) Undo() {
	if len(e.history) == 0 {
		fmt.Println("Немає змін для скасування!")
		return
	}
	lastState := e.history[len(e.history)-1]
	e.history = e.history[:len(e.history)-1] // Видаляємо останній стан
	e.document.RestoreState(lastState)
}

// SetText встановлює новий текст у документ
func (e *TextEditor) SetText(text string) {
	e.document.SetText(text)
}

// GetText отримує поточний текст
func (e *TextEditor) GetText() string {
	return e.document.GetText()
}

// Головний метод програми
func main() {
	editor := NewTextEditor()

	// Додаємо текст
	editor.SetText("Перший рядок")
	editor.Save()

	editor.SetText("Другий рядок")
	editor.Save()

	editor.SetText("Третій рядок")

	fmt.Println("Поточний текст:", editor.GetText()) // Третій рядок

	// Відміна змін
	editor.Undo()
	fmt.Println("Після Undo:", editor.GetText()) // Другий рядок

	editor.Undo()
	fmt.Println("Після другого Undo:", editor.GetText()) // Перший рядок

	editor.Undo() // Немає змін для скасування
}
