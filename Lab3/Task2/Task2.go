package main

import "fmt"

// Hero - інтерфейс для всіх героїв
type Hero interface {
	GetStats() string
}

// Warrior - клас воїна
type Warrior struct{}

func (w Warrior) GetStats() string {
	return "Warrior: Сила 15, Витривалість 10"
}

// Mage - клас мага
type Mage struct{}

func (m Mage) GetStats() string {
	return "Mage: Інтелект 18, Мана 12"
}

// Paladin - клас паладина
type Paladin struct{}

func (p Paladin) GetStats() string {
	return "Paladin: Сила 12, Інтелект 10, Витривалість 12"
}

// --- Декоратори ---

// BaseDecorator - базовий декоратор
type BaseDecorator struct {
	hero Hero
}

func (b BaseDecorator) GetStats() string {
	return b.hero.GetStats()
}

// ArmorDecorator - додає броню
type ArmorDecorator struct {
	BaseDecorator
}

func (a ArmorDecorator) GetStats() string {
	return a.hero.GetStats() + " + Броня +5"
}

// WeaponDecorator - додає зброю
type WeaponDecorator struct {
	BaseDecorator
}

func (w WeaponDecorator) GetStats() string {
	return w.hero.GetStats() + " + Зброя +7 до атаки"
}

// ArtifactDecorator - додає артефакт
type ArtifactDecorator struct {
	BaseDecorator
}

func (a ArtifactDecorator) GetStats() string {
	return a.hero.GetStats() + " + Артефакт +10 до магії"
}

func main() {
	// Створюємо героя
	warrior := Warrior{}
	fmt.Println(warrior.GetStats())

	// Додаємо броню
	warriorWithArmor := ArmorDecorator{BaseDecorator{warrior}}
	fmt.Println(warriorWithArmor.GetStats())

	// Додаємо зброю
	warriorWithWeapon := WeaponDecorator{BaseDecorator{warriorWithArmor}}
	fmt.Println(warriorWithWeapon.GetStats())

	// Додаємо артефакт
	warriorWithArtifact := ArtifactDecorator{BaseDecorator{warriorWithWeapon}}
	fmt.Println(warriorWithArtifact.GetStats())

	// Додаємо предмети паладину
	paladin := Paladin{}
	paladinFullyEquipped := ArtifactDecorator{
		BaseDecorator{
			WeaponDecorator{
				BaseDecorator{
					ArmorDecorator{
						BaseDecorator{paladin},
					},
				},
			},
		},
	}
	fmt.Println(paladinFullyEquipped.GetStats())
}
