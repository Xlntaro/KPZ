package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Інтерфейс для читання файлу
type TextReader interface {
	ReadFile(filename string) ([][]rune, error)
}

// SmartTextReader - клас, що читає текстовий файл і повертає двовимірний масив
type SmartTextReader struct{}

func (r SmartTextReader) ReadFile(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// SmartTextChecker - проксі, що логує операції
type SmartTextChecker struct {
	reader TextReader
}

func (c SmartTextChecker) ReadFile(filename string) ([][]rune, error) {
	fmt.Println("Opening file:", filename)

	content, err := c.reader.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	fmt.Println("Successfully read file:", filename)
	totalLines := len(content)
	totalChars := 0
	for _, line := range content {
		totalChars += len(line)
	}

	fmt.Printf("Total lines: %d, Total characters: %d\n", totalLines, totalChars)
	fmt.Println("Closing file:", filename)

	return content, nil
}

// SmartTextReaderLocker - проксі, що обмежує доступ до файлів за регулярним виразом
type SmartTextReaderLocker struct {
	reader  TextReader
	pattern *regexp.Regexp
}

func NewSmartTextReaderLocker(reader TextReader, regexPattern string) (*SmartTextReaderLocker, error) {
	pattern, err := regexp.Compile(regexPattern)
	if err != nil {
		return nil, err
	}
	return &SmartTextReaderLocker{reader: reader, pattern: pattern}, nil
}

func (l SmartTextReaderLocker) ReadFile(filename string) ([][]rune, error) {
	if l.pattern.MatchString(filename) {
		fmt.Println("Access denied!")
		return nil, fmt.Errorf("access denied")
	}
	return l.reader.ReadFile(filename)
}

func main() {
	filename := "test.txt"

	// Звичайний читач
	reader := SmartTextReader{}

	// Проксі з логуванням
	loggerProxy := SmartTextChecker{reader: reader}

	// Проксі з обмеженням доступу
	lockerProxy, _ := NewSmartTextReaderLocker(reader, `.*secret.*`)

	// Читаємо файл через логуючий проксі
	fmt.Println("\n--- Logging Proxy ---")
	loggerProxy.ReadFile(filename)

	// Читаємо файл через проксі з обмеженням доступу
	fmt.Println("\n--- Access Control Proxy ---")
	lockerProxy.ReadFile("test.txt")       // Дозволений доступ
	lockerProxy.ReadFile("secret_file.txt") // Заборонений доступ
}
