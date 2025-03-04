package main

import (
	"fmt"
)

// Handler - інтерфейс обробника запитів
type Handler interface {
	HandleRequest()
	SetNext(handler Handler)
}

// BaseHandler - базовий обробник для збереження наступного обробника
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) Next() {
	if b.next != nil {
		b.next.HandleRequest()
	} else {
		fmt.Println("Вибачте, ми не змогли знайти відповідь на ваше питання. Зверніться до служби підтримки.")
	}
}

// Level1Handler - перший рівень підтримки (запитує загальну тему)
type Level1Handler struct {
	BaseHandler
}

func (h *Level1Handler) HandleRequest() {
	fmt.Println("Вітаємо в службі підтримки. Оберіть категорію:")
	fmt.Println("1 - Проблеми з інтернетом")
	fmt.Println("2 - Питання по тарифах")
	fmt.Println("3 - Проблеми з мобільним зв'язком")
	fmt.Println("4 - Інші питання")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Переадресація до технічної підтримки...")
		h.Next()
	case 2:
		fmt.Println("Переадресація до фінансового відділу...")
		h.Next()
	case 3:
		fmt.Println("Переадресація до мобільного оператора...")
		h.Next()
	case 4:
		fmt.Println("З'єднання з оператором...")
		h.Next()
	default:
		fmt.Println("Невірний вибір. Спробуйте ще раз.")
		h.HandleRequest()
	}
}

// Level2Handler - другий рівень підтримки (технічні питання)
type Level2Handler struct {
	BaseHandler
}

func (h *Level2Handler) HandleRequest() {
	fmt.Println("Яку проблему з інтернетом ви маєте?")
	fmt.Println("1 - Інтернет не працює")
	fmt.Println("2 - Низька швидкість")
	fmt.Println("3 - Інша проблема")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Перевірте, чи підключений кабель або перезавантажте роутер.")
	case 2:
		fmt.Println("Можливо, завантаження каналу перевантажене. Перевірте пристрої в мережі.")
	case 3:
		fmt.Println("З'єднання з оператором для детальнішої перевірки.")
	default:
		fmt.Println("Невірний вибір. Спробуйте ще раз.")
		h.HandleRequest()
	}
}

// Level3Handler - третій рівень підтримки (питання по тарифах)
type Level3Handler struct {
	BaseHandler
}

func (h *Level3Handler) HandleRequest() {
	fmt.Println("Оберіть питання щодо тарифів:")
	fmt.Println("1 - Поточний тариф")
	fmt.Println("2 - Перехід на новий тариф")
	fmt.Println("3 - Поповнення рахунку")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Ваш поточний тариф: 'Оптимальний'.")
	case 2:
		fmt.Println("Доступні тарифи: 'Оптимальний', 'Максимальний', 'Економ'.")
	case 3:
		fmt.Println("Поповнення можливе через додаток, термінали або банківський рахунок.")
	default:
		fmt.Println("Невірний вибір. Спробуйте ще раз.")
		h.HandleRequest()
	}
}

// Level4Handler - четвертий рівень підтримки (загальні питання)
type Level4Handler struct {
	BaseHandler
}

func (h *Level4Handler) HandleRequest() {
	fmt.Println("З'єднання з оператором... Будь ласка, зачекайте.")
}

func main() {
	// Створення рівнів підтримки
	level1 := &Level1Handler{}
	level2 := &Level2Handler{}
	level3 := &Level3Handler{}
	level4 := &Level4Handler{}

	// Встановлення ланцюжка відповідальності
	level1.SetNext(level2)
	level2.SetNext(level3)
	level3.SetNext(level4)

	// Запуск обробки запитів
	level1.HandleRequest()
}
