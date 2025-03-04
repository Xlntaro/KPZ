package main

import "fmt"

// Інтерфейс пристрою
type Device interface {
	GetName() string
	GetBrand() string
}

// Реалізації пристроїв
type Laptop struct{ brand string }
func (l Laptop) GetName() string  { return "Laptop" }
func (l Laptop) GetBrand() string { return l.brand }

type Netbook struct{ brand string }
func (n Netbook) GetName() string  { return "Netbook" }
func (n Netbook) GetBrand() string { return n.brand }

type EBook struct{ brand string }
func (e EBook) GetName() string  { return "E-Book" }
func (e EBook) GetBrand() string { return e.brand }

type Smartphone struct{ brand string }
func (s Smartphone) GetName() string  { return "Smartphone" }
func (s Smartphone) GetBrand() string { return s.brand }

// Інтерфейс фабрики брендів
type BrandFactory interface {
	CreateLaptop() Device
	CreateNetbook() Device
	CreateEBook() Device
	CreateSmartphone() Device
}

// Фабрика бренду IProne
type IProneFactory struct{}
func (IProneFactory) CreateLaptop() Device      { return Laptop{"IProne"} }
func (IProneFactory) CreateNetbook() Device     { return Netbook{"IProne"} }
func (IProneFactory) CreateEBook() Device       { return EBook{"IProne"} }
func (IProneFactory) CreateSmartphone() Device  { return Smartphone{"IProne"} }

// Фабрика бренду Kiaomi
type KiaomiFactory struct{}
func (KiaomiFactory) CreateLaptop() Device      { return Laptop{"Kiaomi"} }
func (KiaomiFactory) CreateNetbook() Device     { return Netbook{"Kiaomi"} }
func (KiaomiFactory) CreateEBook() Device       { return EBook{"Kiaomi"} }
func (KiaomiFactory) CreateSmartphone() Device  { return Smartphone{"Kiaomi"} }

// Фабрика бренду Balaxy
type BalaxyFactory struct{}
func (BalaxyFactory) CreateLaptop() Device      { return Laptop{"Balaxy"} }
func (BalaxyFactory) CreateNetbook() Device     { return Netbook{"Balaxy"} }
func (BalaxyFactory) CreateEBook() Device       { return EBook{"Balaxy"} }
func (BalaxyFactory) CreateSmartphone() Device  { return Smartphone{"Balaxy"} }

func main() {
	var factory BrandFactory

	fmt.Println("Оберіть бренд:")
	fmt.Println("1 - IProne")
	fmt.Println("2 - Kiaomi")
	fmt.Println("3 - Balaxy")

	var brandChoice int
	fmt.Scanln(&brandChoice)

	switch brandChoice {
	case 1:
		factory = IProneFactory{}
	case 2:
		factory = KiaomiFactory{}
	case 3:
		factory = BalaxyFactory{}
	default:
		fmt.Println("Невірний вибір бренду!")
		return
	}

	fmt.Println("Оберіть пристрій:")
	fmt.Println("1 - Laptop")
	fmt.Println("2 - Netbook")
	fmt.Println("3 - E-Book")
	fmt.Println("4 - Smartphone")

	var deviceChoice int
	fmt.Scanln(&deviceChoice)

	var device Device

	switch deviceChoice {
	case 1:
		device = factory.CreateLaptop()
	case 2:
		device = factory.CreateNetbook()
	case 3:
		device = factory.CreateEBook()
	case 4:
		device = factory.CreateSmartphone()
	default:
		fmt.Println("Невірний вибір пристрою!")
		return
	}

	fmt.Printf("Ви створили: %s (%s)\n", device.GetName(), device.GetBrand())
}
