package main

import (
	"fmt"
	"os"
)

// Logger - клас для логування у консоль
type Logger struct{}

func (l Logger) Log(message string) {
	fmt.Println("\033[32m[INFO]:", message, "\033[0m") // Зелений текст
}

func (l Logger) Error(message string) {
	fmt.Println("\033[31m[ERROR]:", message, "\033[0m") // Червоний текст
}

func (l Logger) Warn(message string) {
	fmt.Println("\033[33m[WARN]:", message, "\033[0m") // Оранжевий текст
}

// FileWriter - клас для запису у файл
type FileWriter struct {
	filename string
}

func NewFileWriter(filename string) *FileWriter {
	return &FileWriter{filename: filename}
}

func (fw *FileWriter) Write(text string) error {
	file, err := os.OpenFile(fw.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	return err
}

func (fw *FileWriter) WriteLine(text string) error {
	return fw.Write(text + "\n")
}

// FileLoggerAdapter - адаптер для запису логів у файл
type FileLoggerAdapter struct {
	writer *FileWriter
}

func NewFileLoggerAdapter(filename string) *FileLoggerAdapter {
	return &FileLoggerAdapter{writer: NewFileWriter(filename)}
}

func (fla *FileLoggerAdapter) Log(message string) {
	_ = fla.writer.WriteLine("[INFO]: " + message)
}

func (fla *FileLoggerAdapter) Error(message string) {
	_ = fla.writer.WriteLine("[ERROR]: " + message)
}

func (fla *FileLoggerAdapter) Warn(message string) {
	_ = fla.writer.WriteLine("[WARN]: " + message)
}

func main() {
	// Використання звичайного Logger
	consoleLogger := Logger{}
	consoleLogger.Log("Це інформаційне повідомлення")
	consoleLogger.Warn("Це попередження")
	consoleLogger.Error("Це помилка")

	// Використання FileLoggerAdapter
	fileLogger := NewFileLoggerAdapter("log.txt")
	fileLogger.Log("Це інформаційне повідомлення")
	fileLogger.Warn("Це попередження")
	fileLogger.Error("Це помилка")

	fmt.Println("Логи записано у файл log.txt")
}
