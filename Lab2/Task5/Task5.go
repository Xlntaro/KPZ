package main

import "fmt"

// Character - структура для героя або ворога
type Character struct {
	Name      string
	Height    float64
	Build     string
	HairColor string
	EyeColor  string
	Clothing  string
	Inventory []string
	Alignment string
	Actions   []string
}

// CharacterBuilder - інтерфейс для білдера персонажа
type CharacterBuilder interface {
	SetName(name string) CharacterBuilder
	SetHeight(height float64) CharacterBuilder
	SetBuild(build string) CharacterBuilder
	SetHairColor(color string) CharacterBuilder
	SetEyeColor(color string) CharacterBuilder
	SetClothing(clothing string) CharacterBuilder
	AddItemToInventory(item string) CharacterBuilder
	DoAction(action string) CharacterBuilder
	SetAlignment(alignment string) CharacterBuilder
	Build() Character
}

// HeroBuilder - білдер для героя
type HeroBuilder struct {
	character Character
}

func NewHeroBuilder() *HeroBuilder {
	return &HeroBuilder{character: Character{Alignment: "Good"}}
}

func (b *HeroBuilder) SetName(name string) CharacterBuilder {
	b.character.Name = name
	return b
}

func (b *HeroBuilder) SetHeight(height float64) CharacterBuilder {
	b.character.Height = height
	return b
}

func (b *HeroBuilder) SetBuild(build string) CharacterBuilder {
	b.character.Build = build
	return b
}

func (b *HeroBuilder) SetHairColor(color string) CharacterBuilder {
	b.character.HairColor = color
	return b
}

func (b *HeroBuilder) SetEyeColor(color string) CharacterBuilder {
	b.character.EyeColor = color
	return b
}

func (b *HeroBuilder) SetClothing(clothing string) CharacterBuilder {
	b.character.Clothing = clothing
	return b
}

func (b *HeroBuilder) AddItemToInventory(item string) CharacterBuilder {
	b.character.Inventory = append(b.character.Inventory, item)
	return b
}

func (b *HeroBuilder) DoAction(action string) CharacterBuilder {
	b.character.Actions = append(b.character.Actions, action)
	return b
}

func (b *HeroBuilder) SetAlignment(alignment string) CharacterBuilder {
	b.character.Alignment = alignment
	return b
}

func (b *HeroBuilder) Build() Character {
	return b.character
}

// EnemyBuilder - білдер для ворога
type EnemyBuilder struct {
	character Character
}

func NewEnemyBuilder() *EnemyBuilder {
	return &EnemyBuilder{character: Character{Alignment: "Evil"}}
}

func (b *EnemyBuilder) SetName(name string) CharacterBuilder {
	b.character.Name = name
	return b
}

func (b *EnemyBuilder) SetHeight(height float64) CharacterBuilder {
	b.character.Height = height
	return b
}

func (b *EnemyBuilder) SetBuild(build string) CharacterBuilder {
	b.character.Build = build
	return b
}

func (b *EnemyBuilder) SetHairColor(color string) CharacterBuilder {
	b.character.HairColor = color
	return b
}

func (b *EnemyBuilder) SetEyeColor(color string) CharacterBuilder {
	b.character.EyeColor = color
	return b
}

func (b *EnemyBuilder) SetClothing(clothing string) CharacterBuilder {
	b.character.Clothing = clothing
	return b
}

func (b *EnemyBuilder) AddItemToInventory(item string) CharacterBuilder {
	b.character.Inventory = append(b.character.Inventory, item)
	return b
}

func (b *EnemyBuilder) DoAction(action string) CharacterBuilder {
	b.character.Actions = append(b.character.Actions, action)
	return b
}

func (b *EnemyBuilder) SetAlignment(alignment string) CharacterBuilder {
	b.character.Alignment = alignment
	return b
}

func (b *EnemyBuilder) Build() Character {
	return b.character
}

// CharacterDirector - директор для контролю процесу створення
type CharacterDirector struct {
	builder CharacterBuilder
}

func NewCharacterDirector(b CharacterBuilder) *CharacterDirector {
	return &CharacterDirector{builder: b}
}

func (d *CharacterDirector) CreateWarrior(name string) Character {
	return d.builder.SetName(name).
		SetHeight(1.85).
		SetBuild("Muscular").
		SetHairColor("Black").
		SetEyeColor("Brown").
		SetClothing("Armor").
		AddItemToInventory("Sword").
		DoAction("Fights evil").
		Build()
}

func (d *CharacterDirector) CreateVillain(name string) Character {
	return d.builder.SetName(name).
		SetHeight(1.80).
		SetBuild("Slim").
		SetHairColor("White").
		SetEyeColor("Red").
		SetClothing("Dark Robe").
		AddItemToInventory("Poison Dagger").
		DoAction("Plots revenge").
		Build()
}

func main() {
	// Створюємо героя
	heroBuilder := NewHeroBuilder()
	director := NewCharacterDirector(heroBuilder)
	hero := director.CreateWarrior("Arthur")

	// Створюємо ворога
	enemyBuilder := NewEnemyBuilder()
	director = NewCharacterDirector(enemyBuilder)
	villain := director.CreateVillain("Dark Lord")

	// Виводимо інформацію про персонажів
	fmt.Println("🌟 Герой:")
	fmt.Printf("Ім'я: %s\nЗріст: %.2f м\nСтатура: %s\nКолір волосся: %s\nКолір очей: %s\nОдяг: %s\nІнвентар: %v\nДії: %v\n\n",
		hero.Name, hero.Height, hero.Build, hero.HairColor, hero.EyeColor, hero.Clothing, hero.Inventory, hero.Actions)

	fmt.Println("💀 Ворог:")
	fmt.Printf("Ім'я: %s\nЗріст: %.2f м\nСтатура: %s\nКолір волосся: %s\nКолір очей: %s\nОдяг: %s\nІнвентар: %v\nДії: %v\n",
		villain.Name, villain.Height, villain.Build, villain.HairColor, villain.EyeColor, villain.Clothing, villain.Inventory, villain.Actions)
}
